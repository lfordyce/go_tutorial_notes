package handler

import (
	"fmt"
	"io"
	"testing"
)

func TestExampleHandlerFunc_RunExample(t *testing.T) {
	end := ExampleHandlerFunc(func(io.Writer, *ExampleInput) {
		fmt.Println("this is the end...")
	})
	he := strExample("Starting...", "End!",
		historyIDExample(1,
			historyIDExample(2,
				historyIDExample(3,
					historyIDExample(4,
						historyIDExample(5,
							historyIDExample(6, end),
						),
					),
				),
			),
		),
	)

	strExample("start", "finish", end)

	w := ExampleWriter{History: []string{}}
	he.RunExample(&w, &ExampleInput{Data: "This is an example"})
	for _, h := range w.History {
		fmt.Println(h)
	}
}
