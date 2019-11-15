package pubsub

//
//import (
//	"fmt"
//	"github.com/gomodule/redigo/redis"
//	"github.com/gorilla/websocket"
//	"github.com/pkg/errors"
//	log "github.com/sirupsen/logrus"
//	"time"
//)
//
//const Channel = "spider"
//
//var (
//	waitingMessage, availableMessage []byte
//	waitSleep                        = time.Second * 10
//)
//
//// redisReceiver receives messages from Redis and broadcasts them to all
//// registered websocket connections that are Registered.
//type redisReceiver struct {
//	pool *redis.Pool
//
//	messages       chan []byte
//	newConnections chan *websocket.Conn
//	rmConnections  chan *websocket.Conn
//}
//
//// newRedisReceiver creates a redisReceiver that will use the provided
//// redis.Pool.
//func newRedisReceiver(pool *redis.Pool) redisReceiver {
//	return redisReceiver{
//		pool:           pool,
//		messages:       make(chan []byte, 1000), // 1000 is arbitrary
//		newConnections: make(chan *websocket.Conn),
//		rmConnections:  make(chan *websocket.Conn),
//	}
//}
//
//func (rr *redisReceiver) wait(_ time.Time) error {
//	rr.broadcast(waitingMessage)
//	time.Sleep(waitSleep)
//	return nil
//}
//
//// run receives pubsub messages from Redis after establishing a connection.
//// When a valid message is received it is broadcast to all connected websockets
//func (rr *redisReceiver) run() error {
//	l := log.WithField("channel", Channel)
//	conn := rr.pool.Get()
//	defer conn.Close()
//	psc := redis.PubSubConn{Conn: conn}
//	psc.Subscribe(Channel)
//	go rr.connHandler()
//	for {
//		switch v := psc.Receive().(type) {
//		case redis.Message:
//			rr.broadcast(v.Data)
//		case redis.Subscription:
//			l.WithFields(log.Fields{
//				"kind":  v.Kind,
//				"count": v.Count,
//			}).Println("Redis Subscription Received")
//			log.Println("Redis Subscription Received")
//		case error:
//			return errors.New("Error while subscribed to Redis channel")
//		default:
//			l.WithField("v", v).Info("Unknown Redis receive during subscription")
//			log.Println("Unknown Redis receive during subscription")
//		}
//	}
//}
//
//
//// broadcast the provided message to all connected websocket connections.
//// If an error occurs while writting a message to a websocket connection it is
//// closed and deregistered.
//func (rr *redisReceiver) broadcast(msg []byte) {
//	rr.messages <- msg
//}
//
//// register the websocket connection with the receiver.
//func (rr *redisReceiver) register(conn *websocket.Conn) {
//	rr.newConnections <- conn
//}
//
//// deRegister the connection by closing it and removing it from our list.
//func (rr *redisReceiver) deRegister(conn *websocket.Conn) {
//	rr.rmConnections <- conn
//}
//
//
//func (rr *redisReceiver) connHandler() {
//	conns := make([]*websocket.Conn, 0)
//	for {
//		select {
//		case msg := <-rr.messages:
//			for _, conn := range conns {
//				log.Printf("message is %v", msg)
//				if err := conn.WriteJSON(string(msg)); err != nil {
//					log.Println(err)
//					log.WithFields(log.Fields{
//						"data": msg,
//						"err":  err,
//						"conn": conn,
//					}).Error("Error writing data to connection! Closing and removing Connection")
//				}
//			}
//		case conn := <-rr.newConnections:
//			conns = append(conns, conn)
//		case conn := <-rr.rmConnections:
//			conns = removeConn(conns, conn)
//		}
//	}
//}
//
//func removeConn(conns []*websocket.Conn, remove *websocket.Conn) []*websocket.Conn {
//	var i int
//	var found bool
//	for i = 0; i < len(conns); i++ {
//		if conns[i] == remove {
//			found = true
//			break
//		}
//	}
//	if !found {
//		fmt.Printf("conns: %#v\nconn: %#v\n", conns, remove)
//		panic("Conn not found")
//	}
//	copy(conns[i:], conns[i+1:]) // shift down
//	conns[len(conns)-1] = nil    // nil last element
//	return conns[:len(conns)-1]  // truncate slice
//}
//
//// redisWriter publishes messages to the Redis CHANNEL
//type redisWriter struct {
//	pool     *redis.Pool
//	messages chan Reply
//}
//
//func newRedisWriter(pool *redis.Pool) redisWriter {
//	return redisWriter{
//		pool:     pool,
//		messages: make(chan Reply),
//	}
//}
//
//// run the main redisWriter loop that publishes incoming messages to Redis.
//func (rw *redisWriter) run() error {
//	conn := rw.pool.Get()
//	defer conn.Close()
//
//	for data := range rw.messages {
//		if err := writeToRedis(conn, data); err != nil {
//			rw.publish(data) // attempt to redeliver later
//			return err
//		}
//	}
//	return nil
//}
//
//func writeToRedis(conn redis.Conn, data Reply) error {
//	if err := conn.Send("PUBLISH", Channel, data); err != nil {
//		return errors.Wrap(err, "Unable to publish message to Redis")
//	}
//	if err := conn.Flush(); err != nil {
//		return errors.Wrap(err, "Unable to flush published message to Redis")
//	}
//	return nil
//}
//
//// publish to Redis via channel.
//func (rw *redisWriter) publish(data Reply) {
//	rw.messages <- data
//}
//
//// WaitFunc to be executed occasionally by something that is waiting.
//// Should return an error to cancel the waiting
//// Should also sleep some amount of time to throttle connection attempts
//type WaitFunc func(time.Time) error
//
//// WaitForAvailability of the redis server located at the provided url, timeout if the Duration passes before being able to connect
//func WaitForAvailability(url string, d time.Duration, f WaitFunc) (bool, error) {
//	conn := make(chan struct{})
//	errs := make(chan error)
//	go func() {
//		for {
//			c, err := redis.Dial("tcp", url)
//			log.Println(err)
//			if err == nil {
//				c.Close()
//				conn <- struct{}{}
//				return
//			}
//			if f != nil {
//				err := f(time.Now())
//				if err != nil {
//					errs <- err
//					return
//				}
//			}
//		}
//	}()
//	select {
//	case err := <-errs:
//		return false, err
//	case <-conn:
//		return true, nil
//	case <-time.After(d):
//		return false, nil
//	}
//}
//
//
//func NewRedisPoolFromURL(url string) (*redis.Pool, error) {
//	return &redis.Pool{
//		MaxIdle:     3,
//		IdleTimeout: 240 * time.Second,
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", url)
//			if err != nil {
//				return nil, err
//			}
//			return c, err
//		},
//		TestOnBorrow: func(c redis.Conn, t time.Time) error {
//			if time.Since(t) < time.Minute {
//				return nil
//			}
//			_, err := c.Do("PING")
//			return err
//		},
//	}, nil
//}
