package internalWriter

import (
	"bytes"
	"dev10/writer"
	"errors"
	"io"
)

type InternalWriter struct {
	wrapped io.Writer
}

func NewWriter(w io.Writer) *InternalWriter {
	writer := InternalWriter{
		wrapped: w,
	}

	return &writer
}

var iaciac []byte = []byte{255, 255}

var errOverflow = errors.New("Overflow")
var errPartialIACIACWrite = errors.New("Partial IAC IAC write.")

func (w *InternalWriter) Write(data []byte) (n int, err error) {
	var n64 int64

	n64, err = w.write64(data)
	n = int(n64)
	if int64(n) != n64 {
		panic(errOverflow)
	}

	return n, err
}

func (w *InternalWriter) write64(data []byte) (n int64, err error) {

	if len(data) <= 0 {
		return 0, nil
	}

	const IAC = 255

	var buffer bytes.Buffer
	for _, datum := range data {

		if IAC == datum {

			if buffer.Len() > 0 {
				var numWritten int64

				numWritten, err = writer.LongWrite(w.wrapped, buffer.Bytes())
				n += numWritten
				if nil != err {
					return n, err
				}
				buffer.Reset()
			}

			var numWritten int64
			//@TODO: Should we worry about "iaciac" potentially being modified by the .Write()?
			numWritten, err = writer.LongWrite(w.wrapped, iaciac)
			if int64(len(iaciac)) != numWritten {
				//@TODO: Do we really want to panic() here?
				panic(errPartialIACIACWrite)
			}
			n += 1
			if nil != err {
				return n, err
			}
		} else {
			buffer.WriteByte(datum) // The returned error is always nil, so we ignore it.
		}
	}

	if buffer.Len() > 0 {
		var numWritten int64
		numWritten, err = writer.LongWrite(w.wrapped, buffer.Bytes())
		n += numWritten
	}

	return n, err
}
