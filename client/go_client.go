package main

import (
	"fmt"
	cli "github.com/HughNian/nmid/pkg/client"
	"github.com/HughNian/nmid/pkg/model"
	"github.com/vmihailenco/msgpack"
	"log"
	"os"
)

const SERVERHOST = "127.0.0.1"
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

	client.ErrHandler = func(e error) {
		if model.RESTIMEOUT == e {
			log.Println("time out here")
		} else {
			log.Println(e)
		}
		fmt.Println("client err here")
		//client.Close()
	}

	respHandler := func(resp *cli.Response) {
		if resp.DataType == model.PDT_S_RETURN_DATA && resp.RetLen != 0 {
			if resp.RetLen == 0 {
				log.Println("ret empty")
				return
			}

			var retStruct model.RetStruct
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

	paramsName1 := make(map[string]interface{})
	paramsName1["text"] = "南京大学城书店，长春市长春药店，研究生命起源"
	paramsName1["p2"] = "2"
	params, err := msgpack.Marshal(&paramsName1)
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
