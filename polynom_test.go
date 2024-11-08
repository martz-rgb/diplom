package main

import (
	"math/big"
	"testing"
)

func TestPolynomInit(t *testing.T) {
	polynom := NewPolynom(5)

	if polynom.varnum != 5 {
		t.Fatalf("incorrect numbers of variable: want 5, got %d\n", polynom.varnum)
	}
	if len(polynom.sets) != 0 {
		t.Fatalf("sets should be empty: got length %d\n", len(polynom.sets))
	}
}

func TestPolynomInsert(t *testing.T) {
	polynom := NewPolynom(5)

	set := big.NewInt(0b01100)

	polynom.Insert(set)

	if len(polynom.sets) != 1 {
		t.Fatalf("incorrect length: want 1, got %d\n", len(polynom.sets))
	}
	if polynom.sets[0].Cmp(set) != 0 {
		t.Fatalf("incorrect set: want 12, got %d\n", polynom.sets[0])
	}

	// should stay the same
	polynom.Insert(set)

	if len(polynom.sets) != 1 {
		t.Fatalf("incorrect length: want 1, got %d\n", len(polynom.sets))
	}
	if polynom.sets[0].Cmp(set) != 0 {
		t.Fatalf("incorrect set: want 12, got %d\n", polynom.sets[0])
	}
}

func TestPolynomInsertSort(t *testing.T) {
	polynom := NewPolynom(6)

	polynom.Insert(big.NewInt(0b010010))
	polynom.Insert(big.NewInt(0b010101))
	polynom.Insert(big.NewInt(0b111110))
	polynom.Insert(big.NewInt(0b010000))
	polynom.Insert(big.NewInt(0b010010))
	polynom.Insert(big.NewInt(0b010010))
	polynom.Insert(big.NewInt(0b011101))
	polynom.Insert(big.NewInt(0b110101))
	polynom.Insert(big.NewInt(0b101110))
	polynom.Insert(big.NewInt(0b000010))
	polynom.Insert(big.NewInt(0b101111))

	for i := 0; i < len(polynom.sets)-1; i++ {
		if polynom.sets[i].Cmp(polynom.sets[i+1]) > 0 {
			t.Fatalf("unsorted sets")
		}
	}
}

func TestPolynomXor(t *testing.T) {
	polynom := NewPolynom(5)

	set := big.NewInt(0b01100)

	polynom.Xor(set)

	if len(polynom.sets) != 1 {
		t.Fatalf("incorrect length: want 1, got %d\n", len(polynom.sets))
	}
	if polynom.sets[0].Cmp(set) != 0 {
		t.Fatalf("incorrect set: want 12, got %d\n", polynom.sets[0])
	}

	// should remove the set
	polynom.Xor(set)

	if len(polynom.sets) != 0 {
		t.Fatalf("incorrect length: want 0, got %d\n", len(polynom.sets))
	}
}

func TestPolynomXorSort(t *testing.T) {
	polynom := NewPolynom(6)

	polynom.Xor(big.NewInt(0b010010))
	polynom.Xor(big.NewInt(0b010101))
	polynom.Xor(big.NewInt(0b111110))
	polynom.Xor(big.NewInt(0b010000))
	polynom.Xor(big.NewInt(0b010010))
	polynom.Xor(big.NewInt(0b010010))
	polynom.Xor(big.NewInt(0b011101))
	polynom.Xor(big.NewInt(0b110101))
	polynom.Xor(big.NewInt(0b101110))
	polynom.Xor(big.NewInt(0b000010))
	polynom.Xor(big.NewInt(0b101111))

	for i := 0; i < len(polynom.sets)-1; i++ {
		if polynom.sets[i].Cmp(polynom.sets[i+1]) > 0 {
			t.Fatalf("unsorted sets")
		}
	}
}

func TestPolynomCopy(t *testing.T) {
	original := NewPolynom(5)

	original.Insert(big.NewInt(0b00000))
	original.Insert(big.NewInt(0b00010))
	original.Insert(big.NewInt(0b01100))
	original.Insert(big.NewInt(0b11010))
	original.Insert(big.NewInt(0b11110))

	copy := Copy(original)

	if copy == original {
		t.Fatalf("structs should point on different memory addresses\n")
	}

	if original.varnum != copy.varnum {
		t.Fatalf("incorrect number of variables: want %d, got %d\n", original.varnum, copy.varnum)
	}

	if len(original.sets) != len(copy.sets) {
		t.Fatalf("incorrect length of sets: want %d, got %d\n", len(original.sets), len(copy.sets))
	}

	for i := 0; i < len(original.sets); i++ {
		if original.sets[i].Cmp(copy.sets[i]) != 0 {
			t.Fatalf("incorrect copied set: want %d, got %d\n", original.sets[i], copy.sets[i])
		}

		if original.sets[i] == copy.sets[i] {
			t.Fatalf("sets should point on different memory addresses\n")
		}
	}
}

func TestPolynomSearch(t *testing.T) {
	polynom := NewPolynom(3)

	polynom.Insert(big.NewInt(0b011))
	polynom.Insert(big.NewInt(0b001))
	polynom.Insert(big.NewInt(0b110))

	index := polynom.Search(big.NewInt(0b001))
	if index == -1 {
		t.Fatalf("unable to found existing set")
	}
	if index != 0 {
		t.Fatalf("incorrect index: want 0, got %d\n", index)
	}

	index = polynom.Search(big.NewInt(0b111))
	if index != -1 {
		t.Fatalf("found not existing set")
	}
}

func TestPolynomFindWeightMax(t *testing.T) {
	polynom := NewPolynom(3)

	polynom.Insert(big.NewInt(0b011))
	polynom.Insert(big.NewInt(0b100))
	polynom.Insert(big.NewInt(0b110))

	index, num := polynom.FindWeightMax()
	if index != 2 {
		t.Fatalf("incorrect index: want 2, got %d\n", index)
	}
	if num.Cmp((big.NewInt(0b110))) != 0 {
		t.Fatalf("found incorrect set: want 6, got %d\n", num)
	}

	polynom.Xor(big.NewInt(0b110))

	index, num = polynom.FindWeightMax()
	if index != 0 {
		t.Fatalf("incorrect index: want 0, got %d\n", index)
	}
	if num.Cmp((big.NewInt(0b011))) != 0 {
		t.Fatalf("found incorrect set: want 3, got %d\n", num)
	}

	polynom.Xor(big.NewInt(0b011))

	index, num = polynom.FindWeightMax()
	if index != 0 {
		t.Fatalf("incorrect index: want 0, got %d\n", index)
	}
	if num.Cmp((big.NewInt(0b100))) != 0 {
		t.Fatalf("found incorrect set: want 4, got %d\n", num)
	}

	polynom.Xor(big.NewInt(0b100))

	index, num = polynom.FindWeightMax()
	if index != -1 {
		t.Fatalf("found not existing index of max")
	}
	if num != nil {
		t.Fatalf("found not existing max")
	}
}
