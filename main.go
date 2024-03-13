package main

import (
	"log"
	"os"

	"github.com/ali-ammar-kazmi/Encrypt_Decrypt/handlers"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Encrypt_Decrypt",
		Usage: "Tool to Encrypt/Decrypt a file.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "action",
				Aliases:  []string{"a"},
				Usage:    "Action item to perform.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "path",
				Aliases:  []string{"p"},
				Usage:    "file path to reach destination",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			path := c.String("path")
			switch c.String("action") {
			case "encrypt":
				handlers.EncryptHandler(path)
			case "decrypt":
				handlers.DecryptHandler(path)
			default:
				handlers.EncryptHandler(path)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
