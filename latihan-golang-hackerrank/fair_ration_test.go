package latihan_golang_hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

/*
 * Complete the 'fairRations' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY B as parameter.
 */

func fairRations(B []int32) string {
	// Write your code here
	counter := 0
	// alreadyGive := 0
	var temp int32 = 0
	for _, value := range B {
		if value%2 != 0 {
			if temp == 1 {
				value += temp
				temp = 0
				counter++
			}
			if value%2 != 0 {
				counter++
				temp++
			}
		} else {
			if temp == 1 {
				value += temp
				temp = 0 // reset
				counter++
			}
			if value%2 != 0 {
				counter++
				temp++
			}
		}
	}
	if temp != 0 {
		return "NO"
	}
	return strconv.Itoa(counter)

}

func TestFairRation(t *testing.T) {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
