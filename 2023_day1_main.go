package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractNumbers(str string) []string {
	re := regexp.MustCompile(`\d+`)
	return re.FindAllString(str, -1)
}

func extractFirstAndLastNumber(str string) (string, string) {
	re := regexp.MustCompile(`\d`)
	allNumbers := re.FindAllString(str, -1)

	if len(allNumbers) == 0 {
		return "0", "0"
	} else {
		return allNumbers[0], allNumbers[len(allNumbers)-1]
	}
}

func replaceStrNumber(str string) {
	writtenNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var x string = str

	for wn := 0; wn < len(writtenNumbers); wn++ {
		strings.ReplaceAll(x, writtenNumbers[wn], numbers[wn])
	}
}

func runPart1() {
	file, err := os.Open("assets/2023_day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 30000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes\n", count)

	dataString := string(data[:count])
	dataList := strings.Split(dataString, "\n")
	total := 0

	for i := 0; i < len(dataList); i++ {
		firstNumber, lastNumber := extractFirstAndLastNumber(dataList[i])

		n, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			log.Fatal(err)
		}

		total += n
	}

	fmt.Println("TOTAL:", total)
}

func main() {
	file, err := os.Open("assets/2023_day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 30000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes\n", count)

	dataString := string(data[:count])
	dataList := strings.Split(dataString, "\n")
	total := 0

	for i := 0; i < len(dataList); i++ {
		var pointer *string = &dataList[i]
		replaceStrNumber(*pointer)
		firstNumber, lastNumber := extractFirstAndLastNumber(dataList[i])

		//fmt.Println(firstNumber, lastNumber)

		n, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			log.Fatal(err)
		}

		total += n
	}
	fmt.Println(dataString)
	fmt.Println("TOTAL:", total)
}
