// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/wgliang/gops/agent"
)

func main() {
	if err := agent.Listen(&agent.Options{
		Addr: "127.0.0.1:4321",
	}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
