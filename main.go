package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func createStackDirectory(stackName string) {
	directoryContents, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range directoryContents {
		if f.IsDir() && f.Name() == stackName {
			log.Fatalf("Stack %s already exists.", stackName)
		}
	}

	err = os.Mkdir(stackName, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func populateStack(stackName string) {
	templateContents, err := os.ReadDir("./templates")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range templateContents {
		copyFile(stackName, f.Name())
	}
	fmt.Printf("Templated stack %s successfully.\n", stackName)
}

func copyFile(stackName, fileName string) {
	source, err := os.Open(fmt.Sprintf("%s/%s", "./templates", fileName))
	if err != nil {
		log.Fatal(err)
	}

	defer source.Close()
	destination, err := os.Create(fmt.Sprintf("%s/%s", stackName, fileName))
	if err != nil {
		log.Fatal(err)
	}

	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.App{
		Name: "terraform-playground",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "s",
				Aliases:  []string{"stack-name"},
				Usage:    "Name of the stack to create.",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			stackName := c.String("s")
			createStackDirectory(stackName)
			populateStack(stackName)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
