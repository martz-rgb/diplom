package main

import (
	"math/big"
	"testing"
)

func TestWeightByte(t *testing.T) {
	b := byte(0)
	weight := WeightByte(b)

	if weight != 0 {
		t.Fatalf("%b: wrong weight: want 0, got %d\n", b, weight)
	}

	b = byte(0b00011001)
	weight = WeightByte(b)

	if weight != 3 {
		t.Fatalf("%b: wrong weight: want 3, got %d\n", b, weight)
	}

	b = byte(0b11111111)
	weight = WeightByte(b)

	if weight != 8 {
		t.Fatalf("%b: wrong weight: want 8, got %d\n", b, weight)
	}
}

func TestSetWeight(t *testing.T) {
	num, _ := new(big.Int).SetString("0", 16)
	weight := Weight(num)

	if weight != 0 {
		t.Fatalf("%x: wrong weight: want 0, got %d\n", num, weight)
	}

	num, _ = new(big.Int).SetString("1101110000010011100", 2)
	weight = Weight(num)

	if weight != 9 {
		t.Fatalf("%x: wrong weight: want 9, got %d\n", num, weight)
	}

	// 1011111100110101000000101001110000100001010011100111111111111111111
	num, _ = new(big.Int).SetString("5F9A814E10A73FFFF", 16)
	weight = Weight(num)

	if weight != 40 {
		t.Fatalf("%x: wrong weight: want 40, got %d\n", num, weight)
	}
}

func TestSetOnesAndZeros(t *testing.T) {
	varnum := 3
	num := big.NewInt(0)

	ones, zeros := OnesAndZeros(varnum, num)
	if len(ones) != 0 {
		t.Fatalf("wrong len of ones: want 0, got %d\n", len(ones))
	}
	if len(zeros) != 3 {
		t.Fatalf("wrong len of zeros: want 3, got %d\n", len(ones))
	}
	for i := 0; i < varnum; i++ {
		if zeros[i] != i {
			t.Fatalf("wrong indices")
		}
	}

	num = big.NewInt(0b011)

	ones, zeros = OnesAndZeros(varnum, num)
	if len(ones) != 2 {
		t.Fatalf("wrong len of ones: want 4, got %d\n", len(ones))
	}
	if len(zeros) != 1 {
		t.Fatalf("wrong len of zeros: want 3, got %d\n", len(ones))
	}
	if ones[0] != 0 {
		t.Fatalf("wrong one index: want 0, got %d\n", zeros[0])
	}
	if ones[1] != 1 {
		t.Fatalf("wrong one index: want 1, got %d\n", zeros[0])
	}
	if zeros[0] != 2 {
		t.Fatalf("wrong zero index: want 2, got %d\n", zeros[0])
	}
}
