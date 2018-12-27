package decorator

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
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
