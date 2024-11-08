package main

import (
	"math"
	"math/big"
	"math/rand"
	"slices"
	"testing"
)

func TestPeriodNone(t *testing.T) {
	polynom := NewPolynom(4)

	polynom.Insert(big.NewInt(0b1011))
	polynom.Insert(big.NewInt(0b0111))
	polynom.Insert(big.NewInt(0b1010))
	polynom.Insert(big.NewInt(0b0011))
	polynom.Insert(big.NewInt(0b0100))

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	if len(period) != 0 {
		t.Fatalf("incorrect length: want 1, got %d", len(period))
	}
}

func TestPeriodSymmetrical(t *testing.T) {
	polynom := NewPolynom(4)

	polynom.Insert(big.NewInt(0b1110))
	polynom.Insert(big.NewInt(0b1101))
	polynom.Insert(big.NewInt(0b1100))
	polynom.Insert(big.NewInt(0b1011))
	polynom.Insert(big.NewInt(0b0111))
	polynom.Insert(big.NewInt(0b0110))
	polynom.Insert(big.NewInt(0b0101))
	polynom.Insert(big.NewInt(0b0001))
	polynom.Insert(big.NewInt(0b0000))

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	if len(period) != 1 {
		t.Fatalf("incorrect length: want 2, got %d", len(period))
	}

	if period[0].Cmp(big.NewInt(0b1111)) != 0 {
		t.Fatalf("incorrect period: want 15, got %d", period[0])
	}
}

func TestPeriodDifferent1(t *testing.T) {
	polynom := NewPolynom(5)

	polynom.Insert(big.NewInt(0b10100))
	polynom.Insert(big.NewInt(0b10000))
	polynom.Insert(big.NewInt(0b01100))
	polynom.Insert(big.NewInt(0b00110))
	polynom.Insert(big.NewInt(0b00001))

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	if len(period) != 2 {
		t.Fatalf("incorrect length: want 3, got %d", len(period))
	}

	if period[0].Cmp(big.NewInt(0b11001)) != 0 && period[1].Cmp(big.NewInt(0b11001)) != 0 {
		t.Fatalf("no period: want 25, got %d and %d", period[0], period[1])
	}

	if period[0].Cmp(big.NewInt(0b01010)) != 0 && period[1].Cmp(big.NewInt(0b01010)) != 0 {
		t.Fatalf("no period: want 10, got %d and %d", period[0], period[1])
	}
}

func TestPeriodDifferent2(t *testing.T) {
	polynom := NewPolynom(5)

	polynom.Insert(big.NewInt(0b10010))
	polynom.Insert(big.NewInt(0b10000))
	polynom.Insert(big.NewInt(0b01010))
	polynom.Insert(big.NewInt(0b00110))
	polynom.Insert(big.NewInt(0b00001))

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	if len(period) != 2 {
		t.Fatalf("incorrect length: want 3, got %d", len(period))
	}

	if period[0].Cmp(big.NewInt(0b11001)) != 0 && period[1].Cmp(big.NewInt(0b11001)) != 0 {
		t.Fatalf("no period: want 25, got %d and %d", period[0], period[1])
	}

	if period[0].Cmp(big.NewInt(0b01100)) != 0 && period[1].Cmp(big.NewInt(0b01100)) != 0 {
		t.Fatalf("no period: want 10, got %d and %d", period[0], period[1])
	}
}

func TestPeriodDifferent3(t *testing.T) {
	polynom := NewPolynom(7)

	polynom.Insert(big.NewInt(0b1100000))
	polynom.Insert(big.NewInt(0b1010000))
	polynom.Insert(big.NewInt(0b0100100))
	polynom.Insert(big.NewInt(0b0010100))
	polynom.Insert(big.NewInt(0b0100001))
	polynom.Insert(big.NewInt(0b0010001))
	polynom.Insert(big.NewInt(0b0001110))

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	if len(period) != 2 {
		t.Fatalf("incorrect length: want 3, got %d", len(period))
	}

	if period[0].Cmp(big.NewInt(0b0110000)) != 0 && period[1].Cmp(big.NewInt(0b1000001)) != 0 {
		t.Fatalf("no period: want 25, got %d and %d", period[0], period[1])
	}

	if period[0].Cmp(big.NewInt(0b0110000)) != 0 && period[1].Cmp(big.NewInt(0b1000001)) != 0 {
		t.Fatalf("no period: want 10, got %d and %d", period[0], period[1])
	}
}

func TestPeriodDifferent4(t *testing.T) {
	polynom := NewPolynom(7)

	polynom.Insert(big.NewInt(0b1001000))
	polynom.Insert(big.NewInt(0b0101000))
	polynom.Insert(big.NewInt(0b0001100))
	polynom.Insert(big.NewInt(0b1010000))
	polynom.Insert(big.NewInt(0b1000100))
	polynom.Insert(big.NewInt(0b0010010))
	polynom.Insert(big.NewInt(0b0000110))
	polynom.Insert(big.NewInt(0b1000000))
	polynom.Insert(big.NewInt(0b0000001))

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	if len(period) != 2 {
		t.Fatalf("incorrect length: want 3, got %d", len(period))
	}

	if period[0].Cmp(big.NewInt(0b0110100)) != 0 && period[1].Cmp(big.NewInt(0b1100011)) != 0 {
		t.Fatalf("no period: want 25, got %d and %d", period[0], period[1])
	}

	if period[0].Cmp(big.NewInt(0b0110100)) != 0 && period[1].Cmp(big.NewInt(0b1100011)) != 0 {
		t.Fatalf("no period: want 10, got %d and %d", period[0], period[1])
	}
}

func TestPeriodDifferent5(t *testing.T) {
	polynom := NewPolynom(7)

	polynom.Insert(big.NewInt(0b1110000))
	polynom.Insert(big.NewInt(0b1101000))
	polynom.Insert(big.NewInt(0b0110001))
	polynom.Insert(big.NewInt(0b0101001))
	polynom.Insert(big.NewInt(0b0000010))

	space := []*big.Int{
		big.NewInt(0b1000001),
		big.NewInt(0b0011000),
		big.NewInt(0b0000100)}

	slices.SortFunc(space, func(a, b *big.Int) int {
		return a.Cmp(b)
	})

	period, err := polynom.FindPeriods()
	if err != nil {
		t.Fatal("error occured: ", err)
	}

	slices.SortFunc(period, func(a, b *big.Int) int {
		return a.Cmp(b)
	})

	if !slices.EqualFunc(space, period, func(a, b *big.Int) bool {
		return a.Cmp(b) == 0
	}) {
		t.Fatalf("incorrect periods")
	}
}

func TestRandom(t *testing.T) {
	n := 4
	pow := int64(math.Pow(2, float64(n)))

	for count := int64(2); count < pow; count++ {

		arr := []*big.Int{}
		for i := int64(0); i < count; i++ {
			arr = append(arr, big.NewInt(i))
		}

		for arr[0].Cmp(big.NewInt(pow-1)) < 0 {

		}

		p := NewPolynom(n)
		c := int64(0)
		for c < count {
			set := big.NewInt(rand.Int63n(pow))
			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		periods, err := p.FindPeriods()
		if err != nil {
			t.Fatal("error occured: ", err)
		}

		if len(periods) == 1 {
			vars := big.NewInt(0)
			for _, set := range p.sets {
				vars.Or(vars, set)
			}

			if Weight(vars) == n {
				t.Log(p)
			}
		}

		if len(periods) > 2 {
			// only_missed := true
			// for _, period := range periods {
			// 	if Weight(period) > 1 {
			// 		only_missed = false
			// 		break
			// 	}
			// }

			// two_or_more := true
			// for _, period := range periods {
			// 	if Weight(period) < 2 {
			// 		two_or_more = false
			// 	}
			// }

			// if two_or_more {
			// 	t.Log(p)
			// 	for i := 0; i < len(periods); i++ {
			// 		t.Logf("\t%0*s\n", p.varnum, periods[i].Text(2))
			// 	}
			// }
		}
	}

}

func BenchmarkRandom5(b *testing.B) {
	n := 5
	pow := int64(math.Pow(2, float64(n)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Int63n(pow)
		if count == 0 {
			count = 1
		}

		p := NewPolynom(n)
		c := int64(0)
		for c < count {
			set := big.NewInt(rand.Int63n(pow))
			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		b.StartTimer()

		p.FindPeriods()
	}
}

func BenchmarkRandom7(b *testing.B) {
	n := 7
	pow := int64(math.Pow(2, float64(n)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Int63n(pow)
		if count == 0 {
			count = 1
		}

		p := NewPolynom(n)
		c := int64(0)
		for c < count {
			set := big.NewInt(rand.Int63n(pow))
			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		b.StartTimer()

		p.FindPeriods()
	}
}

func BenchmarkRandom10(b *testing.B) {
	n := 10
	pow := int64(math.Pow(2, float64(n)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Int63n(pow)
		if count == 0 {
			count = 1
		}

		p := NewPolynom(n)
		c := int64(0)
		for c < count {
			set := big.NewInt(rand.Int63n(pow))
			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		b.StartTimer()

		p.FindPeriods()
	}
}

func BenchmarkRandom15(b *testing.B) {
	n := 15
	pow := int64(math.Pow(2, float64(n)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Int63n(pow)
		if count == 0 {
			count = 1
		}

		p := NewPolynom(n)
		c := int64(0)
		for c < count {
			set := big.NewInt(rand.Int63n(pow))
			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		b.StartTimer()

		p.FindPeriods()
	}
}
