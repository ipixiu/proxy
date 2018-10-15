package main

import (
	"bytes"
	"errors"
)

import (
	"github.com/AlexStocks/getty"
	log "github.com/AlexStocks/log4go"
)

type ProxyGettyBuffer struct {
}

func NewProxyGettyBuffer() *ProxyGettyBuffer {
	return &ProxyGettyBuffer{}
}

func (b *ProxyGettyBuffer) Read(ss getty.Session, data []byte) (interface{}, int, error) {
	var (
		err error
		len int
		pkg Protocol
		buf *bytes.Buffer
	)

	buf = bytes.NewBuffer(data)
	len, err = pkg.Unmarshal(buf)
	if err != nil {
		if err == ErrNotEnoughStream {
			return nil, 0, nil
		}

		return nil, 0, err
	}

	return &pkg, len, nil
}

func (b *ProxyGettyBuffer) Write(ss getty.Session, pkg interface{}) error {
	var (
		err error
		buf *bytes.Buffer
		msg *Protocol
		ok  bool
	)

	if msg, ok = pkg.(*Protocol); !ok {
		log.Error("illegal pkg:%+v\n", pkg)
		return errors.New("invalid echo package!")
	}

	buf, err = msg.Marshal()
	if err != nil {
		log.Warn("binary.Write(echoPkg{%#v}) = err{%#v}", pkg, err)
		return err
	}

	err = ss.WriteBytes(buf.Bytes())

	return err
}
