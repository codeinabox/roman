package roman

import "testing"

func TestSomething(t *testing.T) {
	var n Numeral

	if n.value != "" {
		t.Errorf("something went wrong")
	}
}

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

func TestConvertIntegerToNumeral(t *testing.T) {
	for _, example := range integerToNumeralTests {
		n, err := NewNumeral(example.integer)
		if err != nil {
			t.Fatal(err)
		}
		if n.value != example.numeral {
			t.Fatalf("string representation should be %s, was %s", example.numeral, n.value)
		}
	}
}

func TestShouldntAcceptInvalidString(t *testing.T) {
	_, err := NewNumeral("A")
	if err == nil {
		t.Fatal("We expected an error with A")
	}
}

func TestShouldBeSameIfSameNumeral(t *testing.T) {
	a, _ := NewNumeral("I")
	b, _ := NewNumeral("I")
	if a.SameValueAs(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldBeSameIfIntegerEquivalent(t *testing.T) {
	a, _ := NewNumeral(5)
	b, _ := NewNumeral("V")
	if a.SameValueAs(b) == false {
		t.Fatal("Not same value as")
	}
}

func TestShouldCompareTwoNumeralsAsNotSame(t *testing.T) {
	a, _ := NewNumeral("I")
	b, _ := NewNumeral("X")
	if a.SameValueAs(b) == true {
		t.Fatal("Shouldn't be same value")
	}
}
