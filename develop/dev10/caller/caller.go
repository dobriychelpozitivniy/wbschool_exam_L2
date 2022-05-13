package caller

import (
	"bufio"
	"bytes"
	"dev10/reader"
	"dev10/writer"
	"fmt"
	"io"
	"os"
	"time"
)

type Caller interface {
	CallTELNET(writer.Writer, reader.Reader)
}

var StandardCaller Caller = internalStandardCaller{}

type internalStandardCaller struct{}

func (caller internalStandardCaller) CallTELNET(w writer.Writer, r reader.Reader) {
	standardCallerCallTELNET(os.Stdin, os.Stdout, os.Stderr, w, r)
}

func standardCallerCallTELNET(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, w writer.Writer, r reader.Reader) {

	go func(wr io.Writer, re io.Reader) {

		var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
		p := buffer[:]

		for {
			// Read 1 byte.
			n, err := re.Read(p)
			if n <= 0 && nil == err {
				continue
			} else if n <= 0 && nil != err {
				break
			}

			writer.LongWrite(wr, p)
		}
	}(stdout, r)

	var buffer bytes.Buffer
	var p []byte

	var crlfBuffer [2]byte = [2]byte{'\r', '\n'}
	crlf := crlfBuffer[:]

	scanner := bufio.NewScanner(stdin)
	scanner.Split(scannerSplitFunc)

	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		buffer.Write(crlf)

		p = buffer.Bytes()

		n, err := writer.LongWrite(w, p)
		if nil != err {
			break
		}
		if expected, actual := int64(len(p)), n; expected != actual {
			err := fmt.Errorf("Transmission problem: tried sending %d bytes, but actually only sent %d bytes.", expected, actual)
			fmt.Fprint(stderr, err.Error())
			return
		}

		buffer.Reset()
	}

	// Wait a bit to receive data from the server (that we would send to io.Stdout).
	time.Sleep(3 * time.Millisecond)
}

func scannerSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, nil
	}

	return bufio.ScanLines(data, atEOF)
}
