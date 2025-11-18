package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

// acts as key for HTTP server's address value
// shown in http.Request context
const keyServerAddr = "serverAddr"

// both getRoot and getHello look at context for keyServerAddr
// include that in printed value so you can see which server handled the HTTP request
// root path calls getRoot handler function
func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// returns bool if first query param present
	hasFirst := r.URL.Query().Has("first")
	// gets string of first query param
	// returns empty string if not found
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	// reads the request body until it gets error
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("couldn't read req body: %s\n", err)
	}

	// fmt.Printf("Got root request\n", ctx.Value(keyServerAddr))
	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s\n, body:\n%s\n",
		ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second,
		body,
	)
	io.WriteString(w, "This is my website!\n")
}

// /hello path calls getHello handler fucntion
func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))

	// checks FormData for myName property
	myName := r.PostFormValue("myName")
	if myName == "" {
		myName = "HTTP"
	}

	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
}

func main() {
	// custom multiplexer
	mux := http.NewServeMux()
	// each HandleFunc call sets up a specific request path
	// can use http. or mux., depending on whether you want to use default multiplexer
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf(("Error listening for server: %s\n"))
	}

	// // ctx is the new context, cancelCtx cancels that context
	// ctx, cancelCtx := context.WithCancel(context.Background())
	// // pass server definition with http.Server instead of http.ListenAndServe
	// serverOne := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// 	BaseContext: func(l net.Listener) context.Context {
	// 		// add the address that the server is listening on with l.Addr().String() to the context
	// 		// also uses key defined as keyServerAddr
	// 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
	// 		return ctx
	// 	},
	// }

	// serverTwo := &http.Server{
	// 	Addr:    ":8081",
	// 	Handler: mux,
	// 	BaseContext: func(l net.Listener) context.Context {
	// 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
	// 		return ctx
	// 	},
	// }

	// // starts first server in goroutine
	// go func() {
	// 	// starts and listens on our new server
	// 	// same error handling as below
	// 	err := serverOne.ListenAndServe()
	// 	if errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Printf("server one closed \n")
	// 	} else if err != nil {
	// 		fmt.Printf("error listening on server one: %s\n", err)
	// 	}
	// 	cancelCtx()
	// }()

	// // starts second server in goroutine
	// go func() {
	// 	// starts and listens on our new server
	// 	// same error handling as below
	// 	err := serverTwo.ListenAndServe()
	// 	if errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Printf("server two closed \n")
	// 	} else if err != nil {
	// 		fmt.Printf("error listening on server two: %s\n", err)
	// 	}
	// 	cancelCtx()
	// }()

	<-ctx.Done()

	// err := http.ListenAndServe(":8080", mux)

	// we pass nil to use default multiplexer
	// good if you need a basic handler that calls a single func with a specific req path
	// err := http.ListenAndServe(":8080", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("Server closed\n")
	// 	// checks for any non-ErrServerClosed error
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }
}
