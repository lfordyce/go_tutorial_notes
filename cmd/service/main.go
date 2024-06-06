package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/lfordyce/generalNotes/cmd"
	"github.com/lfordyce/generalNotes/cmd/service/monster"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	ac := make(chan server, 1)
	setupServer(ac)
	ps := <-ac

	ctx, cancel := context.WithCancel(context.Background())
	if err := handleServer(ps)(ctx, cancel); err != nil {
		panic(err)
	}
}

func setupServer(ac chan server) {
	handler := new(monster.Handler)

	router := http.NewServeMux()
	//router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Received a non domain request")
	//	if _, err := w.Write([]byte("Hello, stranger...")); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//	}
	//}))

	// TEST:
	// curl -v -i -X POST -H "Content-Type: multipart/form-data" -F "file=@MOCK_DATA.csv" http://localhost:8080/v1/upload
	router.Handle("POST /upload", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Maximum upload of 10 MB files
		req.ParseMultipartForm(10 << 20)

		// Get handler for filename, size and headers
		file, header, err := req.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", header.Filename)
		fmt.Printf("File Size: %+v\n", header.Size)
		fmt.Printf("MIME Header: %+v\n", header.Header)
	}))
	//
	router.Handle("POST /csv", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		file, header, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		fmt.Printf("File name %s\n", name[0])
		// Copy the file data to my buffer
		io.Copy(&buf, file)
		// do something with the contents...
		// I normally have a struct defined and unmarshal into a struct, but this will
		// work as an example
		contents := buf.String()
		fmt.Println(contents)
		// I reset the buffer in case I want to use it again
		// reduces memory allocations in more intense projects
		buf.Reset()
		// do something else
		// etc write header
		return

		//m, p, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
		//if err != nil {
		//	http.Error(rw, err.Error(), http.StatusBadRequest)
		//	return
		//}
		//boundary := p["boundary"]
		//reader := multipart.NewReader(req.Body, boundary)
		//for {
		//	part, err := reader.NextPart()
		//	if err == io.EOF {
		//		// Done reading body
		//		break
		//	}
		//	//contentType:=part.Header.Get("Content-Type")
		//	//fname:=part.FileName()
		//	// part is an io.Reader, deal with it
		//}

		//r := csv.NewReader(req.Body)
		//for {
		//	record, err := r.Read()
		//	if err == io.EOF {
		//		break
		//	}
		//	if err != nil {
		//		http.Error(rw, err.Error(), http.StatusInternalServerError)
		//		return
		//	}
		//	fmt.Println(record)
		//}
	}))

	router.Handle("GET /echo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received echo request")
		if _, err := io.Copy(w, r.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	router.HandleFunc("POST /monster", handler.Create)
	router.HandleFunc("PUT /monster/{id}", handler.UpdateByID)
	router.HandleFunc("GET /monster/{id}", handler.FindByID)
	router.HandleFunc("DELETE /monster/{id}", handler.DeleteByID)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	stack := CreateStack(Logging, CheckPermissions)

	svc := &http.Server{
		Addr:    ":8080",
		Handler: stack(v1),
	}
	ac <- svc
}

type Start func(ctx context.Context, cancel context.CancelFunc) error

func handleServer(s server) Start {
	return func(ctx context.Context, cancel context.CancelFunc) error {
		exitCh := make(chan error)

		go cmd.WaitForTerminate(ctx, func(msg string) {
			log.Printf("terminate signal: %s\n", msg)
			exitCh <- s.Shutdown(ctx)
		})

		defer cancel()
		serverErr := s.ListenAndServe()

		// When Shutdown is called, Serve, ListenAndServe, and
		// ListenAndServeTLS immediately return ErrServerClosed. Make sure the
		// program doesn't exit and waits instead for Shutdown to return.
		if errors.Is(serverErr, http.ErrServerClosed) {
			// waits instead for Shutdown to return.
			if err := <-exitCh; err != nil {
				log.Printf("shutdown error: %v\n", err)
			}
			// on shutdown, no error or error already logged
			return nil
		}
		return serverErr
	}
}
