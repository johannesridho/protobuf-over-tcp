# protobuf-over-tcp
A simple client-server app to send protobuf message over TCP

### How to Run

#### Server

```
TYPE=server go run github.com/johannesridho/protobuf-over-tcp
Server will send text `Hello World` to each client connected
```

#### Client

```
TYPE=client go run github.com/johannesridho/protobuf-over-tcp
Client will get a text from server and then print it
```
