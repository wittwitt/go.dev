package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	main1()
	// main2()
}

func main1() {
	app := &cli.App{
		Name:    "cli",
		Usage:   "cli for",
		Version: "v0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.toml",
			},
		},
		Before: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			subCmd(),
		},
		Action: func(cctx *cli.Context) error {

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}

func subCmd() *cli.Command {
	return &cli.Command{
		Name: "sub",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "evar",
				Aliases: []string{"c"},
				EnvVars: []string{"GODEV_TEST_EVAR"},
				Value:   "",
			},
		},
		Action: func(cctx *cli.Context) error {

			// env make flag and env ok
			// flag only make flag value

			fmt.Println("flag", cctx.String("evar"))
			fmt.Println("env", os.Getenv("GODEV_TEST_EVAR"))

			go func() {
				for {
					fmt.Println("a")
					time.Sleep(1 * time.Second)
				}
			}()

			select {
			case <-cctx.Context.Done():
				fmt.Println("done")
			}

			return nil
		},
	}
}

func main2() {
	// cli.VersionFlag = &cli.BoolFlag{
	// 	Name:    "print-version",
	// 	Aliases: []string{"V"},
	// 	Usage:   "print only the version",
	// }

	app := &cli.App{
		Name:    "partay",
		Version: "v19.99.0",
	}
	app.Run(os.Args)
}

func main3() {

}
