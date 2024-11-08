package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	interactive_mode()
}

// Программа является реализацией алгоритма, описанного в https://m.mathnet.ru/php/archive.phtml?wshow=paper&jrnid=dm&paperid=1658&option_lang=rus
func interactive_mode() {
	/* на вход программа будет принимать число переменных n и множество S(f) через входной поток данных,
	где корректный вид n — натуральное число, S(f) — некоторое количество строк,
	которые представляют двоичную\восьмиричную\десятичную\шестнадцатиричную запись набора*/

	log_file, err := os.Create("./log/log.txt")
	if err != nil {
		panic(err)
	}
	log.SetOutput(log_file)

	base := 2
	if len(os.Args) > 1 {
		arg_base, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Printf("unable to parse argument to base number system: %s; using binary number system as default.\n", os.Args[1])
		} else {
			base = arg_base
		}
	}

	fmt.Println("Write down number or variables:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	ok := scanner.Scan()
	if !ok {
		if err := scanner.Err(); err != nil {
			log.Printf("error occured while reading from stdin: %s\n", err)
		}
		return
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Printf("unable to parse number of variables: %s\n", scanner.Text())
		return
	}
	varnum := num

	max_num := big.NewInt(0)
	max_num.Exp(big.NewInt(2), big.NewInt(int64(varnum)), nil)

	polynom := NewPolynom(varnum)

	fmt.Println("Write down sets:")

	for scanner.Scan() {
		set := big.NewInt(0)
		set, ok := set.SetString(scanner.Text(), base)
		if !ok {
			log.Printf("unable to parse set of base %d: %s\n", base, scanner.Text())
			continue
		}

		if max_num.Cmp(set) < 0 {
			log.Printf("too large set for %d variables: %s\n", varnum, set.Text(2))
			continue
		}

		polynom.Insert(set)
	}

	fmt.Println("polynom: ", polynom)

	basis, err := polynom.FindPeriods()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(basis) == 0 {
		fmt.Print("basis: none")
	}

	// print basis
	fmt.Println("basis: ")
	for i := 0; i < len(basis); i++ {
		fmt.Printf("\t%0*s\n", polynom.varnum, basis[i].Text(2))
	}
}

func benchmark_all_mode() {
	from_n := 3
	to_n := 4

	average2 := []time.Duration{}
	average3 := []time.Duration{}

	for n := from_n; n <= to_n; n++ {
		d2 := degree_benchmark_all(2, n)
		if !d2.IsInt64() {
			fmt.Printf("%d %d unable to represent as int64: %s", n, 2, d2.Text(10))
		} else {
			average2 = append(average2, time.Duration(d2.Int64()))
		}

		d3 := degree_benchmark_all(3, n)
		if !d3.IsInt64() {
			fmt.Printf("%d %d unable to represent as int64: %s", n, 3, d3.Text(10))
		} else {
			average3 = append(average3, time.Duration(d3.Int64()))
		}
	}

	file, err := os.Create(fmt.Sprintf("result-all-%d-%d-%d.csv", from_n, to_n, time.Now().Unix()))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"n", "degree", "duration_string", "duration"}
	writer.Write(headers)

	for i, duration := range average2 {
		var record []string

		// n
		record = append(record, strconv.Itoa(i+from_n))

		// degree
		record = append(record, "2")

		// duration
		record = append(record, duration.String())
		record = append(record, strconv.FormatInt(duration.Nanoseconds(), 10))

		writer.Write(record)
	}

	for i, duration := range average3 {
		var record []string

		// n
		record = append(record, strconv.Itoa(i+from_n))

		// degree
		record = append(record, "3")

		// duration
		record = append(record, duration.String())
		record = append(record, strconv.FormatInt(duration.Nanoseconds(), 10))

		writer.Write(record)
	}
}

func benchmark_mode() {
	degree := 3
	n := 50
	count := 1000

	_, results := degree_benchmark(degree, n, count)

	file, err := os.Create(fmt.Sprintf("result-%d-%d-%d-%d.csv", degree, n, count, time.Now().Unix()))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"index", "duration_string", "duration", "memory", "periods", "polynom"}
	writer.Write(headers)

	for i, result := range results {
		var record []string

		// index
		record = append(record, strconv.Itoa(i+1))

		// duration
		record = append(record, result.duration.String())
		record = append(record, strconv.FormatInt(result.duration.Nanoseconds(), 10))

		// memory
		record = append(record, strconv.FormatUint(result.memory, 10))

		// periods
		periods := []string{}

		for i := 0; i < len(result.results); i++ {
			periods = append(periods, fmt.Sprintf("%0*s", result.polynom.varnum, result.results[i].Text(2)))
		}

		record = append(record, strings.Join(periods, " "))

		// polynom
		record = append(record, result.polynom.String())

		writer.Write(record)
	}
}
