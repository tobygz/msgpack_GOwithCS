package main

import (
	"bytes"
	"fmt"
	"github.com/vmihailenco/msgpack/v5/main/ext"
	"log"

	"github.com/vmihailenco/msgpack/v5"
)

type HttpBasePacket struct {
	Uuid string
	Data ext.TBuffer
}

type HttpBasePacketTest struct {
	Uuid string
}

type httpMapPacket struct {
	Dic map[int]map[int]string
}

type httpListPacket struct {
	List [][]int
}

type httpSDicPacket struct {
	SDic map[int]int
}



func main() {
	if true {
		tm := HttpBasePacket{}
		//tm.Data.Bin = []byte{1,1,1}
		tm.Uuid = "A"
		var buf bytes.Buffer
		enc := msgpack.NewEncoder(&buf)
		enc.UseArrayEncodedStructs(true)
		err := enc.Encode(&tm)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(buf.Bytes())

		//decode
		tmDec := HttpBasePacket{}
		dec := msgpack.NewDecoder(&buf)
		err = dec.Decode(&tmDec)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}


	if true {
		testQPS()
		return
	}


	if true {
		tm := HttpBasePacket{}
		tm.Data.Bin = []byte{1,1,1}
		tm.Uuid = "A"
		var buf bytes.Buffer
		enc := msgpack.NewEncoder(&buf)
		enc.UseArrayEncodedStructs(true)
		err := enc.Encode(&tm)
		if err != nil {
			log.Panic(err)
		}
		log.Println(buf.Bytes())

		tmDec := HttpBasePacket{}
		dec := msgpack.NewDecoder(&buf)
		err = dec.Decode(&tmDec)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(tmDec.Data.Bin)
		log.Println(tmDec.Uuid)
		//1,1,1 => [145 145 196 3 1 1 1]
		//empty => [145 145 192]
		return
	}


	if true {
		//smap
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
		//enc.UseCompactInts(true)
		err := enc.Encode(&hmp)
		if err != nil {
			log.Panic(err)
		}
		log.Println(buf.Bytes())
		return
	}

	if true {
		//list
		hmp := &httpListPacket{
			List:make([][]int,0,0),
		}
		lst := make([]int,0,0)
		lst = append(lst,1)
		hmp.List = append(hmp.List, lst)
		var buf bytes.Buffer
		enc := msgpack.NewEncoder(&buf)
		enc.UseArrayEncodedStructs(true)
		err := enc.Encode(&hmp)
		if err != nil {
			log.Panic(err)
		}
		log.Println(buf.Bytes())
		return
	}

	if true {
		//map
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
			log.Panic(err)
		}
		log.Println(buf.Bytes())
		return
	}



	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)

	err := enc.Encode(&HttpBasePacketTest{Uuid: "8"})
	if err != nil {
		log.Panic(err)
		return
	}
	for i := 0; i < len(buf.Bytes()); i++ {
		fmt.Printf("%d,", buf.Bytes()[i])
	}
	//145,161,56,

}
