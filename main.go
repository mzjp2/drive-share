package main

import (
	"flag"
	"fmt"

	"github.com/mzjp2/drive-share/drive"
)

func main() {
	fileName := flag.String("name", "default", " The file name of the file you want to share")
	flag.Parse()
	fmt.Printf("%s\n", *fileName)

	drive.ListFiles(fmt.Sprintf("name = '%s'", *fileName))
}
