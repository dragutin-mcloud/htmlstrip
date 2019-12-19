package htmlstrip

import "testing"

func TestStrip(t *testing.T) {

	result := StripHTML([]byte("<html>TEST</html>"))

	if result != nil {
		t.Log("Simple test OK")
	} else {
		t.Fail()
	}

	if string(result) == "test" {
		t.Log("Correct sentence returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("BEFORE<html>TEST</html>AFTER"))
	if string(result) == "beforetestafter" {
		t.Log("Extended sentence BEFORE<html>TEST</html>AFTER returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("BEFORE<html>TEST</h"))
	if string(result) == "beforetest" {
		t.Log("Extended sentence BEFORE<html>TEST</h returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("BEFORE<ht"))
	if string(result) == "before" {
		t.Log("Extended sentence BEFORE<ht returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("BEFORE<html>"))
	if string(result) == "before" {
		t.Log("Extended sentence BEFORE<html> returned OK")
	} else {
		t.Fail()
	}

	result = StripHTML([]byte("<html>AFTER"))
	if string(result) == "after" {
		t.Log("Extended sentence <html>AFTER returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("html>AFTER"))
	if string(result) == "html>after" {
		t.Log("Extended sentence html>AFTER returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("<>"))
	if string(result) == "" {
		t.Log("Extended sentence <> returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte(">"))
	if string(result) == ">" {
		t.Log("Extended sentence > returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("<"))
	if string(result) == "" {
		t.Log("Extended sentence returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("<<"))
	if string(result) == "" {
		t.Log("Extended sentence << returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte(">>"))
	if string(result) == ">>" {
		t.Log("Extended sentence returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("<>>"))
	if string(result) == ">" {
		t.Log("Extended sentence <>> returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("<<>"))
	if string(result) == "" {
		t.Log("Extended sentence <<> returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

	result = StripHTML([]byte("&;"))
	if string(result) == "" {
		t.Log("Extended sentence &; returned OK")
	} else {
		t.Log("Correct sentence returned NOK "+string(result))
		t.Fail()
	}

}
