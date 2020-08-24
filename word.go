package npartword

type Word struct {
	chars     []byte

	rate      int

	pos       string

	distance  float32

	parts	  []*Part //该词的再次切分
}

func NewWord(chars []byte, rate int, pos string) *Word {
	return &Word {
		chars : chars,
		rate  : rate,
		pos   : pos,
	}
}

