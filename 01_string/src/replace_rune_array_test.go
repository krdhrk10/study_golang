package test

import (
	"log"
	"testing"
)

func TestReplaceRuneArray(t *testing.T) {
	str := doReplace("ハローワールド、ワーテルローの戦い", "ワー", "わ〜")

	strShouldBe(str, "ハローわ〜ルド、わ〜テルローの戦い")
}

func TestReplaceRuneArrayEmpty(t *testing.T) {
	str := doReplace("", "ワー", "XXX")

	strShouldBe(str, "")
}

func TestReplaceRuneArrayEmptyPlaceholder(t *testing.T) {
	str := doReplace("ハローワールド、ワーテルローの戦い", "", "XXXX")

	strShouldBe(str, "ハローワールド、ワーテルローの戦い")
}

func TestReplaceRuneArrayEmptyReplacer(t *testing.T) {
	str := doReplace("ハローワールド、ワーテルローの戦い", "ワー", "")

	strShouldBe(str, "ハロールド、テルローの戦い")
}

func TestReplaceRuneArrayHeadMatched(t *testing.T) {
	str := doReplace("ハローワールド、ワーテルローの戦い", "ハロー", "もふ")

	strShouldBe(str, "もふワールド、ワーテルローの戦い")
}

func TestReplaceRuneArrayTailMatched(t *testing.T) {
	str := doReplace("ハローワールド、ワーテルローの戦い", "戦い", "ふるもっふ")

	strShouldBe(str, "ハローワールド、ワーテルローのふるもっふ")
}

func TestReplaceRuneArrayFullMatched(t *testing.T) {
	str := doReplace("ハローワールド、ワーテルローの戦い", "ハローワールド、ワーテルローの戦い", "ジョナサン")

	strShouldBe(str, "ジョナサン")
}

func doReplace(str string, placeholder string, replace string) string {
	phrunes := []rune(placeholder)
	if len(phrunes) < 1 {
		return str
	}
	rprunes := []rune(replace)

	nextRuneIdx := 0
	result := []rune{}
	for _, r := range []rune(str) {
		result = append(result, r)
		if r != phrunes[nextRuneIdx] {
			// reset
			nextRuneIdx = 0
			continue
		}

		if nextRuneIdx == len(phrunes)-1 { // finished
			result = result[:len(result)-len(phrunes)]
			result = append(result, rprunes...)

			// reset
			nextRuneIdx = 0
			continue
		}

		nextRuneIdx++
	}
	return string(result)
}

func strShouldBe(str string, shouldBe string) {
	if str != shouldBe {
		log.Fatalf("Not matched  (%s, %s)", str, shouldBe)
	}
}
