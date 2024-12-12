package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	line, _ := reader.ReadString('\n')
	temps := strings.Fields(line)

	temp := 0
	for i := 0; i < n; i++ {
		fmt.Sscanf(temps[i], "%d", &temp)

		if temp > 15 {
			fmt.Fprintln(writer, "cooler")
		} else {
			fmt.Fprintln(writer, "heater")
		}
	}
}
