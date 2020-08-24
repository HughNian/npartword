package npartword

import (
	"testing"
	"fmt"
)

var (
	str1  = "明"
	str2  = "朝"
	str3  = "明朝"
	str4  = "明洪武"
	str5  = "1号店"
	str6  = "明洪熙"
	str7  = "洪武"
	str8  = "洪武朱元璋"
	str9  = "锦衣卫"
	str10 = "明朝锦衣卫"
	str11 = "朝拜"
	str12 = "翰林院"
	str13 = "明朝翰林院"
	str14 = "南"
	str15 = "南京"
	str16 = "南京大学"
	str17 = "机油0W-30"
)

func Test_func(t *testing.T) {
	trie := NewTrie()
	trie.AddKey(str1, 0)
	trie.AddKey(str2, 1)
	trie.AddKey(str3, 2)
	trie.AddKey(str4, 3)
	trie.AddKey(str5, 4)
	trie.AddKey(str6, 5)
	trie.AddKey(str7, 6)
	trie.AddKey(str8, 7)
	trie.AddKey(str9, 8)
	trie.AddKey(str10, 9)
	trie.AddKey(str11, 10)
	trie.AddKey(str12, 11)
	trie.AddKey(str13, 12)
	trie.AddKey(str14, 13)
	trie.AddKey(str15, 14)
	trie.AddKey(str16, 15)
	trie.AddKey(str17, 16)

	ret, ok := trie.SearchKeyByPre2Next(str15, str16)
	if ok {
		fmt.Println(ret)
	} else {
		fmt.Println("not match")
	}
}

func Test_FindKey(t *testing.T) {
	trie := NewTrie()
	trie.AddKey(str15, 14)
	trie.AddKey(str16, 15)
	trie.AddKey(str17, 16)

	n, ok := trie.FindKey(str17)
	if ok {
		fmt.Println(n.path)
		fmt.Println(n.meta)
	} else {
		fmt.Println("not find")
	}
}

func Test_FindKeyNode(t *testing.T) {
	trie := NewTrie()
	trie.AddKey(str14, 14)
	trie.AddKey(str15, 15)
	trie.AddKey(str16, 16)

	key, ok := trie.SearchKeyNodeByPre2Next(str15, "大")
	if !ok {
		fmt.Println("not find")
	}

	fmt.Println(key)
}