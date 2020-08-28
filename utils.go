package npartword

import "fmt"

//将英文词转化为小写
func ToLower(text []byte) []byte {
	output := make([]byte, len(text))
	for i, t := range text {
		if t >= 'A' && t <= 'Z' {
			output[i] = t - 'A' + 'a'
		} else {
			output[i] = t
		}
	}
	return output
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func FilterPartNil(part []*Part) []*Part {
	newPart := make([]*Part, 0, len(part))
	for _, item := range part {
		if item != nil {
			newPart = append(newPart, item)
		}
	}

	return newPart
}

func PartToStrings(part []*Part, posShow bool) (output string) {
	if len(part) <= 0 {
		return ``
	}

	xchars := make([]byte, 0)
	for i := range part {
		for _, w := range part[i].Word() {
			if w.pos == "x" && w.chars[0] != SPACE {
				xchars = append(xchars, w.chars...)
			} else {
				if len(xchars) > 0 {
					if posShow {
						output += fmt.Sprintf("%s/%s ", string(xchars), "x")
					} else {
						output += fmt.Sprintf("%s|", string(xchars))
					}
					xchars = make([]byte, 0)
				}

				if posShow {
					output += fmt.Sprintf("%s/%s ", string(w.chars), w.pos)
				} else {
					output += fmt.Sprintf("%s|", string(w.chars))
				}
			}
		}
	}
	if len(xchars) > 0 {
		if posShow {
			output += fmt.Sprintf("%s/%s ", string(xchars), "x")
		} else {
			output += fmt.Sprintf("%s|", string(xchars))
		}
		xchars = make([]byte, 0)
	}

	return output
}

func PartToTexts(part []*Part) (output []string) {
	if len(part) <= 0 {
		return nil
	}

	output = make([]string, 0)
	xchars := make([]byte, 0)
	for i := range part {
		for _, w := range part[i].Word() {
			if w.pos == "x" && w.chars[0] != SPACE {
				xchars = append(xchars, w.chars...)
			} else {
				if len(xchars) > 0 {
					output = append(output, string(xchars))
					xchars = make([]byte, 0)
				}

				output = append(output, string(w.chars))
			}
		}
	}
	if len(xchars) > 0 {
		output = append(output, string(xchars))
		xchars = make([]byte, 0)
	}

	return output
}