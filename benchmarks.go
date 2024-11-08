package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"runtime"
	"time"

	"gonum.org/v1/gonum/stat/combin"
)

type Result struct {
	polynom  *Polynom
	results  []*big.Int
	duration time.Duration
	memory   uint64
}

type ResultSmall struct {
	duration time.Duration
	memory   uint64
}

func degree_benchmark(degree int, n int, count int) ([]time.Duration, []Result) {
	max_sets := 0
	for i := 1; i <= degree; i++ {
		max_sets += combin.Binomial(n, i)
	}

	results := []Result{}
	durations := []time.Duration{}

	for i := 0; i < count; i++ {
		var count int
		for count == 0 {
			count = rand.Intn(max_sets)
		}

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

		var m1, m2 runtime.MemStats
		var start time.Time
		var elapsed time.Duration

		runtime.ReadMemStats(&m1)
		start = time.Now()

		periods, err := p.FindPeriods()

		elapsed = time.Since(start)
		runtime.ReadMemStats(&m2)

		if err != nil {
			panic(err)
		}

		durations = append(durations, elapsed)
		results = append(results, Result{
			p,
			periods,
			elapsed,
			m2.TotalAlloc - m1.TotalAlloc,
		})

		fmt.Printf("%d %s\n", i+1, elapsed.String())
	}

	return durations, results
}

// https://www.mathnet.ru/links/e23e2c835202c5f1f0d9a8ad8a19e1df/zvmmf7536.pdf
func convert_index(index int, degree int, n int, binoms [][]int) *big.Int {
	num := big.NewInt(0)

	d := index
	l := 0
	for i := 0; i < n; i++ {
		var c int
		if degree-l-1 == 0 {
			c = 1
		} else if degree-l-1 < 0 {
			c = 0
		} else {
			c = binoms[n-i-1-1][degree-l-1-1]
		}
		a := d - c

		if a > 0 {
			// bit is zero
			d = a
		} else {
			num.SetBit(num, i, 1)
			l += 1
		}
	}

	return num
}

func diff_monoms(ordinal *big.Int) (result []int) {
	new := &big.Int{}
	new.Add(ordinal, big.NewInt(1))

	for i := 0; i < ordinal.BitLen(); i++ {
		if new.Bit(i) != ordinal.Bit(i) {
			result = append(result, i)
		}
	}

	if new.BitLen() > ordinal.BitLen() {
		result = append(result, ordinal.BitLen())
	}

	return
}

func degree_benchmark_all(degree int, n int) *big.Int {
	binoms := make([][]int, n)
	floors := make([]int, degree)

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

		if i == 0 {
			floors[i] = binoms[n-1][i]
		} else {
			floors[i] = floors[i-1] + binoms[n-1][i]
		}
	}

	polynom_ordinal := big.NewInt(1)
	polynom_ordinal.Lsh(polynom_ordinal, uint(floors[degree-1]-(binoms[n-1][degree-1])))

	//results := []ResultSmall{}
	duration := big.NewInt(0)
	polynom := NewPolynom(n)

	for polynom_ordinal.BitLen() <= floors[degree-1] {
		// var m1, m2 runtime.MemStats
		var start time.Time
		var elapsed time.Duration

		// runtime.ReadMemStats(&m1)
		start = time.Now()

		_, err := polynom.FindPeriods()

		elapsed = time.Since(start)
		// runtime.ReadMemStats(&m2)

		if err != nil {
			panic(err)
		}

		duration.Add(duration, big.NewInt(elapsed.Nanoseconds()))
		// results = append(results, ResultSmall{
		// 	elapsed,
		// 	m2.TotalAlloc - m1.TotalAlloc,
		// })

		monoms := diff_monoms(polynom_ordinal)
		polynom_ordinal.Add(polynom_ordinal, big.NewInt(1))

		for i := 0; i < len(monoms); i++ {

			floor := 0
			for j := range floors {
				if monoms[i] >= floors[j] {
					floor = j
				} else {
					break
				}
			}

			polynom.Xor(convert_index(i+1-floors[floor], floor, n, binoms))
		}

	}
	fmt.Printf("%d\n", polynom_ordinal)

	return duration.Div(duration, polynom_ordinal)
}
