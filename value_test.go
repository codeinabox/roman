package roman_test

import (
	"github.com/codeinabox/roman"
	"testing"
)

var integerToNumeralTests = []struct {
	integer int
	numeral string
}{
	{1, "I"},
	{2, "II"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{10, "X"},
	{50, "L"},
	{100, "C"},
	{257, "CCLVII"},
	{500, "D"},
	{1000, "M"},
	{2000, "MM"},
	{2257, "MMCCLVII"},
}

// Create a different type of value object used in Equals() check
type NotNumeral struct {
	value string
}

func (n NotNumeral) String() string {
	return n.value
}

func (n NotNumeral) Equals(value roman.Value) bool {
	return false
}

func TestConvertIntegerToNumeral(t *testing.T) {
	for _, example := range integerToNumeralTests {
		n, err := roman.NewNumeral(example.integer)
		if err != nil {
			t.Fatal(err)
		}
		if n.String() != example.numeral {
			t.Fatalf("string representation should be %s, was %s", example.numeral, n.String())
		}
	}
}

func TestShouldntAcceptInvalidString(t *testing.T) {
	_, err := roman.NewNumeral("B")
	if err == nil {
		t.Fatal("We expected an error with A")
	}
}

func TestShouldNotBeEqualIfNotNumeral(t *testing.T) {
	var notNumeral NotNumeral
	numeral, _ := roman.NewNumeral("I")

	if numeral.Equals(notNumeral) == true {
		t.Fatal("Different value object types can not be equal")
	}
}

func TestShouldBeEqualIfSameNumeral(t *testing.T) {
	a, _ := roman.NewNumeral("I")
	b, _ := roman.NewNumeral("I")
	if a.Equals(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldBeEqualIfIntegerEquivalent(t *testing.T) {
	a, _ := roman.NewNumeral(5)
	b, _ := roman.NewNumeral("V")
	if a.Equals(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldCompareTwoNumeralsAsNotEqual(t *testing.T) {
	a, _ := roman.NewNumeral("I")
	b, _ := roman.NewNumeral("X")
	if a.Equals(b) == true {
		t.Fatal("Shouldn't be same value")
	}
}
