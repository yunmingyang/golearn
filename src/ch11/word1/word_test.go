package word

import "testing"



func TestPalindrom(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrom("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrom("kayak") = false`)
	}
}

func TestNonPalindrom(t *testing.T) {
	if IsPalindrome("palindrom") {
		t.Error(`IsPalindrom("palindrom") = true`)
	}
}

func TestFrenchPalindrom(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrom("été") = false`)
	}
}

func TestCanalPalindrom(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrom(%q) = false \n`, input)
	}
}