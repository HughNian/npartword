package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/HughNian/nmid/pkg/logger"
	"github.com/HughNian/nmid/pkg/model"
	wor "github.com/HughNian/nmid/pkg/worker"
	npw "github.com/HughNian/npartword"
	"github.com/vmihailenco/msgpack"
)

var parter *npw.Parter

// 普通分词
func PartWordsM1(job wor.Job) ([]byte, error) {
	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	retStruct := model.GetRetStruct()
	if len(resp.ParamsMap) == 0 {
		retStruct.Code = 100
		retStruct.Msg = "error"
		retStruct.Data = []byte(``)
		ret, err := msgpack.Marshal(retStruct)
		if nil != ret {
			return []byte(``), err
		}

		resp.RetLen = uint32(len(ret))
		resp.Ret = ret

		return ret, err
	}

	text := resp.ParamsMap["text"].(string)
	p2 := resp.ParamsMap["p2"].(string)
	tag, _ := strconv.ParseInt(p2, 10, 64)
	//partStr := parter.PartWords(text, npw.PART_MODE_ONE, int(tag))
	//partStr := parter.Part(text, npw.PART_MODE_ONE, int(tag)).ToStrings()
	op := parter.Part(text, npw.PART_MODE_ONE, int(tag)).ToStrings()
	partStr := npw.OpRet(op).GetEmoScore()

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

// mmseg过滤
func PartWordsM2(job wor.Job) ([]byte, error) {
	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	retStruct := model.GetRetStruct()
	if len(resp.ParamsMap) == 0 {
		retStruct.Code = 100
		retStruct.Msg = "error"
		retStruct.Data = []byte(``)
		ret, err := msgpack.Marshal(retStruct)
		if nil != ret {
			return []byte(``), err
		}

		resp.RetLen = uint32(len(ret))
		resp.Ret = ret

		return ret, err
	}

	text := resp.ParamsMap["text"].(string)
	p2 := resp.ParamsMap["p2"].(string)
	tag, _ := strconv.ParseInt(p2, 10, 64)
	partStr := parter.PartWords(text, npw.PART_MODE_TWO, int(tag))

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

// 隐马尔可夫模型
func PartWordsM3(job wor.Job) ([]byte, error) {
	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	retStruct := model.GetRetStruct()
	if len(resp.ParamsMap) == 0 {
		retStruct.Code = 100
		retStruct.Msg = "error"
		retStruct.Data = []byte(``)
		ret, err := msgpack.Marshal(retStruct)
		if nil != ret {
			return []byte(``), err
		}

		resp.RetLen = uint32(len(ret))
		resp.Ret = ret

		return ret, err
	}

	text := resp.ParamsMap["text"].(string)
	p2 := resp.ParamsMap["p2"].(string)
	tag, _ := strconv.ParseInt(p2, 10, 64)
	partStr := parter.PartWords(text, npw.PART_MODE_THREE, int(tag))

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
	showLogo()

	//worker服务
	var workerName = "npartword"
	var worker *wor.Worker
	var err error
	serverAddr := os.Getenv("NMIDSERVERHOST") + ":" + os.Getenv("NMIDSERVERPORT")
	// worker = wor.NewWorker().SetWorkerName(workerName).WithTrace(os.Getenv("SKYWALKINGTRACEOAPURL"))
	worker = wor.NewWorker().SetWorkerName(workerName)
	err = worker.AddServer("tcp", serverAddr)
	if err != nil {
		log.Fatalln(err)
		worker.WorkerClose()
		return
	}

	//加载字典
	parter = npw.NewParter()
	parter.LoadDictionary("./data/dictionary.txt")
	parter.LoadEmoDictionary("./data/claim.txt,./data/degree.txt,./data/gainsay.txt,./data/negative_comment.txt," +
		"./data/negative_emotions.txt,./data/positive_comment.txt,./data/positive_emotions.txt")

	worker.AddFunction("PartWordsM1", PartWordsM1)
	worker.AddFunction("PartWordsM2", PartWordsM2)
	worker.AddFunction("PartWordsM3", PartWordsM3)

	if err = worker.WorkerReady(); err != nil {
		log.Fatalln(err)
		worker.WorkerClose()
		return
	}

	go worker.WorkerDo()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	logger.Infof("Server Exit: %s", s.String())
	worker.WorkerClose()
}

func showLogo() {
	logo := `
  ____  ____ _      __
  / __ \/ __ \ | /| / /
 / / / / /_/ / |/ |/ / 
/_/ /_/ .___/|__/|__/  
     /_/               
`
	fmt.Println(logo)
}
