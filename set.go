package main

import (
	"fmt"
	"math/big"
	"strings"
)

func WeightByte(b byte) byte {
	b = ((b >> 1) & 0x55) + (b & 0x55)
	b = ((b >> 2) & 0x33) + (b & 0x33)
	b = ((b >> 4) & 0x0F) + (b & 0x0F)

	return b
}

// https://habr.com/ru/articles/276957/
func Weight(set *big.Int) int {
	bytes := set.Bytes()

	count := 0
	for i := 0; i < len(bytes); i++ {
		count += int(WeightByte(bytes[i]))
	}

	return count
}

func OnesAndZeros(varnum int, num *big.Int) (ones, zeros []int) {
	ones = []int{}
	zeros = []int{}

	for i := 0; i < varnum; i++ {
		if i >= num.BitLen() {
			zeros = append(zeros, i)
			continue
		}

		if num.Bit(i) == 1 {
			ones = append(ones, i)
		} else {
			zeros = append(zeros, i)
		}
	}

	return
}

func StringSet(varnum int, set *big.Int) string {
	if set.BitLen() == 0 {
		return "1"
	}

	vars := []string{}

	for i := set.BitLen(); i > -1; i-- {
		if set.Bit(i) == 1 {
			vars = append(vars, fmt.Sprintf("x%d", varnum-i))
		}
	}

	return strings.Join(vars, "*")
}
