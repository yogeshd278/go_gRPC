---

# gRPC Server in Go

This project is a simple **gRPC server implementation in Golang**.
It demonstrates how to set up a gRPC server, define protobuf services, and test the server with a client.

---

## 📂 Project Structure

```
.
├── main.go    # gRPC server implementation
├── client
│   ├── client.go
├── protoc/            # Generated gRPC code from .proto files
│   ├── hello.pb.go
│   ├── hello_grpc.pb.go
│   └── hello.proto
├── server
│   ├── server.go
└── go.mod
```

---

## ⚡ Prerequisites

* [Go](https://go.dev/dl/) (>= 1.18)
* [protoc](https://grpc.io/docs/protoc-installation/) (Protocol Buffers compiler)
* Go plugins for protobuf and gRPC:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure `$(go env GOPATH)/bin` is in your `PATH`.

---

## 🚀 How to Run

1. **Clone the repo**

```bash
git clone https://github.com/yogeshd278/go_gRPC.git
cd go_gRPC
```

2. **Generate gRPC code**

```bash
protoc --go_out=. --go-grpc_out=. protoc/hello.proto
```

3. **Run the server**

```bash
go run main.go
```

Server starts on port **:9000**.

---

## 🧪 Testing the gRPC Service

### Option 1: Using a Go client

Run it:

```bash
go run client.go
```

Output:

```
Server Response: Hello from server, World
```

---

### Option 2: Using [grpcurl](https://github.com/fullstorydev/grpcurl)

Install grpcurl:

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

Call server:

```bash
grpcurl -plaintext -d '{"message": "World"}' localhost:9000 proto.HelloServer/ServerReply
```

## 📖 References

* [gRPC in Go](https://grpc.io/docs/languages/go/)
* [Protocol Buffers](https://developers.google.com/protocol-buffers)

---
