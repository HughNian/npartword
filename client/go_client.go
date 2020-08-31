package main

import (
	"github.com/vmihailenco/msgpack"
	cli "github.com/HughNian/nmid/client"
	"log"
	"fmt"
	"os"
)

const SERVERHOST = "192.168.1.176"
const SERVERPORT = "6808"

func main() {
	var client *cli.Client
	var err error

	serverAddr := SERVERHOST + ":" + SERVERPORT
	client, err = cli.NewClient("tcp", serverAddr)
	if nil == client || err != nil {
		log.Println(err)
		return
	}
	defer client.Close()

	client.ErrHandler= func(e error) {
		if cli.RESTIMEOUT == e {
			log.Println("time out here")
		} else {
			log.Println(e)
		}
		fmt.Println("client err here")
		//client.Close()
	}

	respHandler := func(resp *cli.Response) {
		if resp.DataType == cli.PDT_S_RETURN_DATA && resp.RetLen != 0 {
			if resp.RetLen == 0 {
				log.Println("ret empty")
				return
			}

			var retStruct cli.RetStruct
			err := msgpack.Unmarshal(resp.Ret, &retStruct)
			if nil != err {
				log.Fatalln(err)
				return
			}

			if retStruct.Code != 0 {
				log.Println(retStruct.Msg)
				return
			}

			fmt.Println(string(retStruct.Data))
		}
	}

	text := []string{"南京大学城书店，长春市长春药店，研究生命起源", "2"}
	params, err := msgpack.Marshal(&text)
	if err != nil {
		log.Fatalln("params msgpack error:", err)
		os.Exit(1)
	}

	err = client.Do("PartWordsM1", params, respHandler)
	if nil != err {
		fmt.Println(err)
	}

	err = client.Do("PartWordsM2", params, respHandler)
	if nil != err {
		fmt.Println(err)
	}

	err = client.Do("PartWordsM3", params, respHandler)
	if nil != err {
		fmt.Println(err)
	}
}