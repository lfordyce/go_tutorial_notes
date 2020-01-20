package decorator

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// SortFunc type
type SortFunc func(sort.Interface)

// TimedSortFunc method for SortFunc
func TimedSortFunc(f SortFunc) SortFunc {
	return func(data sort.Interface) {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed: %v ---\n", time.Since(t))
		}(time.Now())
		f(data)
	}
}

// RandomHeroSort method for HeroScoreList
func RandomHeroSort(size int, upperLimit int) HeroScoreList {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	heroScores := make(HeroScoreList, size)

	for i := range heroScores {
		hero := heroScores[i]
		hero.name = "Batman"
		hero.likeCount = r.Intn(upperLimit)
	}
	return heroScores
}

// HeroScore struct
type HeroScore struct {
	likeCount int
	name      string
}

// HeroScoreList type
type HeroScoreList []HeroScore

func (s HeroScoreList) Len() int {
	return len(s)
}

func (s HeroScoreList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s HeroScoreList) Less(i, j int) bool {
	return s[i].likeCount < s[j].likeCount
}

// LikeFunc type
type LikeFunc func(int, int) bool

// LikePost function
func LikePost(userID int, postID int) bool {
	fmt.Printf("Upldate Complete!\n")
	return true
}

// DecoratedLike function
func DecoratedLike(f LikeFunc) LikeFunc {
	return func(userID int, postID int) bool {
		fmt.Printf("likePost Log: User %v liked post# %v\n", userID, postID)
		return f(userID, postID)
	}
}

// Adder interface to demonstrate decorator pattern
type Adder interface {
	Add(x, y int) int
}

// AdderFunc type
type AdderFunc func(x, y int) int

// AdderMiddleware type
type AdderMiddleware func(Adder) Adder

// Add function
func (a AdderFunc) Add(x, y int) int {
	return a(x, y)
}

// WrapLogger middleware
func WrapLogger(logger *log.Logger) AdderMiddleware {
	return func(a Adder) Adder {
		// Using `AdderFunc` to implement the `Adder` interface
		fn := func(x, y int) (result int) {
			defer func(t time.Time) {
				log.Printf("took=%v, x=%v, y=%v, result=%v", time.Since(t), x, y, result)
			}(time.Now())
			return a.Add(x, y)
		}
		return AdderFunc(fn)
	}
}

// WrapCache method for AdderMiddleware
func WrapCache(cache *sync.Map) AdderMiddleware {
	return func(a Adder) Adder {
		fn := func(x, y int) int {
			key := fmt.Sprintf("x=%dy=%d", x, y)
			val, ok := cache.Load(key)
			if ok {
				return val.(int)
			}
			result := a.Add(x, y)
			cache.Store(key, result)
			return result
		}
		return AdderFunc(fn)
	}
}

// Chain method for AdderMiddleware
func Chain(outer AdderMiddleware, middleware ...AdderMiddleware) AdderMiddleware {
	return func(a Adder) Adder {
		topIndex := len(middleware) - 1
		for i := range middleware {
			a = middleware[topIndex-i](a)
		}
		return outer(a)
	}
}

type fnString func(string) string

func compose2(a fnString, b fnString) fnString {
	return func(s string) string {
		return a(b(s))
	}
}

func compose(fns ...fnString) fnString {
	return func(s string) string {
		f := fns[0]
		fs := fns[1:]

		if len(fns) == 1 {
			return f(s)
		}
		return f(compose(fs...)(s))
	}
}

type fnUrl func(url.URL) url.URL

func composeUrl(fns ...fnUrl) fnUrl {
	return func(s url.URL) url.URL {
		f := fns[0]
		fs := fns[1:]
		if len(fns) == 1 {
			return f(s)
		}
		return f(composeUrl(fs...)(s))
	}
}

func changeScheme(s string) fnUrl {
	return func(i url.URL) url.URL {
		i.Scheme = s
		return i
	}
}

func addPathToUrl(path string) fnUrl {
	return func(i url.URL) url.URL {
		fmt.Println(i.Scheme)
		urlPath := url.URL{Path: path}
		i = *i.ResolveReference(&urlPath)
		return i
	}
}

type StringManipulator func(string) string

func ToLower(m StringManipulator) StringManipulator {
	return func(s string) string {
		lower := strings.ToLower(s)
		return m(lower)
	}
}

func ToBase64(m StringManipulator) StringManipulator {
	return func(s string) string {
		b64 := base64.StdEncoding.EncodeToString([]byte(s))
		return m(b64)
	}
}

func AppendDecorator(x string) func(m StringManipulator) StringManipulator {
	return func(m StringManipulator) StringManipulator {
		return func(s string) string {
			return m(s + x)
		}
	}
}

func PrependDecorator(x string, m StringManipulator) StringManipulator {
	return func(s string) string {
		return m(x + s)
	}
}

// "identity" just return the same string
func ident(s string) string {
	return s
}
