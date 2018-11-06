# TCP Echo Server

[![Docker Pulls](https://img.shields.io/docker/pulls/venilnoronha/tcp-echo-server.svg?logo=docker)](https://hub.docker.com/r/venilnoronha/tcp-echo-server/)
[![License](https://img.shields.io/github/license/venilnoronha/tcp-echo-server.svg)](LICENSE)

TCP Echo Server is a simple echo server written in Golang that works plainly
over the TCP protocol. The server listens to connection requests, prepends a
preconfigured prefix to the data it receives from a client and writes it back
on the connection.

## Deploy Locally

To start the TCP Echo Server locally, simply run the following command.

```console
$ go run -v main.go 9000 hello
command-line-arguments
listening on [::]:9000, prefix: hello
```

## Deploy Over Docker

To run the TCP Echo Server container, simply execute the following command.

```console
$ docker run -p 9000:9000 venilnoronha/tcp-echo-server:latest
listening on [::]:9000, prefix: hello
```

## Manual Test

To test the TCP Echo Server, run the following command.

```console
$ echo world | nc localhost 9000
hello world
```

## Unit Test

To run the unit test, execute the following command.

```console
$ go test -v ./...
=== RUN   TestTcpEchoServer
listening on [::]:9000, prefix: hello
request: world
response: hello world
--- PASS: TestTcpEchoServer (2.01s)
PASS
ok  	github.com/venilnoronha/tcp-echo-server	2.015s
```

## License

The TCP Echo Server project is distributed under the BSD-style license found in
the [LICENSE](LICENSE) file.
