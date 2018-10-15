package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	PROTOCAL_HEAD_LEN = 4
	MAX_PROTOCAL_LEN  = 60000
)

var (
	ErrNotEnoughStream = errors.New("packet stream is not enough")
)

type ProtocolHead struct {
	Len  uint16
	Type uint16
}

type Protocol struct {
	//网络头
	Head ProtocolHead

	//数据
	Data []byte
}

func (p *Protocol) Marshal() (*bytes.Buffer, error) {
	var (
		err error
		buf *bytes.Buffer
	)

	buf = &bytes.Buffer{}
	err = binary.Write(buf, binary.LittleEndian, p.Head) //是否有按1对齐的待确定
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (p *Protocol) Unmarshal(buf *bytes.Buffer) (int, error) {
	var (
		err error
	)

	if buf.Len() < PROTOCAL_HEAD_LEN {
		return 0, ErrNotEnoughStream
	}

	err = binary.Read(buf, binary.LittleEndian, &p)
	if err != nil {
		return 0, err
	}

	if MAX_PROTOCAL_LEN < p.Head.Len {
		return 0, errors.New(fmt.Sprintf("packet len exceed the max len, maxLen=%d, msgLen=%d", MAX_PROTOCAL_LEN, p.Head.Len))
	}
	if buf.Len() < (int)(p.Head.Len) {
		return 0, ErrNotEnoughStream
	}
	p.Data = buf.Next(int(buf.Len()))

	return int(PROTOCAL_HEAD_LEN + p.Head.Len), nil
}
