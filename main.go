package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	fmt.Println("welcome to file Writer")

	fmt.Printf("message : ")

	userData := userInput()

	cleanUserInput := userData + "\n"

	fmt.Printf("times : ")

	timesString := userInput()

	times, err := strconv.ParseInt(timesString, 10, 64)

	handleError(err)

	writeFile(cleanUserInput, int(times))

}

func writeFile(data string, times int) {

	file, err := os.Create("./data.txt")
	handleError(err)
	defer file.Close()

	for i := 0; i <= times; i++ {
		_, err = io.WriteString(file, data)
		handleError(err)
	}

	fmt.Println("completed")

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
