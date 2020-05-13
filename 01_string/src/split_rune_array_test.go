package test

import (
	"log"
	"testing"
)

func TestSplitRuneArray(t *testing.T) {
	strarr := doSplit("ハローワールド、ワーテルローの戦い", "ワー")

	strarrShouldBe(strarr, []string{
		"ハロー",
		"ルド、",
		"テルローの戦い",
	})
}

func TestSplitRuneArrayEmpty(t *testing.T) {
	strarr := doSplit("", "ワー")

	strarrShouldBe(strarr, []string{})
}

func TestSplitRuneArrayEmptySplitter(t *testing.T) {
	strarr := doSplit("ハローワールド、ワーテルローの戦い", "")

	strarrShouldBe(strarr, []string{
		"ハローワールド、ワーテルローの戦い",
	})
}

func TestSplitRuneArrayEmptyHeadMatched(t *testing.T) {
	strarr := doSplit("ハローワールド、ワーテルローの戦い", "ハロー")

	strarrShouldBe(strarr, []string{
		"ワールド、ワーテルローの戦い",
	})
}

func TestSplitRuneArrayEmptyTailMatched(t *testing.T) {
	strarr := doSplit("ハローワールド、ワーテルローの戦い", "戦い")

	strarrShouldBe(strarr, []string{
		"ハローワールド、ワーテルローの",
	})
}

func TestSplitRuneArrayFullMatched(t *testing.T) {
	strarr := doSplit("ハローワールド、ワーテルローの戦い", "ハローワールド、ワーテルローの戦い")

	strarrShouldBe(strarr, []string{})
}

func doSplit(str string, splitter string) []string {
	splitrunes := []rune(splitter)
	if len(splitrunes) < 1 {
		return []string{str}
	}

	runeBuf := []rune{}
	nextRuneIdx := 0

	result := []string{}
	for _, r := range []rune(str) {
		runeBuf = append(runeBuf, r)
		if r != splitrunes[nextRuneIdx] {
			// reset
			nextRuneIdx = 0
			continue
		}

		if nextRuneIdx == len(splitrunes)-1 { // finished
			if 0 < len(runeBuf)-len(splitrunes) {
				result = append(result, string(runeBuf[:len(runeBuf)-len(splitrunes)]))
			}

			// reset
			runeBuf = []rune{}
			nextRuneIdx = 0
			continue
		}

		nextRuneIdx++
	}
	if 0 < len(runeBuf) {
		result = append(result, string(runeBuf))
	}
	return result
}

func strarrShouldBe(strarr []string, shouldBe []string) {
	if len(strarr) != len(shouldBe) {
		log.Fatalf("Not matched length (%q, %q)", strarr, shouldBe)
	}

	for i, s := range strarr {
		if s != shouldBe[i] {
			log.Fatalf("Not matched index(%d): (%s, %s)", i, s, shouldBe[i])
		}
	}
}
