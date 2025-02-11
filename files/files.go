package files

import (
	"fmt"
	"os"
)

func Files() {

	file, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	tamanho, err := file.Write([]byte("Hello, world!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("tamanho: %d", tamanho)

}
