package main

import (
	"fmt"
	"math/big"
)

func calc(a, b *big.Int) {

	sum := new(big.Int)
	difference := new(big.Int)
	product := new(big.Int)
	quotient := new(big.Int)
	remainder := new(big.Int)

	// Сложение
	sum.Add(a, b)
	fmt.Printf("Сумма: %s\n", sum.String())

	// Вычитание
	difference.Sub(a, b)
	fmt.Printf("Вычитание: %s\n", difference.String())

	// Умножение
	product.Mul(a, b)
	fmt.Printf("Умножение: %s\n", product.String())

	// Деление с остатком
	quotient.DivMod(b, a, remainder)
	fmt.Printf("Деление: %s\n", quotient.String())
	fmt.Printf("Остаток от деления: %s\n", remainder.String())
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("2097153", 10)
	calc(a, b)
}
