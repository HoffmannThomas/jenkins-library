package main

//go:generate go run pkg/generator/step-metadata.go --metadataDir=./resources/metadata/ --targetDir=./cmd/

import (
	"github.com/HoffmannThomas/jenkins-library/cmd"
)

func main() {
	cmd.Execute()
}
