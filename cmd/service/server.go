package main

import "context"

type server interface {
	Shutdown(context.Context) error
	ListenAndServe() error
}
