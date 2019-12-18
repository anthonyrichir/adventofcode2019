package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
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
		result := ((module - (module%3))/3) - 2
		total += result
	}
	fmt.Println(total)
}
