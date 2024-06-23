package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Sum all the values in the buffer and return the final float64 sum
func sumBuffer(buf []int, maxExp int) *big.Float {
	var result = new(big.Float).SetPrec(1024)
	sumValue := new(big.Int)
	var i int
	for i = 1; len(buf)-i >= 0; i++ {
		value := big.NewInt(int64(buf[len(buf)-i]))
		bigIntExp := value.Mul(value, big.NewInt(int64(math.Pow10(i-1))))
		sumValue = sumValue.Add(sumValue, bigIntExp)
	}
	bigFloatSum := new(big.Float).SetInt(sumValue).SetPrec(1024)
	bigFloatMaxExp := new(big.Float).SetFloat64(math.Pow10(maxExp)).SetPrec(1024)
	result = new(big.Float).Quo(bigFloatSum, bigFloatMaxExp).SetPrec(1024)
	return result
}

func readFloatsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []string
	scanner := bufio.NewScanner(file)
	// Чтение строк из файла
	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "Имя файла для чтения чисел с плавающей запятой")
	flag.Parse()

	// Проверка, что имя файла указано
	if filename == "" {
		fmt.Println("Имя файла не указано. Используйте флаг -file <имя файла>")
		return
	}

	// Чтение чисел из файла
	numbers, err := readFloatsFromFile(filename)
	if err != nil {
		fmt.Printf("Ошибка чтения файла %s: %v\n", filename, err)
		return
	}
	buf := make([]int, 0)
	maxExp := 0

	for _, num := range numbers {

		parts := strings.Split(num, ".")
		integerPart := parts[0]
		decimalPart := parts[1]
		expIndex := len(decimalPart)
		if expIndex > maxExp {
			maxExp = expIndex
		}
		for len(buf) <= expIndex {
			buf = append(buf, 0)
		}
		firstChar := string(integerPart[0])
		if firstChar == "-" {
			decimalPart = "-" + decimalPart
		}
		var intPart int
		intPart, _ = strconv.Atoi(integerPart)
		buf[0] += intPart
		var decPart int
		decPart, _ = strconv.Atoi(decimalPart)
		buf[expIndex] += decPart
	}

	sum := sumBuffer(buf, maxExp).Text('f', -1)
	fmt.Printf("Sum: %s\n", sum) // Should be the sum of the numbers
}
