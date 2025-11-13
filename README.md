# Go HttpServer

- Resource: https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
- This is a HttpServer written in Go
- Ensure Go is installed on your machine to run
- Cmd to run server: `go run main.go`

- Go HTTP Server has two major components: server that listens for requests coming from HTTP clients, and one or more request handlers that respond to the requests coming in
  - `http.HandlerFunc` in `main.go` tells the server which function to call to handle a request to the server
    - This is used for `getRoot` and `getHello` functions
    - Both take `http.ResponseWriter` and `*http.Request` value
      - When request is made to server, it sets up the two values with info about the request made, then calls the handler function with those values
      - `http.ResponseWriter` value is used to control the response info being written back to the client that made the request
        - This can include response body or status code
      - `*http.Request` value is used to get info about the incoming request-- such as body or info about the client making the request
    - this gets called in `main` function for each route we want to define
  - `http.ListenAndServe` starts the server and tells it to listen for new HTTP requests, and then serve them using the handler functions
    - Tells the global HTTP server to listen for incoming requests on a specific port
      - We specify port as `:3333`, but without this, the server will listen on every IP address associated with your computer
    - Also passes `nil` value for `http.Handler` parameter
      - This tells `ListenAndServe` function you want to use default server multiplexer, not the one you've setup
        - Multiplexer handles things like routing, 3rd party deps, req/resp handling
    - `ListenAndServe` is a blocking call, programs won't continue running until AFTER `ListenAndServe` finishes running
      - HOWEVER this won't finish running until your program finishes running, OR the HTTP server is told to 
    - `ListenAndServe` needs error handling as calling the function can fail
  - `http.ErrServerClosed` is returned when the server is told to shu down or close
    - Is usually an expected error because you'll be shutting the server yourself
      - Can also be used to show why the server stopped in the output
    - Errors NOT related to this can include things like not being able to listen on the specified port, can happen if port is commonly used/is being used by another program on your machine
    - 