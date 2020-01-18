package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type binOp int

const (
	ADD binOp = iota
	SUB
	MULT
	DIV
)

type lexeme int

const (
	OPERATOR lexeme = iota
	LITERAL
	IDENTIFIER
	ASSIGNMENT
	FUNCTION
	STATEMENT
	UNKNOWN
)

type scanner struct {
	tokenList     map[int]string
	tokenTypeList map[int]lexeme
}

var tokenListType scanner

func checkError(err error) {
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
}

func fileGetContents(filename string) string {
	contents := new(bytes.Buffer)
	f, err := os.Open(filename)
	checkError(err)
	_, err = io.Copy(contents, f)
	if err != io.EOF {
		checkError(err)
	}
	checkError(f.Close())
	return contents.String()
}

func checkCharacter(c byte) lexeme {

	if (c >= byte(65) && c <= byte(90)) || (c >= byte(97) && c <= byte(122)) {
		return IDENTIFIER
	}

	if c == byte(46) {
		return LITERAL // .
	}

	if c >= byte(40) && c <= byte(47) {
		return OPERATOR // common operators
	}

	if c >= byte(48) && c <= byte(57) {
		return LITERAL // 0-9
	}

	if c == byte(61) {
		return ASSIGNMENT // =
	}

	if c == byte(10) {
		return STATEMENT
	}

	return UNKNOWN
}

func lex(text string) {

	var tokenList = make(map[int]string)
	var tokenTypeList = make(map[int]lexeme)
	var tokenIndex int
	var token string
	var lookAhead byte
	var lookAheadType lexeme

	for i := 0; i < len(text); i++ {
		if text[i] != byte(32) {
			if i+1 < len(text) {
				lookAhead = text[i+1]
				lookAheadType = checkCharacter(lookAhead)
			}
			tokenType := checkCharacter(text[i])
			if tokenType != lookAheadType {
				// -literal
				if text[i] == byte(45) && lookAheadType == LITERAL {
					token += string(text[i])
				} else {
					token += string(text[i])
					tokenList[tokenIndex] = string(token)
					tokenTypeList[tokenIndex] = tokenType
					tokenIndex++
					token = ""
				}
			} else {
				token += string(text[i])
			}
		}
	}

	tokenListType.tokenList = tokenList
	tokenListType.tokenTypeList = tokenTypeList

}

func showTokens() {
	var lookup = make(map[lexeme]string)
	lookup[IDENTIFIER] = "IDENTIFIER"
	lookup[OPERATOR] = "OPERATOR"
	lookup[LITERAL] = "LITERAL"
	lookup[ASSIGNMENT] = "ASSIGNMENT"
	lookup[STATEMENT] = "END OF LINE"
	lookup[UNKNOWN] = "UNKNOWN"

	for n := 0; n < len(tokenListType.tokenList); n++ {
		if tokenListType.tokenTypeList[n] != STATEMENT {
			fmt.Printf("tokenList[%v]\t%10s\t%v\n", n, tokenListType.tokenList[n], lookup[tokenListType.tokenTypeList[n]])
		} else {
			fmt.Printf("tokenList[%v]\t%10s\t%v\n", n, "NEWLINE", lookup[tokenListType.tokenTypeList[n]])
		}
	}

}

func main() {
	fmt.Println("Scanner")
	text := fileGetContents("./txt/test.slw")

	lex(text)
	showTokens()

	parse()
}
