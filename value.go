package roman

import (
	"errors"
	"regexp"
	"strings"
)

// Errors block
var (
	ErrInvalidNumeral = errors.New("Not a valid numeral")
	ErrNegativeValue  = errors.New("can not represent negative numbers")
)

var (
	lookup = map[string]uint{
		"I":  1,
		"IV": 4,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}

	order = []string{"M", "D", "C", "L", "X", "V", "IV", "I"}
)

// Value stores values
type Value interface {
	String() string
	Equals(value Value) bool
}

// Numeral represents a number
type Numeral struct {
	value string
}

func pattern() *regexp.Regexp {
	return regexp.MustCompile("[MDCLXVI]+")
}

// NewNumeral creates a new numeral
func NewNumeral(v interface{}) (Numeral, error) {
	var n Numeral
	switch t := v.(type) {
	case string:
		if !pattern().Match([]byte(t)) {
			return n, ErrInvalidNumeral
		}
		n.value = t
	case uint:
		n.value = itoa(t)
	case int:
		if t < 0 {
			return n, ErrNegativeValue
		}
		n.value = itoa(uint(t))
	}

	return n, nil
}

// String returns string representation
func (n Numeral) String() string {
	return n.value
}

// Equals checks that two values are the same
func (n Numeral) Equals(value Value) bool {
	otherNumeral, ok := value.(Numeral)
	return ok && n.value == otherNumeral.value
}

func itoa(in uint) string {
	remainder := in
	acc := ""

	for _, k := range order {
		currentVal := lookup[k]
		c := remainder / currentVal
		acc += strings.Repeat(k, int(c))
		remainder = remainder - (c * currentVal)
	}

	return acc
}
