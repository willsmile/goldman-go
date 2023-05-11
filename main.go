package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// version of goldman-go
const version = "0.1.0"

func main() {
	var (
		schedule []Schedule
		path     string
		start    string
		end      string
		day      int
		week     int
	)

	app := &cli.App{
		Name:    "goldman-go",
		Usage:   "A tool for generating option list of schedule",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "Wei Chen (willsmile)",
				Email: "willsmile.me@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "Load configuration from `FILE`",
				Destination: &path,
				EnvVars:     []string{"GOLDMAN_GO_PATH"},
			},
			&cli.StringFlag{
				Name:        "start_date",
				Aliases:     []string{"s"},
				Value:       time.Now().Format(time.DateOnly),
				Usage:       "Set start date from `DATE`",
				Destination: &start,
			},
			&cli.StringFlag{
				Name:        "end_date",
				Aliases:     []string{"e"},
				Value:       "",
				Usage:       "Set end date from `DATE`",
				Destination: &end,
			},
			&cli.IntFlag{
				Name:        "day",
				Aliases:     []string{"d"},
				Value:       7,
				Usage:       "Set day interval from start date from `INTEGER`",
				Destination: &day,
			},
			&cli.IntFlag{
				Name:        "week",
				Aliases:     []string{"w"},
				Value:       0,
				Usage:       "Set day interval from start date from `INTEGER`",
				Destination: &week,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate list of schedule options",
				Action: func(c *cli.Context) error {
					drg, err := NewDateRange(start, end, day, week)
					if err != nil {
						return err
					}

					cfg, err := LoadConfig(path)
					if err != nil {
						return err
					}

					data := cfg.UserDefinedData
					schedule = data.Generate(drg)
					if err != nil {
						return err
					}

					format := cfg.UserDefinedFormat
					for _, s := range schedule {
						fmt.Println(s.Format(format))
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("[Cli Error] ", err)
	}
}
