package channel

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"simusco.com/go-netty/util"
	"sync"
)

type Channel struct {
	Id int64

	attrs     map[string]string
	conn      net.Conn
	writeChan chan []byte
	lock      sync.Mutex
	closeFlag bool
}

func NewChannel(c net.Conn) *Channel {
	ch := Channel{
		Id:   util.GetRand(1000000000, 9999999999),
		conn: c,
	}
	return &ch
}

func (c *Channel) Attr(key string, value string) {
	c.attrs[key] = value
}

func (c *Channel) Write(p []byte) (n int, err error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(len(p)))

	buf := new(bytes.Buffer)
	buf.Write(b)
	buf.Write(p)

	return c.conn.Write(buf.Bytes())
}

func (c *Channel) Read(p []byte) (n int, err error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return -1, err
	}

	msgLen := binary.BigEndian.Uint32(buf)

	if msgLen > 4096 {
		return -1, errors.New("message to long")
	}

	msg := make([]byte, msgLen)
	if _, err := io.ReadFull(c.conn, msg); err != nil {
		return -1, err
	}

	p = msg

	return -1, nil
}

func (c *Channel) Close() {
	_ = c.conn.Close()
}
