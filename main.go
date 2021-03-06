package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func checkStackExists(stackName string) error {
	directoryContents, err := os.ReadDir("./stacks")
	if err != nil {
		return fmt.Errorf("unable to read current directory %w", err)
	}

	for _, f := range directoryContents {
		if f.IsDir() && f.Name() == stackName {
			return fmt.Errorf("stack %s already exists", stackName)
		}
	}
	return nil
}

func createStackDirectory(stackName string) error {
	err := os.Mkdir(stackName, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s, %w", stackName, err)
	}
	return nil
}

func populateStackDirectory(stackNameDirectory string) error {
	templateContents, err := os.ReadDir("./templates")
	if err != nil {
		return fmt.Errorf("unable to read templates directory, %w", err)
	}

	for _, f := range templateContents {
		sourceFile := fmt.Sprintf("%s/%s", "./templates", f.Name())
		err = copyFile(sourceFile, stackNameDirectory)
		if err != nil {
			return fmt.Errorf("error while copying source to destination, %w", err)
		}
	}
	fmt.Printf("Templated stack %s successfully.\n", stackNameDirectory)
	return nil
}

func copyFile(sourceFile, destinationDirectory string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("unable to open source file %s, %w", sourceFile, err)
	}

	defer source.Close()
	destinationFile := fmt.Sprintf("%s/%s", destinationDirectory, path.Base(sourceFile))
	destination, err := os.Create(destinationFile)
	if err != nil {
		return fmt.Errorf("unable to create destination file %s, %w", destinationFile, err)
	}

	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("unable to copy source file to destination, %w", err)
	}
	return nil
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
			stackName := "stacks/" + c.String("s")

			err := checkStackExists(stackName)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			err = createStackDirectory(stackName)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			err = populateStackDirectory(stackName)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
