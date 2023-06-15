package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/willsmile/goldman-go/internal/config"
	"github.com/willsmile/goldman-go/internal/date"
)

// version of goldman-go
const version = "0.1.0"

func New() *cli.App {
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
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "",
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"GOLDMAN_GO_PATH"},
			},
			&cli.StringFlag{
				Name:    "start_date",
				Aliases: []string{"s"},
				Value:   time.Now().Format(time.DateOnly),
				Usage:   "Set start date from `DATE`",
			},
			&cli.StringFlag{
				Name:    "end_date",
				Aliases: []string{"e"},
				Value:   "",
				Usage:   "Set end date from `DATE`",
			},
			&cli.IntFlag{
				Name:    "day",
				Aliases: []string{"d"},
				Value:   7,
				Usage:   "Set day interval from start date from `INTEGER`",
			},
			&cli.IntFlag{
				Name:    "week",
				Aliases: []string{"w"},
				Value:   0,
				Usage:   "Set day interval from start date from `INTEGER`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate list of schedule options",
				Action: func(c *cli.Context) error {
					err := generate(c)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	return app
}

func generate(c *cli.Context) error {
	path := c.String("config")
	start := c.String("start_date")
	end := c.String("end_date")
	day := c.Int("day")
	week := c.Int("week")

	drg, err := date.NewDateRange(start, end, day, week)
	if err != nil {
		return err
	}

	cfg, err := config.LoadConfig(path)
	if err != nil {
		return err
	}

	data := cfg.UserDefinedData
	schedule := data.Generate(drg)
	if err != nil {
		return err
	}

	format := cfg.UserDefinedFormat
	for _, s := range schedule {
		fmt.Println(s.Format(format))
	}

	return nil
}
