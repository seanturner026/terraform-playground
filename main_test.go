package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateStackDirectory(t *testing.T) {
	t.Run("Successfully created mock stack", func(t *testing.T) {
		const mockDirectory string = "mock-create-stack"

		err := createStackDirectory(mockDirectory)
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(mockDirectory)
		directoryContents, err := os.ReadDir("./")
		if err != nil {
			t.Fatalf("Unable to make read current directory, %s", err)
		}

		var mockDirectoryExists bool
		for _, d := range directoryContents {
			if d.Name() == mockDirectory && d.IsDir() {
				mockDirectoryExists = true
				break
			}
		}

		if !mockDirectoryExists {
			t.Fatalf("Mock directory %s never created", mockDirectory)
		}
	})
}

func TestPopulateDirectoryStack(t *testing.T) {
	t.Run("Successfully populated mock stack", func(t *testing.T) {
		const mockDirectory string = "mock-populate-stack"
		expectedTerraformFiles := []string{"main.tf", "outputs.tf", "provider.tf", "terraform.tfvars", "variables.tf"}

		err := os.Mkdir(mockDirectory, 0755)
		if err != nil {
			t.Fatalf("Unable to make mock directory, %s", err)
		}

		err = populateStackDirectory(mockDirectory)
		if err != nil {
			t.Fatal(err)
		}

		defer os.RemoveAll(mockDirectory)
		mockDirectoryContents, err := os.ReadDir(mockDirectory)
		if err != nil {
			t.Fatal(err)
		}

		if len(expectedTerraformFiles) != len(mockDirectoryContents) {
			t.Fatal("Mock directory has a different number of files than expected")
		}

		for _, f := range mockDirectoryContents {
			switch f.Name() {
			case "main.tf":
			case "outputs.tf":
			case "provider.tf":
			case "terraform.tfvars":
			case "variables.tf":
			default:
				t.Fatal("Mock directory has different terraform files than expected")
			}
		}
	})
}

func TestCopyFile(t *testing.T) {
	t.Run("Successfully copied file mock file", func(t *testing.T) {
		const mockDirectory string = "mock-copy-file"

		err := os.Mkdir(mockDirectory, 0755)
		if err != nil {
			t.Fatalf("Unable to make mock directory, %s", err)
		}

		err = copyFile("README.md", mockDirectory)
		if err != nil {
			t.Fatal(err)
		}

		defer os.RemoveAll(mockDirectory)
		mockFile, err := os.Open(fmt.Sprintf("%s/%s", mockDirectory, "README.md"))
		if err != nil {
			t.Fatalf("Unable to open mocked README.md, %s", err)
		}

		readmeFile, err := os.Open("README.md")
		if err != nil {
			t.Fatalf("Unable to open README.md, %s", err)
		}

		mockFileInfo, err := mockFile.Stat()
		if err != nil {
			t.Fatalf("Unable to obtain mocked README.md file info, %s", err)
		}

		readmeFileInfo, err := readmeFile.Stat()
		if err != nil {
			t.Fatalf("Unable to obtain README.md file info, %s", err)
		}

		if mockFileInfo.Name() != readmeFileInfo.Name() {
			t.Fatal("README.md and mocked README.md are have different names")
		} else if mockFileInfo.Size() != readmeFileInfo.Size() {
			t.Fatal("README.md and mocked README.md are have different sizes")
		} else if mockFileInfo.Mode() != readmeFileInfo.Mode() {
			t.Fatal("README.md and mocked README.md are have different modes")
		}
	})
}
