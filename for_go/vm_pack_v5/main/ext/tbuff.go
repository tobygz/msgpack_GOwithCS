package ext

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

type TBuffer struct {
	Bin []byte
}



func WriteBinHeader(len uint32, buff bytes.Buffer) bytes.Buffer {
	if len <= 0xff {
		buff.Write([]byte{msgpack.MessagePackCode_Bin8})
		buff.Write([]byte{byte(len)})
	}else if (len <= 0xffff) {
		buff.Write([]byte{msgpack.MessagePackCode_Bin16})
		b := make([]byte,2)
		binary.BigEndian.PutUint16(b, uint16(len))
		buff.Write(b)
	}else {
		buff.Write([]byte{msgpack.MessagePackCode_Bin32})
		b := make([]byte,4)
		binary.BigEndian.PutUint32(b, len)
		buff.Write(b)
	}
	return buff
}


func init() {
	msgpack.RegisterExt(1, (*TBuffer)(nil))
}


var (
	_ msgpack.Marshaler   = (*TBuffer)(nil)
	_ msgpack.Unmarshaler = (*TBuffer)(nil)
)

func (tm *TBuffer) MarshalMsgpack() ([]byte, error) {
	bin := bytes.Buffer{}
	if len(tm.Bin) <= 0 {
		//write nil
		bin.Write([]byte{msgpack.MessagePackCode_Nil})
		return bin.Bytes(),nil
	}
	//write header
	bin.Write([]byte{msgpack.MessagePackCode_MinFixArray|1})
	bin = WriteBinHeader(uint32( len(tm.Bin) ), bin)
	bin.Write(tm.Bin)
	return bin.Bytes(), nil
}

func (tm *TBuffer) UnmarshalMsgpack(b []byte) error {
	tm.Bin = b
	return nil
}

func (tm *TBuffer) Comp(bb []byte) error {
	if len(tm.Bin) != len(bb) {
		return fmt.Errorf("bin not match for len")
	}
	for i:=0;i<len(bb);i++{
		if tm.Bin[i] != bb[i] {
			return fmt.Errorf("bin not match for idx: %d",i)
		}
	}
	return nil
}