package htmlstrip

import (
	"bytes"
	"unicode"
)

func charsTo(input []byte, charPoint *int, to rune) string {
	var result string
	result = ""
	for *charPoint < len(input)-1 && input[*charPoint] != byte(to) {
		result += string(input[*charPoint])
		*charPoint++
	}

	*charPoint++

	return result
}

func endWith(input []byte, charPoint *int, to string) {
	var toLength = len(to)
	for string(input[*charPoint:toLength+*charPoint]) != to {
		*charPoint++
	}
}

func removeSpace(input []byte, charPoint *int) {
	if *charPoint+3 > len(input) {
		return
	}
	for bytes.Compare(input[*charPoint:*charPoint+2], []byte("  ")) == 0 {
		*charPoint++
	}
}

func getSymbol(code string) string {

	return code
}

func specialTags(input []byte, charPoint *int) {

	var thisChar []byte = input[*charPoint : *charPoint+1]
	var tag, code string

	if bytes.Compare(thisChar, []byte("&")) == 0 {
		code = charsTo(input, charPoint, ';')
		tag = getSymbol(code)
		return
	}

	if bytes.Compare(thisChar, []byte("<")) == 0 {
		*charPoint++
		tag = charsTo(input, charPoint, '>')
	}

	if tag != "" {
		switch tag {
		case "br":
		case "script":
			_ = removeBetweenScript(&input) // Remove script
		default:
		}
	}

}

func removeNewLine(input []byte, charPoint *int) {
	if bytes.Compare(input[*charPoint:*charPoint+2], []byte("\n\n")) == 0 {
		charsTo(input, charPoint, '>')
	}
}

func alltoLowerCase(input []byte) []byte { // 3.75ms
	return bytes.ToLower(input)
}

func alltoLowerCaseIterate(input []byte) []byte { // 2.26ms
	result := []byte{}
	var b byte
	for c := 0; c < len(input); c++ {
		b = byte(unicode.ToLower(rune(input[c])))
		result = append(result, b)
	}
	return result
}

func findOccurence(input []byte, find string, from int) int {
	var f int = 0
	for c := from; c < len(input); c++ {
		f = 0
		for c<len(input) && f < len(find) && input[c] == find[f] {
			f++
			c++
		}
		if f == len(find) {
			return c - len(find)
		}
	}
	return -1
}

func removeBetweenScript(input *[]byte) (bool) {
	find1 := findOccurence(*input, "<script", 0)
	if find1 < 0 {
		return false
	}
	find2 := findOccurence(*input, "</script>", find1)
	if find2 < 0 {
		return false
	}
	*input = append((*input)[:find1], (*input)[find2+9:]...)
	return true
}

func removeBetweenStyle(input *[]byte) (bool) {
	find1 := findOccurence(*input, "<style", 0)
	if find1 < 0 {
		return false
	}
	find2 := findOccurence(*input, "</style>", find1)
	if find2 < 0 {
		return false
	}
	*input = append((*input)[:find1], (*input)[find2+8:]...)
	return true
}

func StripHTML(input []byte) []byte {
	var charPoint int
	var output []byte

	input = alltoLowerCase(input)	// Convert all to lowercase
	for removeBetweenScript(&input)==true {
	}
	for removeBetweenStyle(&input)==true {
	}

	for charPoint = 0; charPoint < len(input); charPoint++ {

		removeSpace(input, &charPoint)
		if charPoint >= len(input) {
			continue
		}

		specialTags(input, &charPoint)
		if charPoint > len(input) {
			continue
		}

		for charPoint < len(input)-2 && input[charPoint] == '<' {
			specialTags(input, &charPoint)
		}
		if charPoint > len(input)-1 {
			continue
		}

		switch input[charPoint] {
		case '<':
			charsTo(input, &charPoint, '>')
			continue
		case '\t', '\f', '\r', '\n':
			continue
		default:
		}
		if input[charPoint] == ' ' && len(output) > 0 && output[len(output)-1] == ' ' {
			continue;
		}
		output = append(output, input[charPoint])
	}
	return output
}
