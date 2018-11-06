// Copyright 2018 Venil Noronha. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

// TestTcpEchoServer tests the behavior of the TCP Echo Server.
func TestTcpEchoServer(t *testing.T) {
	// set up test parameters
	prefix := "hello"
	request := "world"
	want := prefix + " " + request

	// start the TCP Echo Server
	os.Args = []string{"main", "9000", prefix}
	go main()

	// wait for the TCP Echo Server to start
	time.Sleep(2 * time.Second)

	// connect to the TCP Echo Server
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		t.Fatalf("couldn't connect to the server: %v", err)
	}
	defer conn.Close()

	// test the TCP Echo Server output
	if _, err := conn.Write([]byte(request + "\n")); err != nil {
		t.Fatalf("couldn't send request: %v", err)
	} else {
		reader := bufio.NewReader(conn)
		if response, err := reader.ReadBytes(byte('\n')); err != nil {
			t.Fatalf("couldn't read server response: %v", err)
		} else if !strings.HasPrefix(string(response), want) {
			t.Errorf("output doesn't match, wanted: %s, got: %s", want, response)
		}
	}
}
