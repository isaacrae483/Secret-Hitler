package main

import (
	"SecretHitlerBackend/environment"
	"github.com/pressly/goose/v3"
	"log"
	"os"

	_ "SecretHitlerBackend/migrations"
)

func main() {
	args := os.Args[1:]
	command := args[0]

	appConfig := environment.Connect(true)

	var arguments []string
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, appConfig.DB, "../migrations", arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err.Error())
	}
}
