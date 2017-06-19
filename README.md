A simple client in go serving a ruby server through gRPC.

Data are read from a local file and stream into a unix socket to the go code, that will transmit it to the gRPC ruby server.

The goal is to reproduce and track the issue describe here: https://github.com/grpc/grpc/issues/11331

To make work the test do
- `bundle install`
- start ruby server gRPC with `bundle exec bin/my_agent.rb`
- start go client gRPC with `go run main.go`
