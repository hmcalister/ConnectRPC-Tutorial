# Connect RPC Tutorial

Just a quick follow through the [ConnectRPC](https://connectrpc.com/) tutorial for Golang, which can be found [here](https://connectrpc.com/docs/go/getting-started/).

ConnectRPC is a library for making APIs that are callable in both browsers and through gRPC clients, which seems ideal for multi-client apps. For example, the same server can serve both web clients (e.g. a webapp frontend) and programmatic clients (terminal UIs, other programs, etc...) while retaining the nice type enforcement of RPC compared to HTTP.

### Define the Protocol Buffer

These schemas are shared between client and server, so need to be in a shared location. Here, we use `greet/v1`, which must match the exported package name. See `greet/v1/greet.proto` for an explanation of the proto format and some good resources on learning more.

### Generate Code

Once the proto file is ready, use [Buf](https://buf.build/) to generate code from the proto file. Create an initial config using 
```bash
buf config init
```
then add project-specific configurations to a new file called `buf.gen.yaml`. In this example, since we are targeting Golang, we can use the simple file:
```yaml
version: v2
plugins:
  - local: protoc-gen-go
    out: gen
    opt: paths=source_relative
  - local: protoc-gen-connect-go
    out: gen
    opt: paths=source_relative
```
Run `buf lint` to check for errors. Then `buf generate` to make the code.

That's all the generation done for us! The generated files are a little unhappy to the go linter, but ignore that. 

### Implement Server Logic

Inside `cmd/server/main.go` implement the server logic to listen for incoming requests and server responses. We use `greetv1connect.NewGreetServiceHandler` to easily create a handler from the specific implementor of the protobuf methods.

### Run and Request the Server

Run `go run cmd/server/main.go` to start serving the API on `localhost:8080`, with the API specific methods bound under the `/greet.v1.GreetService/` subpath. We can easily request the server using curl, for example:
```bash
curl \
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
    http://localhost:8080/greet.v1.GreetService/Greet
```
or with gRPC:
```bash
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"name": "Jane"}' \
    localhost:8080 greet.v1.GreetService/Greet
```

### Making a Client

We can also make a full fleshed client in Go. This time, in `cmd/client/main.go`, we use `greetv1connect.NewGreetServiceClient` to make a client that will send `GreetRequest` to the server.