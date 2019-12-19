package htmlstrip

import "testing"

func TestStrip1(t *testing.T) {

	result := StripHTML([]byte("<html>TEST</html>"))

	if result != nil {
		t.Log("Simple test OK")
	} else {
		t.Fail()
	}
}
func TestStrip2(t *testing.T) {
	result := StripHTML([]byte("<html>TEST</html>"))
	if string(result) == "test" {
		t.Log("Correct sentence returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip3(t *testing.T) {
	result := StripHTML([]byte("BEFORE<html>TEST</html>AFTER"))
	if string(result) == "beforetestafter" {
		t.Log("Extended sentence BEFORE<html>TEST</html>AFTER returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip4(t *testing.T) {
	result := StripHTML([]byte("BEFORE<html>TEST</h"))
	if string(result) == "beforetest" {
		t.Log("Extended sentence BEFORE<html>TEST</h returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip5(t *testing.T) {
	result := StripHTML([]byte("BEFORE<ht"))
	if string(result) == "before" {
		t.Log("Extended sentence BEFORE<ht returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip6(t *testing.T) {
	result := StripHTML([]byte("BEFORE<html>"))
	if string(result) == "before" {
		t.Log("Extended sentence BEFORE<html> returned OK")
	} else {
		t.Fail()
	}
}
func TestStrip7(t *testing.T) {
	result := StripHTML([]byte("<html>AFTER"))
	if string(result) == "after" {
		t.Log("Extended sentence <html>AFTER returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip8(t *testing.T) {
	result := StripHTML([]byte("html>AFTER"))
	if string(result) == "html>after" {
		t.Log("Extended sentence html>AFTER returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip9(t *testing.T) {
	result := StripHTML([]byte("<>"))
	if string(result) == "" {
		t.Log("Extended sentence <> returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip10(t *testing.T) {
	result := StripHTML([]byte(">"))
	if string(result) == ">" {
		t.Log("Extended sentence > returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip11(t *testing.T) {
	result := StripHTML([]byte("<"))
	if string(result) == "" {
		t.Log("Extended sentence returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip12(t *testing.T) {
	result := StripHTML([]byte("<<"))
	if string(result) == "" {
		t.Log("Extended sentence << returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip13(t *testing.T) {
	result := StripHTML([]byte(">>"))
	if string(result) == ">>" {
		t.Log("Extended sentence returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip14(t *testing.T) {
	result := StripHTML([]byte("<>>"))
	if string(result) == ">" {
		t.Log("Extended sentence <>> returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip15(t *testing.T) {
	result := StripHTML([]byte("<<>"))
	if string(result) == "" {
		t.Log("Extended sentence <<> returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
func TestStrip16(t *testing.T) {
	result := StripHTML([]byte("&;"))
	if string(result) == "" {
		t.Log("Extended sentence &; returned OK")
	} else {
		t.Log("Correct sentence returned NOK " + string(result))
		t.Fail()
	}
}
