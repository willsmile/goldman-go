package main

import (
	"log"
	"os"

	"github.com/willsmile/goldman-go/internal/cli"
)

func main() {
	if err := cli.New().Run(os.Args); err != nil {
		log.Fatal("[Error] ", err)
	}
}
