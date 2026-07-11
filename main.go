package main

import (
	"fmt"
	"os"

	"github.com/arimatakao/comicread/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "comicread:", err)
		os.Exit(1)
	}
}
