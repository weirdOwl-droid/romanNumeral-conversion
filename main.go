package main

import (
	"fmt"
	"strings"
)

var RomanMapping = []struct {
	value  int
	symbol string
}{
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

func romanConversion(input int) (string, error) {
	var result strings.Builder
	if input < 1 || input > 3999 {
		return "", fmt.Errorf("Input Error, hanya mengkonversi angka 1-3999")
	}
	for _, mapping := range RomanMapping {
		for input >= mapping.value {
			result.WriteString(mapping.symbol)
			input -= mapping.value
		}
	}
	return result.String(), fmt.Errorf("")
}

func main() {

	var input int
	var control string
	for {
		fmt.Println("Silahkan masukan angka")
		fmt.Scanln(&input)
		fmt.Println("Hasil konversi . . .")
		fmt.Println(romanConversion(input))
		fmt.Println("Ingin Mengonversi angka lainnya?")
		fmt.Println("Ketik (next) jika ingin melanjutkan")
		fmt.Scanln(&control)
		if control == "next" {
			continue
		} else {
			panic("Closing application")
		}
	}

}
