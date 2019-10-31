package main

//go:generate rm -rf internal/generator/resources.go
//go:generate go-bindata -pkg generator -o internal/generator/resources.go assets/templates/...

import (
	"gitlab.com/mikrowezel/backend/cli/cmd"
)

func main() {
	cmd.Execute()
}
