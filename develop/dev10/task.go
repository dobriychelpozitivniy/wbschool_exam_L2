package main

import (
	"dev10/caller"
	"dev10/connect"
	"dev10/internalReader"
	"dev10/internalWriter"
	"dev10/reader"
	"dev10/writer"
	"flag"
	"fmt"
	"net"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Int("timeout", 10, "timeout for connect to telnet server")
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)

	if host == "" || port == "" {
		panic("Command format is go-telnet --timeout='n'(optional) host port")
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	call := caller.StandardCaller

	DialToAndCall(addr, call, time.Duration(*timeout))
}

type Client struct {
	Caller caller.Caller
}

func (client *Client) Call(conn *connect.Conn) error {

	c := client.Caller
	if nil == c {
		c = caller.StandardCaller
	}

	var w writer.Writer = conn
	var r reader.Reader = conn

	c.CallTELNET(w, r)
	conn.Close()

	return nil
}

func DialToAndCall(srvAddr string, caller caller.Caller, timeout time.Duration) error {
	conn, err := DialTo(srvAddr, timeout)
	if nil != err {
		return err
	}

	client := &Client{Caller: caller}

	return client.Call(conn)
}

// Создает дефолтное tcp соединение с кастомными reader и writer,
func DialTo(addr string, timeout time.Duration) (*connect.Conn, error) {

	const network = "tcp"

	if "" == addr {
		addr = "127.0.0.1:telnet"
	}

	conn, err := net.DialTimeout(network, addr, timeout*time.Second)
	if nil != err {
		return nil, err
	}

	dataReader := internalReader.NewReader(conn)
	dataWriter := internalWriter.NewWriter(conn)

	clientConn := connect.Conn{
		Conn:       conn,
		DataReader: dataReader,
		DataWriter: dataWriter,
	}

	return &clientConn, nil
}
