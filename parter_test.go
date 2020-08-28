package npartword

import (
	"testing"
	"fmt"
)

func Test_Dict(t *testing.T) {
	parter := NewParter()
	parter.LoadDictionary("./data/dictionary.txt")

	text := "南京大学城书店"
	//text := "长春市长春药店"
	//text := "研究生命起源"

	str := parter.PartWords(text, PART_MODE_TWO)
	fmt.Println(str)

	texts := parter.PartWordsTexts(text, PART_MODE_TWO)
	fmt.Println(texts)
}
