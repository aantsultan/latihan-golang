package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "aant,sultan,rahmanya\n"+
		"seele,veli,ona\n"+
		"ani,budi,dodo"

	reader := csv.NewReader(strings.NewReader(csvString))

	for { // looping tidak terbatas
		record, err := reader.Read()
		if err == io.EOF { // baris paling akhir akan mengehentikan looping
			break
		}
		fmt.Println(record)
	}
}
