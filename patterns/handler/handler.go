package handler

import "fmt"

//ExampleWriter is the object that will be the result of the chain, implements writer
type ExampleWriter struct {
	History []string
}

func (e *ExampleWriter) Write(data []byte) (int, error) {
	e.History = append(e.History, string(data))
	return len(data), nil
}

// ExampleInput helper object for the input data in handlerfuncs, could be a
// simple string or `type ExampleInput string`
type ExampleInput struct {
	Data string
}

// ExampleHandler is the interface that will be implemeted by the handlerFunc type
type ExampleHandler interface {
	RunExample(*ExampleWriter, *ExampleInput)
}

// ExampleHandlerFunc type will be a function compatible with RunExample so
// it can execute itself, in other words, it implements ExampleHandler interface
// by executing itself
type ExampleHandlerFunc func(*ExampleWriter, *ExampleInput)

// RunExample will execute itself (remember this object is a function)
func (e ExampleHandlerFunc) RunExample(ew *ExampleWriter, in *ExampleInput) {
	e(ew, in)
}

// historyExample will add history to our handlers
func historyIDExample(id int, eh ExampleHandler) ExampleHandler {
	return ExampleHandlerFunc(func(ew *ExampleWriter, in *ExampleInput) {
		fmt.Fprintf(ew, "Start example %d: %s", id, in.Data)
		defer fmt.Fprintf(ew, "Finish example %d: %s", id, in.Data) // We could use defer!
		eh.RunExample(ew, in)
	})
}

func strExample(start, end string, eh ExampleHandler) ExampleHandler {
	return ExampleHandlerFunc(func(ew *ExampleWriter, in *ExampleInput) {
		if start != "" {
			fmt.Fprint(ew, start)
		}
		if end != "" {
			defer fmt.Fprint(ew, end)
		}
		eh.RunExample(ew, in)
	})
}
