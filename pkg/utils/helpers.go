package utils

import "fmt"

func RGBtoHEX(r, g, b, a int) string {
	return "#" + DecToHex(r) + DecToHex(g) + DecToHex(b) + DecToHex(a)
}

func DecToHex(n int) string {
	hex := ""
	nm := make(map[int]string, 0)
	for i := 0; i < 16; i++ {
		if i > 9 {
			nm[i] = string(i%10 + 65)
		} else {
			nm[i] = fmt.Sprintf("%d", i)
		}
	}
	for n != 0 {
		hex = nm[n%16] + hex
		n = n / 16
	}
	return fmt.Sprintf("%02s", hex)
}
