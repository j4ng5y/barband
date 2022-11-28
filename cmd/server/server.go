package main

import (
	servercli "github.com/j4ng5y/barband/pkg/cmd/server"
	"github.com/rs/zerolog/log"
)

//go:generate make generate

func main() {
	if err := servercli.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
