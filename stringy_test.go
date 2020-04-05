package string_manipulation

import (
	"testing"
)

var sm StringManipulation = New("This is example.")

func TestInput_Between(t *testing.T) {
	val := sm.Between("This", "example").ToUpper()
	if val != "IS" {
		t.Errorf("Expected: %s but got: %s", "IS", val)
	}
}

func TestInput_Boolean(t *testing.T) {
	str := New("on")
	val := str.Boolean()
	if !val {
		t.Errorf("Expected: to be true but got: %v", val)
	}
}

func TestInput_CamelCase(t *testing.T) {
	str := New("Camel case this_complicated__string%%")
	val := str.CamelCase("%", "")
	if val != "CamelCaseThisComplicatedString" {
		t.Errorf("Expected: to be %s but got: %s", "CamelCaseThisComplicatedString", val)
	}
}

func TestInput_ContainsAll(t *testing.T) {
	contains := New("hello mam how are you??")
	if val := contains.ContainsAll("mam", "?"); !val {
		t.Errorf("Expected value to be true but got false")
	}
	if val := contains.ContainsAll("non existent"); val {
		t.Errorf("Expected value to be false but got true")
	}
}

func TestInput_Delimited(t *testing.T) {
	str := New("Delimited case this_complicated__string@@")
	against := "delimited.case.this.complicated.string"
	if val := str.Delimited(".", "@", "").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_KebabCase(t *testing.T) {
	str := New("Kebab case this-complicated___string@@")
	against := "Kebab-case-this-complicated-string"
	if val := str.KebabCase("@", "").Get(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_LcFirst(t *testing.T) {
	str := New("this is an all lower")
	against := "This is an all lower"
	if val := str.LcFirst(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_Lines(t *testing.T) {
	lines := New("fòô\r\nbàř\nyolo")
	strSlic := lines.Lines()
	if len(strSlic) != 3 {
		t.Errorf("Length expected to be 3 but got: %d", len(strSlic))
	}
	if strSlic[0] != "fòô" {
		t.Errorf("Expected: %s but got: %s", "fòô", strSlic[0])
	}
}

func TestInput_Pad(t *testing.T) {
	pad := New("Roshan")
	if result := pad.Pad(10, "0", "both"); result != "00Roshan00" {
		t.Errorf("Expected: %s but got: %s", "00Roshan00", result)
	}
	if result := pad.Pad(10, "0", "left"); result != "0000Roshan" {
		t.Errorf("Expected: %s but got: %s", "0000Roshan", result)
	}
	if result := pad.Pad(10, "0", "right"); result != "Roshan0000" {
		t.Errorf("Expected: %s but got: %s", "Roshan0000", result)
	}
}

func TestInput_RemoveSpecialCharacter(t *testing.T) {
	cleanString := New("special@#remove%%%%")
	against := "specialremove"
	if result := cleanString.RemoveSpecialCharacter(); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_ReplaceFirst(t *testing.T) {
	replaceFirst := New("Hello My name is Roshan and his name is Alis.")
	against := "Hello My nombre is Roshan and his name is Alis."
	if result := replaceFirst.ReplaceFirst("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_ReplaceLast(t *testing.T) {
	replaceLast := New("Hello My name is Roshan and his name is Alis.")
	against := "Hello My name is Roshan and his nombre is Alis."
	if result := replaceLast.ReplaceLast("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_Reverse(t *testing.T) {
	reverseString := New("roshan")
	against := "nahsor"
	if result := reverseString.Reverse(); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_Shuffle(t *testing.T) {
	check := "roshan"
	shuffleString := New(check)
	if result := shuffleString.Shuffle(); len(result) != len(check) && check == result {
		t.Errorf("Shuffle string gave wrong output")
	}
}

func TestInput_SnakeCase(t *testing.T) {
	str := New("SnakeCase this-complicated___string@@")
	against := "snake_case_this_complicated_string"
	if val := str.SnakeCase("@", "").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_Surround(t *testing.T) {
	str := New("this")
	against := "__this__"
	if val := str.Surround("__"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_Tease(t *testing.T) {
	str := New("This is just simple paragraph on lorem ipsum.")
	against := "This is just..."
	if val := str.Tease(12, "..."); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_UcFirst(t *testing.T) {
	str := New("this is test")
	against := "This is test"
	if val := str.UcFirst(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}