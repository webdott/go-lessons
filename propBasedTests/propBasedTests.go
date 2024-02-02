package main

import (
	"strings"
)

type RomanNumErr string

const (
	RomanNumberOverFlowError = RomanNumErr("number provided is greater than 3999")
	InvalidRomanStringError  = RomanNumErr("The roman numeral provided is invalid")
)

type RomanNumerals []RomanNumeral
type windowedRoman string

func (rne RomanNumErr) Error() string {
	return string(rne)
}

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)

	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabicNum uint16) (string, error) {
	if arabicNum > 3999 {
		return "", RomanNumberOverFlowError
	}

	var result strings.Builder

	for i := 0; i < len(allRomanNumerals); {
		romanNumeral := allRomanNumerals[i]

		if arabicNum >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)
			arabicNum -= romanNumeral.Value
		} else {
			i += 1
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) (uint16, error) {
	total := uint16(0)

	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}

	if total > 3999 {
		return 0, InvalidRomanStringError
	}

	return total, nil
}

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
