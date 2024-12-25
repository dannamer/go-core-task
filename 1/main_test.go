package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestConvertToString(t *testing.T) {
	numDecimal := 42
	numOctal := 052
	numHexadecimal := 0x2A
	pi := 3.14
	name := "Golang"
	isActive := true
	var complexNum complex64 = 1 + 2i

	expected := strconv.Itoa(numDecimal) + strconv.FormatInt(int64(numOctal), 8) + strconv.FormatInt(int64(numHexadecimal), 16) + strconv.FormatFloat(pi, 'f', -1, 64) + name + strconv.FormatBool(isActive) + fmt.Sprintf("%v", complexNum)

	result := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	if result != expected {
		t.Errorf("convertToString() = %v; want %v", result, expected)
	}
}

func TestStringToRunes(t *testing.T) {
	str := "Hello"
	expected := []rune{'H', 'e', 'l', 'l', 'o'}

	result := stringToRunes(str)

	if !equalRunes(result, expected) {
		t.Errorf("stringToRunes() = %v; want %v", result, expected)
	}
}

func equalRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("testString")
	salt := "randomSalt"

	str := string(runes)
	mid := len(str) / 2
	strWithSalt := str[:mid] + salt + str[mid:]

	hash := sha256.New()
	hash.Write([]byte(strWithSalt))
	expectedHash := hex.EncodeToString(hash.Sum(nil))

	result := hashWithSalt(runes, salt)

	if result != expectedHash {
		t.Errorf("hashWithSalt() = %v; want %v", result, expectedHash)
	}
}
