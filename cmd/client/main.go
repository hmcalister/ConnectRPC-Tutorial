package main

import (
	"context"
	"log"
	"net/http"

	greetv1 "hmcalister/connectrpcTutorial/gen/greet/v1"
	"hmcalister/connectrpcTutorial/gen/greet/v1/greetv1connect"

	"connectrpc.com/connect"
)

func main() {
	// Make a new client with all the same magic as `greetv1connect.NewGreetServiceHandler`
	// Note the client (and server) are proto specific, since the protobuf definition is so rigorous.
	//
	// We use the http.DefaultClient as a base, and we specify the *base* api, that is
	// *without* the additional path `/greetgreet.v1.GreetService`
	//
	// There are many more options to use in this method, including using gRPC and so on.
	// Check out the documentation (generated) or the connect documentation for more.
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		// connect.WithGRPC(),
	)

	// client.Greet implements the method defined in the protobuf, but from the client side this time.
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
