package testmona

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/main/ext"
	"log"
	"testing"
)


type HttpBasePacket struct {
	Uuid string
	Data ext.TBuffer
}

func Test_mona_tbuff(t *testing.T){
	//encode
	tm := HttpBasePacket{}
	tm.Data.Bin = []byte{1,1,1}
	tm.Uuid = "A"
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	err := enc.Encode(&tm)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(buf.Bytes())


	//decode
	tmDec := HttpBasePacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		t.Fatal(err)
		return
	}

	if err = tmDec.Data.Comp(tm.Data.Bin); err != nil {
		t.Fatal("data.bin not match")
	}

	log.Println(tmDec.Data.Bin)

	//1,1,1 => [145 145 196 3 1 1 1]
	//empty => [145 145 192]
	return
}

func Test_mona_tbuff_empty(t *testing.T){
	//encode
	tm := HttpBasePacket{}
	//tm.Data.Bin = []byte{1,1,1}
	//tm.Uuid = "A"
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	err := enc.Encode(&tm)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(buf.Bytes())

	//decode
	tmDec := HttpBasePacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		t.Fatal(err)
		return
	}

	if tmDec.Data.Bin != nil {
		t.Fatal("buf should be empty")
	}
	return
}


func Test_mona_tbuff_empty_bin(t *testing.T){
	//encode
	tm := HttpBasePacket{}
	//tm.Data.Bin = []byte{1,1,1}
	tm.Uuid = "A"
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	err := enc.Encode(&tm)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(buf.Bytes())

	//decode
	tmDec := HttpBasePacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		t.Fatal(err)
		return
	}

	if tmDec.Uuid != "A" || tmDec.Data.Bin != nil {
		t.Fatal("decode succ, result not match")
	}
	return
}

func Benchmark_monabuff(t *testing.B) {
	//encode
	tm := HttpBasePacket{}
	tm.Data.Bin = []byte{1,1,1}
	tm.Uuid = "A"
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	err := enc.Encode(&tm)
	if err != nil {
		t.Fatal(err)
	}
	//log.Println(buf.Bytes())


	//decode
	tmDec := HttpBasePacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		t.Fatal(err)
		return
	}

	//t.Log(tmDec.Data.Bin)
}