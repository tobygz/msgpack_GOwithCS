package main

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
	"log"
	"sync"
	"time"
)

var ct int =0
var mtx sync.Mutex

func Inc(){
	mtx.Lock()
	defer mtx.Unlock()
	ct++
}

func Fetch() (ret int) {
	mtx.Lock()
	defer mtx.Unlock()
	ret = ct
	ct = 0
	return
}


func testQPS(){
	lastTick := time.Now()

	go func(){
		for {
			if time.Now().Sub(lastTick).Seconds() >= 1 {
				log.Println("qps:", Fetch())
				lastTick = time.Now()
			}
		}
	}()

	for {
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
		//log.Println(buf.Bytes())

		tmDec := HttpBasePacket{}
		dec := msgpack.NewDecoder(&buf)
		err = dec.Decode(&tmDec)
		if err != nil {
			log.Fatal(err)
			return
		}
		Inc()
	}

}