package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	pbGameengine "github.com/rwbailey/m-apis/game-engine/v1"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to address: " + *addressPtr)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := pbGameengine.NewGameEngineClient(conn)
	if c == nil {
		log.Info().Msg("client nil")
	}

	resp, err := c.GetSize(timeoutCtx, &pbGameengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("error calling GetHighScore")
	}

	log.Info().Msg(fmt.Sprintf("%f", resp.GetSize()))
}
