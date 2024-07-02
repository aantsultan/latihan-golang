package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("selamat datang\n")
	writer.Flush()
}
