package main

import (
	"math/big"
)

func (p *Polynom) FindPeriods() ([]*big.Int, error) {
	system := []*Polynom{}

	next := Copy(p)

	for len(next.sets) > 0 {
		linears := find_linears(next)
		if len(linears) == 0 {
			break
		}

		system = append(system, linears...)

		product := product_linears(linears)

		for i := 0; i < len(product.sets); i++ {
			next.Xor(product.sets[i])
		}
	}

	matrix := convert_system_to_matrix(system)
	matrix = GaussianElimination(next.varnum, matrix)

	if len(matrix) > next.varnum { // only trivial solution
		return []*big.Int{big.NewInt(0)}, nil
	}

	return extract_basis(next.varnum, matrix), nil
}

func find_linears(p *Polynom) (linears []*Polynom) {
	_, max := p.FindWeightMax()

	ones, zeros := OnesAndZeros(p.varnum, max)

	linears = []*Polynom{}

	for i := 0; i < len(ones); i++ {
		linear := NewPolynom(p.varnum)

		set := new(big.Int)
		set.SetBit(set, ones[i], 1)

		linear.Insert(set)

		for j := 0; j < len(zeros); j++ {
			check := new(big.Int).Set(max)

			check.SetBit(check, ones[i], 0)
			check.SetBit(check, zeros[j], 1)

			index := p.Search(check)
			if index != -1 {
				set := new(big.Int)
				set.SetBit(set, zeros[j], 1)

				linear.Insert(set)
			}
		}

		linears = append(linears, linear)
	}

	return
}

func product_linears(linears []*Polynom) (product *Polynom) {
	product = linears[0]

	for i := 1; i < len(linears); i++ {
		next_product := NewPolynom(product.varnum)

		for j := 0; j < len(linears[i].sets); j++ {
			for k := 0; k < len(product.sets); k++ {
				new := new(big.Int).Set(product.sets[k])
				new.Or(new, linears[i].sets[j])
				next_product.Xor(new)
			}
		}

		product = next_product
	}

	return
}

func convert_system_to_matrix(system []*Polynom) []*big.Int {
	matrix := make([]*big.Int, len(system))

	for i := 0; i < len(system); i++ {
		matrix[i] = system[i].XorAllSets()
	}

	return matrix
}

func GaussianElimination(cols int, matrix []*big.Int) []*big.Int {
	for i := cols - 1; i > -1; i-- {
		row := cols - i - 1

		if row >= len(matrix) {
			break
		}
		// find first one
		if matrix[row].Bit(i) != 1 {
			found := false

			for j := row + 1; j < len(matrix); j++ {
				if matrix[j].Bit(i) == 1 {
					matrix[row], matrix[j] = matrix[j], matrix[row]
					found = true
					break
				}
			}

			if !found {
				empty := big.NewInt(0)
				empty.SetBit(empty, cols, 1)

				matrix = append(matrix[:row+1], matrix[row:]...)
				matrix[row] = empty
				continue
			}
		}

		for j := 0; j < len(matrix); j++ {
			if j == row {
				// skip pivot
				continue
			}

			if matrix[j].Bit(i) == 1 {
				matrix[j].Xor(matrix[j], matrix[row])
			}
		}
	}

	// delete zeros rows
	for i := len(matrix) - 1; i > -1; i-- {
		if matrix[i].BitLen() == 0 {
			matrix = append(matrix[:i], matrix[i+1:]...)
		}
	}

	for len(matrix) != cols {
		empty := big.NewInt(0)
		empty.SetBit(empty, cols, 1)

		matrix = append(matrix, empty)
	}

	return matrix
}

func extract_basis(varnum int, matrix []*big.Int) (basis []*big.Int) {
	basis = []*big.Int{}

	col := varnum - 1
	for col > -1 {
		row := varnum - col - 1
		if row >= len(matrix) || matrix[row].Bit(col) == 0 {
			period := big.NewInt(0)

			for i := 0; i < row; i++ {
				period.SetBit(period, varnum-1-i, matrix[i].Bit(col))
			}
			for i := row; i < varnum; i++ {
				if i == row {
					period.SetBit(period, col, 1)
				}
			}

			basis = append(basis, period)
		}
		col--
	}

	return
}
