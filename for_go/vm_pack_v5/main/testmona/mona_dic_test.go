package testmona

import (
	"bytes"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"log"
	"testing"
)

type httpMapPacket struct {
	Dic map[int]map[int]string
}

func Test_mona_dic(t *testing.T){
	hmp := &httpMapPacket{
		Dic:make(map[int]map[int]string),
	}
	mp := make(map[int]string)
	mp[2] = "A"
	hmp.Dic[1] = mp
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)
	err := enc.Encode(&hmp)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(buf.Bytes())


	tmDec := httpMapPacket{}
	dec := msgpack.NewDecoder(&buf)
	err = dec.Decode(&tmDec)
	if err != nil {
		log.Fatal(err)
		return
	}

	jss,_ := json.Marshal(tmDec.Dic)
	log.Println( string(jss))
}