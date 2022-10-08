package testmona

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
	"log"
	"testing"
)

type httpListPacket struct {
	List [][]int
}

func Test_mona_list(t *testing.T){
	hmp := &httpListPacket{
		List:make([][]int,0,0),
	}
	lst := make([]int,0,0)
	lst = append(lst,1)
	lst = append(lst,2)
	lst = append(lst,3)
	lst = append(lst,2)
	hmp.List = append(hmp.List, lst)
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	err := enc.Encode(&hmp)
	if err != nil {
		log.Panic(err)
	}
	log.Println(buf.Bytes())

	tmDec := httpListPacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(tmDec.List)
}