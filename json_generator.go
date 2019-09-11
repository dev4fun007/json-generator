package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	f, err := os.OpenFile(buildFilename(), os.O_CREATE|os.O_WRONLY, os.ModeExclusive)
	checkError(err)

	w := bufio.NewWriter(f)

	//Start json array - [ and newline
	_, err = w.WriteString("[")
	checkError(err)

	for i:= 1; i <= MaxLen; i++ {
		if i < MaxLen {
			_, err = w.WriteString(JsonSeed + ",")
		} else if i == MaxLen {
			_, err = w.WriteString(JsonSeed)
		}
		checkError(err)
	}

	//Close json array - ]
	_, err = w.WriteString("]")
	checkError(err)

	//Flush the writer
	err = w.Flush()
	checkError(err)

	err = f.Close()
	checkError(err)

}


func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func buildFilename() string {
	fileName := "generated_" + time.Now().Format("2006_01_02_15_04_05") + ".json"
	fmt.Printf("Filename: %v\n", fileName)
	return fileName
}