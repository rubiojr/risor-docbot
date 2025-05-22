package main

import (
	"context"
	"fmt"
	"log"
	"os"

	docbot "github.com/rubiojr/risor-docbot"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "risor-docbot",
		Usage: "AI-powered documentation assistant for Risor code",
		Description: "Automatically documents Risor code files using AI. " +
			"Requires OPENAI_API_KEY to be set in the environment.",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Path to the Risor file to document",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose output",
			},
		},
		Action: func(c *cli.Context) error {
			filePath := c.String("file")

			// If file is not provided via flag, try to get it from args
			if filePath == "" && c.Args().Len() > 0 {
				filePath = c.Args().First()
			}

			if filePath == "" {
				return fmt.Errorf("no file specified, use --file flag or provide as argument")
			}

			// Check if file exists
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				return fmt.Errorf("file not found: %s", filePath)
			}

			// Check if OPENAI_API_KEY is set
			if os.Getenv("OPENAI_API_KEY") == "" {
				return fmt.Errorf("OPENAI_API_KEY environment variable is not set")
			}

			bot := docbot.New(filePath)

			if c.Bool("verbose") {
				fmt.Printf("Documenting file: %s\n", filePath)
			}

			_, err := bot.DocumentCode(context.Background())
			if err != nil {
				return fmt.Errorf("documentation failed: %w", err)
			}

			if c.Bool("verbose") {
				fmt.Printf("Documentation completed.")
			} else {
				fmt.Println("Documentation completed successfully!")
			}

			return nil
		},
	}

	// Enable shell completion
	app.EnableBashCompletion = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
