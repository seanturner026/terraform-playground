package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

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

func populateStack(stackNameDirectory string) {
	templateContents, err := os.ReadDir("./templates")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range templateContents {
		sourceFile := fmt.Sprintf("%s/%s", "./templates", f.Name())
		copyFile(sourceFile, stackNameDirectory)
	}
	fmt.Printf("Templated stack %s successfully.\n", stackNameDirectory)
}

func copyFile(sourceFile, destinationDirectory string) {
	source, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	defer source.Close()
	destinationFile := fmt.Sprintf("%s/%s", destinationDirectory, path.Base(sourceFile))
	destination, err := os.Create(destinationFile)
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
