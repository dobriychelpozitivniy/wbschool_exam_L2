package writer

import "io"

type Writer interface {
	Write([]byte) (int, error)
}

func LongWrite(w io.Writer, p []byte) (int64, error) {

	numWritten := int64(0)
	for {
		n, err := w.Write(p)
		numWritten += int64(n)
		if nil != err && io.ErrShortWrite != err {
			return numWritten, err
		}

		if !(n < len(p)) {
			break
		}

		p = p[n:]

		if len(p) < 1 {
			break
		}
	}

	return numWritten, nil
}
