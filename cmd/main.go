package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yuganksinghal/mEmIFy"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("With a seed of 0: ", mEmIFy.SpongebobCaseSeed(input, 0))
	fmt.Println("With convenience function: ", mEmIFy.SpongebobCase(input))
}
