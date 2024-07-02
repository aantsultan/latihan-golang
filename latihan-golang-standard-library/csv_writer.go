package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"aant","sultan","rahmanya"})
	_ = writer.Write([]string{"seele","veli","ona"})
	_ = writer.Write([]string{"ani","budi","dodo"})

	writer.Flush()
}
