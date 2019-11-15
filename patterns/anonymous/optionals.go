package anonymous

import (
	"fmt"
	"net"
	"net/url"
)

const (
	defaultScheme = "http"
	defaultPort   = "8080"
)

func createURL(mgmtURI string, serviceId string, transformFunc TransformationTypeFunc) string {

	var urlToUse *url.URL

	urlToUse, err := url.ParseRequestURI(mgmtURI)
	if err != nil {
		urlToUse = &url.URL{Host: mgmtURI}
	}

	apply := transformFunc.Apply(*urlToUse)(defaultScheme, serviceId)
	return apply.String()
}

var (
	EncoderURL1  = TransformationTypeFunc(encoderStrategy)
	PackagerURL1 = TransformationTypeFunc(packagerStrategy)
)

type TransformationTypeFunc func(url.URL, ...string) url.URL

//type Strategic interface {
//	SetTransformation(TransformationTypeFunc)
//	Result() string
//}

//type createStrategy struct {
//	parsedUrl *url.URL
//	mgmtURI string
//	serviceId string
//	transformFunc TransformationTypeFunc
//	result string
//}
//
//func (c createStrategy) SetTransformation(t TransformationTypeFunc) {
//	c.transformFunc = t
//}
//
//func (c createStrategy) Result() string {
//	return c.transformFunc.Apply(*c.parsedUrl)(defaultScheme, c.serviceId).String()
//}

func (h TransformationTypeFunc) Apply(i url.URL) func(...string) url.URL {
	return func(values ...string) url.URL {
		values = append([]string{}, values...)
		return h(i, values...)
	}
}

func encoderStrategy(i url.URL, s ...string) url.URL {
	i.Scheme = s[0]
	return i
}

func packagerStrategy(i url.URL, s ...string) url.URL {
	i.Scheme = s[0]

	port := i.Port()

	if port == "" {
		port = net.JoinHostPort(i.Host, defaultPort)
		i.Host = port
	}

	if len(s) > 1 && s[1] != "" {

		serviceID := s[1]

		path := url.URL{Path: fmt.Sprintf("/Halo/service/%v", serviceID)}
		i = *i.ResolveReference(&path)
		return i
	}
	return i
}

type TransformOption interface {
	apply(*transformOptions)
}

type funcOptions struct {
	f func(*transformOptions)
}

func (fdo *funcOptions) apply(do *transformOptions) {
	fdo.f(do)
}

func newFuncOption(f func(*transformOptions)) *funcOptions {
	return &funcOptions{
		f: f,
	}
}

type SchemeChange func(parse url.URL, scheme string) url.URL

type PackagerTransform func(parse url.URL, update SchemeChange, serviceId string) url.URL

type transformOptions struct {
	schemeChane SchemeChange
}

func WithScheme(f SchemeChange) TransformOption {
	return newFuncOption(func(options *transformOptions) {
		options.schemeChane = f
	})
}

func schemeChange(parse url.URL, scheme string) url.URL {
	parse.Scheme = scheme
	return parse
}

func Create(mgmtURI string, opts ...TransformOption) *createTransform {

	var uri *url.URL

	uri, err := url.ParseRequestURI(mgmtURI)
	if err != nil {
		uri = &url.URL{Host: mgmtURI}
	}

	transform := &createTransform{
		parser: uri,
	}

	for _, opt := range opts {
		opt.apply(&transform.opts)
	}
	return transform
}

type createTransform struct {
	parser *url.URL
	opts   transformOptions
}

func (c *createTransform) Execute() string {
	options := c.opts
	chane := options.schemeChane(*c.parser, "http")
	return chane.String()
}
