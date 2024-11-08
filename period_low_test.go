package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

func BenchmarkLinear(b *testing.B) {
	degree := 1

	n := 5
	max_sets := 0
	for i := 1; i <= degree; i++ {
		max_sets += combin.Binomial(n, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Intn(max_sets)

		p := NewPolynom(n)

		c := 0
		for c < count {
			set := big.NewInt(0)

			for j := 0; j < degree; j++ {
				set.SetBit(set, rand.Intn(n), 1)
			}

			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		b.StartTimer()

		p.FindPeriods()
	}
}

func BenchmarkQuadric(b *testing.B) {
	degree := 2

	n := 10
	max_sets := 0
	for i := 1; i <= degree; i++ {
		max_sets += combin.Binomial(n, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Intn(max_sets)

		p := NewPolynom(n)

		c := 0
		for c < count {
			set := big.NewInt(0)

			for j := 0; j < degree; j++ {
				set.SetBit(set, rand.Intn(n), 1)
			}

			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		b.StartTimer()

		p.FindPeriods()
	}
}

func BenchmarkTriple(b *testing.B) {
	degree := 3

	n := 5
	max_sets := 0
	for i := 1; i <= degree; i++ {
		max_sets += combin.Binomial(n, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		count := rand.Intn(max_sets)

		p := NewPolynom(n)

		c := 0
		for c < count {
			set := big.NewInt(0)

			for j := 0; j < degree; j++ {
				set.SetBit(set, rand.Intn(n), 1)
			}

			if p.Search(set) == -1 {
				p.Insert(set)
				c++
			}
		}

		fmt.Println(p)

		b.StartTimer()

		periods, err := p.FindPeriods()

		b.StopTimer()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(periods)
	}
}
