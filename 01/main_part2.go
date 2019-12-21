package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	total := 0
	file, err := os.Open("modules.txt")
	if (err != nil) {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		module,err := strconv.Atoi(scanner.Text())
		if (err != nil) {
			log.Fatal(err)
		}
		total += fuel(module)
	}
	fmt.Println(total)
}

func fuel(mass int) int {
	result := ((mass - (mass%3))/3) - 2
	if (result <= 0) {
		return 0
	} else {
		return result + fuel(result)
	}
}
