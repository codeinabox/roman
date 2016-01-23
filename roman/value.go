package roman

import (
	"errors"
	"regexp"
	"strings"
)

type Numeral struct {
	value string
}

var lookup = map[string]uint{
	"I":  1,
	"IV": 4,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
}

var order = []string{"M", "D", "C", "L", "X", "V", "IV", "I"}

func pattern() *regexp.Regexp {
	return regexp.MustCompile("[MDCLXVI]+")
}

func NewNumeral(v interface{}) (Numeral, error) {
	var n Numeral
	switch t := v.(type) {
	case string:
		if !pattern().Match([]byte(t)) {
			return n, errors.New("Not a valid numeral")
		}
		n.value = t
	case uint:
		n.value = Itoa(t)
	case int:
		if t < 0 {
			return n, errors.New("can not represent negative numbers")
		}
		n.value = Itoa(uint(t))
	}

	return n, nil
}

func (n Numeral) SameValueAs(nn Numeral) bool {
	return n.value == nn.value
}

func Itoa(in uint) string {
	remainder := in
	acc := ""

	for _, k := range order {
		currentVal := lookup[k]
		c := remainder / currentVal
		// fmt.Printf("%v / %v = %v\n", remainder, currentVal, c)
		acc += strings.Repeat(k, int(c))
		remainder = remainder - (c * currentVal)
	}

	return acc
}
