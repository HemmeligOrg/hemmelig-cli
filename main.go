package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/HemmeligOrg/hemmelig-cli/core"
	"github.com/urfave/cli/v2"
)

func input() string {
	data := ""

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return data
}

func createSecret(c *cli.Context) error {
	if c.Int("ttl") < 0 || c.Int("ttl") > 604800 {
		log.Fatal("Sorry, the TTL is expected to be from 0 - 605800 seconds")
	}

	data := ""
	secret := c.Args().Get(0)
	if secret == "" {
		data += input()
	} else {
		data += secret
	}

	url, err := core.CreateSecret(data, c.String("password"), c.String("ttl"), c.String("url"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The secret URL: " + url)

	return nil
}

func main() {
	app := &cli.App{
		Name:      "[he`m:(ə)li]",
		HelpName:  "hemmelig",
		Usage:     "Create a secret URL directly from your CLI.",
		UsageText: "cat your_secret_file.txt | hemmelig --password=cantguessthislol \nOr just pass it as the first argument: hemmelig \"This is my secret\" --password=secret",
		Copyright: fmt.Sprintf("(c) %d Hemmelig.app", core.Year()),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "password",
				Value:   "",
				Usage:   "Set a password to protect the secret",
				Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:    "ttl",
				Value:   "14400",
				Usage:   "Secret expiration time in seconds. 0 - 605800 seconds.",
				Aliases: []string{"t"},
			},
			&cli.StringFlag{
				Name:    "url",
				Value:   "https://hemmelig.app/",
				Usage:   "Override the Hemmelig app URL if you host it yourself",
				Aliases: []string{"u"},
			},
		},
		Action: createSecret,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error starting hemmelig cli: %v", err)
	}
}
