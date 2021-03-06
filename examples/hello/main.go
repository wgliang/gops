// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/wgliang/gops/agent"
)

func main() {
	addr := "127.0.0.1:7890"

	var mux = http.NewServeMux()

	go func(serverAddr string, m *http.ServeMux) {
		if err := http.ListenAndServe(serverAddr, m); err != nil {
			log.Fatalln(`Binding Ip and Port Err, Please check whether port is occupied:`, err)
		}
	}(addr, mux)

	if err := agent.Listen(&agent.Options{
		Addr:            addr,
		EnableProfiling: true,
		ProfilingMux:    mux,
	}); err != nil {
		log.Fatal(err)
	}

	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	select {
	case <-chExit:
		log.Println("EXIT...Bye.")
	}
}
