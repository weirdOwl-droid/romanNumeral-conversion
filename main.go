package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanConversion(input int) (string, error) {
	var RomanMapping = []struct {
		value  int
		symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

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
	return result.String(), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Silahkan masukan angka:")
		inputStr, _ := reader.ReadString('\n')
		inputStr = strings.TrimSpace(inputStr)

		input, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Error: Input tidak valid. Masukkan angka numerik.")
			continue
		}

		if value, err := romanConversion(input); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Hasil konversi:", value)
		}

		fmt.Println("Ingin Mengonversi angka lainnya? (ketik 'next' untuk melanjutkan, atau ketik apa saja untuk keluar)")
		control, _ := reader.ReadString('\n')
		control = strings.TrimSpace(control)
		if control != "next" {
			fmt.Println("Menutup aplikasi...")
			break
		}
	}
}
