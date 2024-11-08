package main

import (
	"math/big"
	"strings"
)

type Polynom struct {
	varnum int
	sets   []*big.Int
}

func NewPolynom(varnum int) *Polynom {
	return &Polynom{
		varnum,
		make([]*big.Int, 0),
	}
}

func Copy(p *Polynom) *Polynom {
	new := &Polynom{
		p.varnum,
		make([]*big.Int, len(p.sets)),
	}

	for i := 0; i < len(p.sets); i++ {
		new.sets[i] = big.NewInt(0).Set(p.sets[i])
	}

	return new
}

func (p *Polynom) search_index(num *big.Int) int {
	low, high := 0, len(p.sets)

	for low < high {
		middle := int(uint(low+high) >> 1)

		if p.sets[middle].Cmp(num) < 0 {
			low = middle + 1
		} else {
			high = middle
		}
	}

	return low
}

// insert num, if already exists do nothing
func (p *Polynom) Insert(num *big.Int) {
	index := p.search_index(num)

	if index == len(p.sets) {
		p.sets = append(p.sets, num)
		return
	}

	if p.sets[index].Cmp(num) == 0 {
		return
	}

	p.sets = append(p.sets[:index+1], p.sets[index:]...)
	p.sets[index] = num
}

// insert num, if already exists -- marked it for deletion
func (p *Polynom) Xor(num *big.Int) {
	index := p.search_index(num)

	if index == len(p.sets) {
		p.sets = append(p.sets, num)
		return
	}

	if p.sets[index].Cmp(num) == 0 {
		p.sets = append(p.sets[:index], p.sets[index+1:]...)
		return
	}

	p.sets = append(p.sets[:index+1], p.sets[index:]...)
	p.sets[index] = num
}

// return index if success and -1 if not exist
func (p *Polynom) Search(num *big.Int) int {
	index := p.search_index(num)

	if index == len(p.sets) {
		return -1
	}

	if p.sets[index].Cmp(num) == 0 {
		return index
	}

	return -1
}

func (p *Polynom) FindWeightMax() (int, *big.Int) {
	if len(p.sets) == 0 {
		return -1, nil
	}

	max_weight := 0
	max_index := 0
	for i := range p.sets {
		weight := Weight(p.sets[i])

		if weight > max_weight {
			max_weight = weight
			max_index = i
			continue
		}

		if weight == max_weight && p.sets[max_index].Cmp(p.sets[i]) < 0 {
			max_index = i
		}
	}

	return max_index, p.sets[max_index]
}

func (p *Polynom) XorAllSets() *big.Int {
	xor := big.NewInt(0)

	for i := 0; i < len(p.sets); i++ {
		xor.Xor(xor, p.sets[i])
	}

	return xor
}

func (p *Polynom) String() string {
	if len(p.sets) == 0 {
		return "0"
	}

	sets := []string{}

	for i := 0; i < len(p.sets); i++ {
		sets = append(sets, StringSet(p.varnum, p.sets[i]))
	}

	// for better view
	for i, j := 0, len(sets)-1; i < j; i, j = i+1, j-1 {
		sets[i], sets[j] = sets[j], sets[i]
	}

	return strings.Join(sets, " + ")
}
