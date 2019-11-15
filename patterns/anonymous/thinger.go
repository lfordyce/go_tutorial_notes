package anonymous

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

type AbstractThing interface {
	DoThing()
}

type DoThingWith func()

// Satisfy AbstractThing interface.
// So we can now pass an anonymous function using DoThingWith,
// which implements AbstractThing
func (thing DoThingWith) DoThing() {
	thing()
}

type App struct {
}

func (a App) DoThing(f AbstractThing) {
	f.DoThing()
}

type I interface {
	M()
}

type T1 struct {
	field1 string
}

func (t *T1) M() {
	t.field1 = t.field1 + t.field1
}

type T2 struct {
	field2 string
	T1
}

type bindFunc func(int, int) int

func add(x, y int) int { return x + y }

func (f bindFunc) Error() string {
	return "bindFunc error"
}

type Add func(int, int) int

func (f Add) Apply(i int) func(int) int {
	return func(j int) int {
		return f(i, j)
	}
}

type Option func(Add) Add

func applyOptions(add Add, opts ...Option) Add {
	for _, opt := range opts {
		add = opt(add)
	}
	return add
}

func Third(n int) Option {
	return func(i Add) Add {
		n := n
		return func(i2 int, i int) int {
			return n + i2 + i
		}
	}
}

var adding Add = func(i int, j int) int {
	return i + j
}

type Multiply func(...int) int

func (f Multiply) Apply(i int) func(...int) int {
	return func(values ...int) int {
		values = append([]int{i}, values...)
		return f(values...)
	}
}

var multiply Multiply = func(values ...int) int {
	var total = 1
	for _, value := range values {
		total *= value
	}
	return total
}

func trimIndexRune(s string) string {
	if idx := strings.IndexRune(s, ':'); idx != -1 {
		return s[idx:]
	}
	return s
}

func trimStringFromSlash(s string) string {
	if idx := strings.Index(s, "://"); idx != -1 {
		return s[idx:]
	}
	return s
}

type URLOptions func(*url.URL, ...string) *url.URL

func (h URLOptions) Apply(i *url.URL) func(...string) *url.URL {

	return func(values ...string) *url.URL {

		values = append([]string{}, values...)

		return h(i, values...)
	}
}

func setScheme(i *url.URL, s ...string) *url.URL {
	i.Scheme = s[0]
	return i
}

func setSchemeWithOptions(i *url.URL, s ...string) *url.URL {
	i.Scheme = s[0]

	port := i.Port()

	if port == "" {
		port = net.JoinHostPort(i.Host, "8080")
		i.Host = port
	}

	if len(s) > 1 {

		serviceID := s[1]

		path := url.URL{Path: fmt.Sprintf("/Halo/service/%v", serviceID)}

		i = i.ResolveReference(&path)
		return i
	}

	return i
}

var (
	EncoderURL  = URLOptions(setScheme)
	PackagerURL = URLOptions(setSchemeWithOptions)
)
