package npartword

type Dictionary struct {
	trie        *Trie
	maxCharLen  int
	words       []*Word
	totalRate   int64
}

func NewDictionary() *Dictionary {
	return &Dictionary {
		trie       : NewTrie(),
		maxCharLen : 0,
		words      : make([]*Word, 0),
		totalRate  : 0,
	}
}

func (dict *Dictionary) Trie() *Trie {
	return dict.trie
}

func (dict *Dictionary) MaxCharLen() int {
	return dict.maxCharLen
}

func (dict *Dictionary) WordsNum() int {
	return len(dict.words)
}

func (dict *Dictionary) TotalRate() int64 {
	return dict.totalRate
}

func (dict *Dictionary) AddWord(word *Word) {
	_, exist := dict.trie.FindKey(string(word.chars))
	if exist {
		return
	}

	dict.trie.AddKey(string(word.chars), dict.WordsNum())
	dict.words = append(dict.words, word)
	dict.totalRate += int64(word.rate)
	if dict.maxCharLen < len(word.chars) {
		dict.maxCharLen = len(word.chars)
	}
}

func (dict *Dictionary) FindWord(chars []rune) ([]*Word, int) {
	if len(chars) == 0 {
		return nil, 0
	}
	var (
		pre string
		i   int
		num int
		findNode bool
	)
	words := make([]*Word, 0)
	pre = ``
	num = 0
	findNode = false
	for i = range chars {
		next := string(chars[i])
		fpre, ok := dict.Trie().SearchKeyByPre2Next(pre, next)
		if !ok {
			fpre, ok = dict.Trie().SearchKeyNodeByPre2Next(pre, next)
			if !ok {
				break
			}
			findNode = true
		}
		pre = fpre

		num++

		n, exist := dict.Trie().FindKey(pre)
		if !exist {
			if findNode {
				findNode = false
				num--
				continue
			} else {
				word := NewWord([]byte(pre), 1, "x")
				words = append(words, word)
				break
			}
		}
		windex := n.meta.(int)
		word := dict.words[windex]
		//fmt.Println("chars:", string(word.chars))
		words = append(words, word)
	}

	return words, num
}