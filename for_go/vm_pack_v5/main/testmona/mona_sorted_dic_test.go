package testmona

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
	"log"
	"testing"
)

type httpSDicPacket struct {
	SDic map[int]int
}

func Test_mona_sorted_dic(t *testing.T){
	hmp := &httpSDicPacket{
		SDic:make(map[int]int),
	}
	hmp.SDic[3] = 3
	hmp.SDic[1] = 1
	hmp.SDic[2] = 2
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	enc.SetSortMapKeys(true)
	err := enc.Encode(&hmp)
	if err != nil {
		log.Panic(err)
	}
	log.Println(buf.Bytes())


	//decode
	tmDec := httpSDicPacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(tmDec.SDic)
}