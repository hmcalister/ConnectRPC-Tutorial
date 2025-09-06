# Connect RPC Tutorial

Just a quick follow through the [ConnectRPC](https://connectrpc.com/) tutorial for Golang, which can be found [here](https://connectrpc.com/docs/go/getting-started/).

ConnectRPC is a library for making APIs that are callable in both browsers and through gRPC clients, which seems ideal for multi-client apps. For example, the same server can serve both web clients (e.g. a webapp frontend) and programmatic clients (terminal UIs, other programs, etc...) while retaining the nice type enforcement of RPC compared to HTTP.