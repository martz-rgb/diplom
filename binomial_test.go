package main

import (
	"math/big"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

func TestConvert(t *testing.T) {
	n := 10
	degree := 4
	index := 35

	binoms := make([][]int, n)

	for i := 0; i < len(binoms); i++ {
		max := degree + 1
		if i < degree {
			max = i + 1
		}
		binoms[i] = make([]int, max)
	}

	for i := 0; i < degree; i++ {
		for j := 0; j < len(binoms); j++ {
			if i <= j {
				binoms[j][i] = combin.Binomial(j+1, i+1)
			}
		}
	}

	result := big.NewInt(0b0000110101)

	num := convert_index(index, degree, n, binoms)

	if result.Cmp(num) != 0 {
		t.Fatalf("wrong answer, expected: %b, got: %b\n", result, num)
	}
}

func TestConvertLarge(t *testing.T) {
	n := 3
	max := 3 + 3 + 1

	results := []*big.Int{
		big.NewInt(0b001),
		big.NewInt(0b010),
		big.NewInt(0b100),
		big.NewInt(0b011),
		big.NewInt(0b101),
		big.NewInt(0b110),
		big.NewInt(0b111),
	}

	binoms := make([][]int, n)

	for i := 0; i < len(binoms); i++ {
		max := 3 + 1
		if i < 3 {
			max = i + 1
		}
		binoms[i] = make([]int, max)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < len(binoms); j++ {
			if i <= j {
				binoms[j][i] = combin.Binomial(j+1, i+1)
			}
		}
	}

	nums := []*big.Int{}
	cur_degree := 1
	bottom := 0
	floor := combin.Binomial(n, cur_degree)
	for i := 1; i <= max; i++ {
		if i > floor {
			bottom += combin.Binomial(n, cur_degree)
			cur_degree++
			floor += combin.Binomial(n, cur_degree)
		}

		nums = append(nums, convert_index(i-bottom, cur_degree, n, binoms))
	}

	for i := range results {
		if results[i].Cmp(nums[i]) != 0 {
			t.Fatalf("wrong answer, expected: %b, got: %b\n", results[i], nums[i])
		}
	}

}
