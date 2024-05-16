package cmd

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func TestFrameConn(t *testing.T) {
	//server
	l, err := net.Listen("tcp", ":9981")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
				return
			}

			c := NewLineBasedFrameConn(conn)
			go func(conn FrameConn) {
				for {
					b, err := c.ReadFrame()
					if err != nil {
						if err == io.EOF {
							return
						}
					}
					fmt.Println(string(b))
				}
			}(c)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:9981")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	fc := NewLineBasedFrameConn(conn)
	fc.WriteFrame([]byte("hello "))
	fc.WriteFrame([]byte("world"))

	time.Sleep(1 * time.Second)
}

func TestFrameConnLengthBased(t *testing.T) {
	//server
	l, err := net.Listen("tcp", ":9981")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	encoderConfig := EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               2,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}

	decoderConfig := DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldLength:   2,
		LengthFieldOffset:   0,
		LengthAdjustment:    0,
		InitialBytesToStrip: 2,
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
				return
			}

			c := NewLengthFieldBasedFrameConn(encoderConfig, decoderConfig, conn)
			go func(conn FrameConn) {
				for {
					b, err := c.ReadFrame()
					if err != nil {
						if err == io.EOF {
							return
						}
					}
					fmt.Println(string(b))
				}
			}(c)
		}

	}()

	conn, err := net.Dial("tcp", "127.0.0.1:9981")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	fc := NewLengthFieldBasedFrameConn(encoderConfig, decoderConfig, conn)
	fc.WriteFrame([]byte("hello"))
	fc.WriteFrame([]byte("world"))

	time.Sleep(1 * time.Second)
}
