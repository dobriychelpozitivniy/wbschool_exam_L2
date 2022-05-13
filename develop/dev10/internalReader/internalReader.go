package internalReader

import (
	"bufio"
	"fmt"
	"io"
)

type InternalReader struct {
	wrapped  io.Reader
	buffered *bufio.Reader
}

func NewReader(r io.Reader) *InternalReader {
	buffered := bufio.NewReader(r)

	reader := InternalReader{
		wrapped:  r,
		buffered: buffered,
	}

	return &reader
}

func (r *InternalReader) Read(data []byte) (n int, err error) {

	const IAC = 255

	const SB = 250
	const SE = 240

	const WILL = 251
	const WONT = 252
	const DO = 253
	const DONT = 254

	p := data

	for len(p) > 0 {
		var b byte

		b, err = r.buffered.ReadByte()
		if nil != err {
			return n, err
		}

		if IAC == b {
			var peeked []byte

			peeked, err = r.buffered.Peek(1)
			if nil != err {
				return n, err
			}

			switch peeked[0] {
			case WILL, WONT, DO, DONT:
				_, err = r.buffered.Discard(2)
				if nil != err {
					return n, err
				}
			case IAC:
				p[0] = IAC
				n++
				p = p[1:]

				_, err = r.buffered.Discard(1)
				if nil != err {
					return n, err
				}
			case SB:
				for {
					var b2 byte
					b2, err = r.buffered.ReadByte()
					if nil != err {
						return n, err
					}

					if IAC == b2 {
						peeked, err = r.buffered.Peek(1)
						if nil != err {
							return n, err
						}

						if IAC == peeked[0] {
							_, err = r.buffered.Discard(1)
							if nil != err {
								return n, err
							}
						}

						if SE == peeked[0] {
							_, err = r.buffered.Discard(1)
							if nil != err {
								return n, err
							}
							break
						}
					}
				}
			case SE:
				_, err = r.buffered.Discard(1)
				if nil != err {
					return n, err
				}
			default:
				// If we get in here, this is not following the TELNET protocol.
				//@TODO: Make a better error.
				err = fmt.Errorf("I dont now what is this error")
				return n, err
			}
		} else {
			p[0] = b
			n++
			p = p[1:]
		}
	}

	return n, nil
}
