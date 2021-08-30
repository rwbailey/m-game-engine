package main

import (
	"flag"

	"github.com/rs/zerolog/log"
	grpcSetup "github.com/rwbailey/m-game-engine/internal/server/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":60051", "address to connec to m-game-engine service")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start m-game-engine service")
	}
}
