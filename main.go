package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(10) + 1
	fmt.Println("Угадайте число от 1 до 10! У вас есть 5 попыток")

	reader := bufio.NewReader(os.Stdin)
	maxAttempts := 5
	attempts := 0

	for attempts < maxAttempts {
		fmt.Println("... осталось попыток:", maxAttempts-attempts)
		fmt.Print("Ваш вариант ")

		input, _ := reader.ReadString('\n')
		guess := checkGuess(input)

		if guess == -1 {
			fmt.Print("\nНеобходимо ввести числовое значение от 1 до 10")
			continue
		}

		if guess == target {
			fmt.Printf("Вы угадали! 🎉 Это \"%d\"\n", target)
			break
		}

		if guess > target {
			fmt.Print("\nЗагаданное число меньше!")
		} else {
			fmt.Print("\nЗагаданное число больше!")
		}

		attempts++
	}

	if attempts == maxAttempts {
		fmt.Printf("\nПопытки закончились😢 Было загадано число \"%d\"\n", target)
	}
}

func checkGuess(input string) int {
	input = strings.TrimSpace(input)

	if len(input) > 1 && strings.HasPrefix(input, "0") {
		return -1
	}

	i, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}

	if i < 1 || i > 10 {
		return -1
	}

	return i
}
