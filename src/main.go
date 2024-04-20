package main

import (
	"fmt"

	sbox "github.com/k1910177/funsec-report1/src/sbox"
	ttable "github.com/k1910177/funsec-report1/src/ttable"
)

func main() {
	input := [16]byte{
		0x05,
		0x02,
		0x03,
		0x04,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x05,
	}

	key := [16]byte{
		0x01,
		0x02,
		0x03,
		0x04,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x01,
	}

	{
		result := sbox.Encrypt(input[:], key[:], 10)

		for i := 0; i < len(result); i++ {
			h := fmt.Sprintf("%x", result[i])
			fmt.Printf("%s", h)
			if (i+1)%4 == 0 {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}

	{
		result := ttable.Encrypt(input[:], key[:], 10)

		for i := 0; i < len(result); i++ {
			h := fmt.Sprintf("%x", result[i])
			fmt.Printf("%s", h)
			if (i+1)%4 == 0 {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}
