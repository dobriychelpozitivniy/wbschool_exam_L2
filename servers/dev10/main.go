package main

import (
	"fmt"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

func main() {

	var handler telnet.Handler = internalEchoHandler{}

	err := telnet.ListenAndServe(":5555", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}

type internalEchoHandler struct{}

func (handler internalEchoHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {

	var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	p := buffer[:]

	for {
		n, err := r.Read(p)

		if n > 0 {
			fmt.Println("Read from client: ", string(p[:n]))
			oi.LongWrite(w, p[:n])
		}

		if nil != err {
			break
		}
	}
}
