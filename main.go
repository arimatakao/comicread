package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/arimatakao/comicread/internal/cli"
	"github.com/arimatakao/comicread/internal/i18n"
)

func main() {
	err := cli.Run(os.Args[1:])
	if errors.Is(err, flag.ErrHelp) {
		return
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "comicread:", err)
		if errors.Is(err, cli.ErrUsage) {
			fmt.Fprintln(os.Stderr, i18n.T("cli.help_hint"))
		}
		os.Exit(1)
	}
}
