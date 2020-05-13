package test

import (
	"log"

	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	str := "Hello, "
	str = str +  "World!"

	log.Printf("str : %s\n", str)
}

func TestConcatMulti(t *testing.T) {
	str := "ハロー、"
	str = str + "ワールド"

	log.Printf("str : %s\n", str)
}

func TestLen(t *testing.T) {
	str := "Hello, World!"

	log.Printf("str : %d\n", len(str))
}

func TestLenMulti(t *testing.T) {
	str := "ハローワールド"

	log.Printf("str : %d\n", len(str)) // not 7
	log.Printf("rune : %d\n", len([]rune(str))) // 7
}

func TestSubstr(t *testing.T) {
	str := "Hello, World!"

	log.Printf("str[1:] : %s\n", str[1:])
	log.Printf("str[2:4] : %s\n", str[2:4])
	log.Printf("str[:5] : %s\n", str[:5])
}

func TestSubstrMulti(t *testing.T) {
	str := "ハローワールド"
	log.Printf("str[1:] : %s\n", str[1:]) // not ローワールド
	log.Printf("str[2:4] : %s\n", str[2:4]) // not ローワ
	log.Printf("str[:5] : %s\n", str[:5]) // not ハローワールド

	log.Printf("rune[1:] : %s\n", string([]rune(str)[1:])) // ローワールド
	log.Printf("rune[2:4] : %s\n", string([]rune(str)[2:4])) // ーワ
	log.Printf("rune[:5] : %s\n", string([]rune(str)[:5])) // ハローワールド
}

func TestStrCopy(t *testing.T) {
	str := "Hello, World!"
	str2 := str // value copy
	str3 := &str // pointer copy
	str = str[2:3]
	
	log.Printf("str : %s\n", str)
	log.Printf("str2 : %s\n", str2)
	log.Printf("str3 : %s\n", *str3)
}

func TestIndexOf(t *testing.T) {
	str := "Hello, World!"

	log.Printf("index of %s = %d\n", "l", strings.Index(str, "l")) // 2
	log.Printf("index of %s = %d\n", "l", strings.Index(str[3:], "l")) // 0
}

func TestIndexOfMulti(t *testing.T) {
	str := "ハローワールド"

	log.Printf("index of %s = %d\n", "ー", strings.Index(str, "ー")) // not 2
	log.Printf("index of %s = %d\n", "ー", strings.IndexRune(str, 'ー'))
}

func TestRuneOperation(t *testing.T) {
	r := 'a' // this is rune

	// increment
	for i := 0; i < 10; i++ {
		r ++
		log.Printf("r(%v) = %s\n", r, string(r))
	}

	r = 'あ'
	for i := 0; i < 10; i++ {
		r ++
		log.Printf("r(%v) = %s\n", r, string(r))
	}
}

func TestSplit(t *testing.T) {
	str := "Hello, World!"
	strarr := strings.Split(str, ", ")

	log.Printf("Splitted[0] : %s\n", strarr[0])
	log.Printf("Splitted[1] : %s\n", strarr[1])
}

func TestSplitRune(t *testing.T) {
	str := "ハローワールド"
	strarr := strings.FieldsFunc(str, func(r rune) bool {
		return r == 'ー'
	})

	log.Printf("Splitted[0] : %s\n", strarr[0])
	log.Printf("Splitted[1] : %s\n", strarr[1])
	log.Printf("Splitted[1] : %s\n", strarr[2])
}

func TestReplace(t *testing.T) {
	str := "Hello, World!"

	str = strings.Replace(str, "World!", "Universe!", 1)
	log.Printf("Replaced : %s\n", str)

	str = strings.ReplaceAll(str, "l", "L")
	log.Printf("Replaced : %s\n", str)
}
