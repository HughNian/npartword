package main

import (
	wor "github.com/HughNian/nmid/worker"
	npw "github.com/HughNian/npartword"
	"github.com/vmihailenco/msgpack"
	"fmt"
	"log"
	"strconv"
)

const SERVERHOST = "192.168.1.176"
const SERVERPORT = "6808"

var parter *npw.Parter

//普通分词
func PartWordsM1(job wor.Job) ([]byte, error) {
	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	text := resp.StrParams[0]
	p2   := resp.StrParams[1]
	tag, _  := strconv.ParseInt(p2, 10, 64)
	partStr := parter.PartWords(text, npw.PART_MODE_ONE, int(tag))

	retStruct := wor.GetRetStruct()
	retStruct.Msg = "ok"
	retStruct.Data = []byte(partStr)
	ret, err := msgpack.Marshal(retStruct)
	if nil != err {
		return []byte(``), err
	}

	resp.RetLen = uint32(len(ret))
	resp.Ret = ret

	return ret, nil
}

//mmseg过滤
func PartWordsM2(job wor.Job) ([]byte, error) {
	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	text := resp.StrParams[0]
	p2   := resp.StrParams[1]
	tag, _  := strconv.ParseInt(p2, 10, 64)
	partStr := parter.PartWords(text, npw.PART_MODE_TWO, int(tag))

	retStruct := wor.GetRetStruct()
	retStruct.Msg = "ok"
	retStruct.Data = []byte(partStr)
	ret, err := msgpack.Marshal(retStruct)
	if nil != err {
		return []byte(``), err
	}

	resp.RetLen = uint32(len(ret))
	resp.Ret = ret

	return ret, nil
}

//隐马尔可夫模型
func PartWordsM3(job wor.Job) ([]byte, error) {
	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	text := resp.StrParams[0]
	p2   := resp.StrParams[1]
	tag, _  := strconv.ParseInt(p2, 10, 64)
	partStr := parter.PartWords(text, npw.PART_MODE_THREE, int(tag))

	retStruct := wor.GetRetStruct()
	retStruct.Msg = "ok"
	retStruct.Data = []byte(partStr)
	ret, err := msgpack.Marshal(retStruct)
	if nil != err {
		return []byte(``), err
	}

	resp.RetLen = uint32(len(ret))
	resp.Ret = ret

	return ret, nil
}

func main() {
	//worker服务
	var worker *wor.Worker
	var err error
	serverAddr := SERVERHOST + ":" + SERVERPORT
	worker = wor.NewWorker()
	err = worker.AddServer("tcp", serverAddr)
	if err != nil {
		log.Fatalln(err)
		worker.WorkerClose()
		return
	}

	//加载字典
	parter = npw.NewParter()
	parter.LoadDictionary("./data/dictionary.txt")

	worker.AddFunction("PartWordsM1", PartWordsM1)
	worker.AddFunction("PartWordsM2", PartWordsM2)
	worker.AddFunction("PartWordsM3", PartWordsM3)

	if err = worker.WorkerReady(); err != nil {
		log.Fatalln(err)
		worker.WorkerClose()
		return
	}

	worker.WorkerDo()
}