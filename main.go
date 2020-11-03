package main

import (
	"fmt"
	esg "github.com/p-rs/eschargo/pkg/escape"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) > 0 {
		joined := strings.Join(os.Args[1:], " ")
		out, err := esg.EscString(joined)
		if err != nil {
			_ = fmt.Errorf("ERROR:\t%s", err)
			os.Exit(1)
		}

		fmt.Println(out)
	} else {
		out, err := esg.EscPipe(os.Stdin)
		if err != nil {
			_ = fmt.Errorf("ERROR:\t%s", err)
			os.Exit(1)
		}

		fmt.Print(out)
	}
}
