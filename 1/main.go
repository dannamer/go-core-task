package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func convertToString(numDecimal int, numOctal int, numHexadecimal int, pi float64, name string, isActive bool, complexNum complex64) string {
	strDecimal := strconv.Itoa(numDecimal)
	strOctal := strconv.FormatInt(int64(numOctal), 8)
	strHexadecimal := strconv.FormatInt(int64(numHexadecimal), 16)
	strPi := strconv.FormatFloat(pi, 'f', -1, 64)
	strName := name
	strIsActive := strconv.FormatBool(isActive)
	strComplex := fmt.Sprintf("%v", complexNum)

	return strDecimal + strOctal + strHexadecimal + strPi + strName + strIsActive + strComplex
}

func stringToRunes(str string) []rune {
	return []rune(str)
}

func hashWithSalt(runes []rune, salt string) string {
	str := string(runes)
	mid := len(str) / 2
	strWithSalt := str[:mid] + salt + str[mid:]

	hash := sha256.New()
	hash.Write([]byte(strWithSalt))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println("Type of numDecimal:", fmt.Sprintf("%T", numDecimal))
	fmt.Println("Type of numOctal:", fmt.Sprintf("%T", numOctal))
	fmt.Println("Type of numHexadecimal:", fmt.Sprintf("%T", numHexadecimal))
	fmt.Println("Type of pi:", fmt.Sprintf("%T", pi))
	fmt.Println("Type of name:", fmt.Sprintf("%T", name))
	fmt.Println("Type of isActive:", fmt.Sprintf("%T", isActive))
	fmt.Println("Type of complexNum:", fmt.Sprintf("%T", complexNum))

	combinedStr := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nCombined string:", combinedStr)

	runes := stringToRunes(combinedStr)
	fmt.Println("\nSlice of runes:", runes)

	salt := "go-2024"
	hashedStr := hashWithSalt(runes, salt)
	fmt.Println("\nHashing result:", hashedStr)
}