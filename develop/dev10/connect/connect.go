package connect

import (
	"dev10/internalReader"
	"dev10/internalWriter"
	"net"
)

type Conn struct {
	Conn interface {
		Read(b []byte) (n int, err error)
		Write(b []byte) (n int, err error)
		Close() error
		LocalAddr() net.Addr
		RemoteAddr() net.Addr
	}
	DataReader *internalReader.InternalReader
	DataWriter *internalWriter.InternalWriter
}

// Close closes the client connection.
//
// Typical usage might look like:
//
//	telnetsClient, err = telnet.DialToTLS(addr, tlsConfig)
//	if nil != err {
//		//@TODO: Handle error.
//		return err
//	}
//	defer telnetsClient.Close()
func (clientConn *Conn) Close() error {
	return clientConn.Conn.Close()
}

// Read receives `n` bytes sent from the server to the client,
// and "returns" into `p`.
//
// Note that Read can only be used for receiving TELNET (and TELNETS) data from the server.
//
// TELNET (and TELNETS) command codes cannot be received using this method, as Read deals
// with TELNET (and TELNETS) "unescaping", and (when appropriate) filters out TELNET (and TELNETS)
// command codes.
//
// Read makes Client fit the io.Reader interface.
func (clientConn *Conn) Read(p []byte) (n int, err error) {
	return clientConn.DataReader.Read(p)
}

// Write sends `n` bytes from 'p' to the server.
//
// Note that Write can only be used for sending TELNET (and TELNETS) data to the server.
//
// TELNET (and TELNETS) command codes cannot be sent using this method, as Write deals with
// TELNET (and TELNETS) "escaping", and will properly "escape" anything written with it.
//
// Write makes Conn fit the io.Writer interface.
func (clientConn *Conn) Write(p []byte) (n int, err error) {
	return clientConn.DataWriter.Write(p)
}

// LocalAddr returns the local network address.
func (clientConn *Conn) LocalAddr() net.Addr {
	return clientConn.Conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (clientConn *Conn) RemoteAddr() net.Addr {
	return clientConn.Conn.RemoteAddr()
}
