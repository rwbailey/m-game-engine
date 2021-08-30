package grpc

import (
	"context"
	"net"

	"github.com/rs/zerolog/log"
	pbGameengine "github.com/rwbailey/m-apis/game-engine/v1"
	"github.com/rwbailey/m-game-engine/internal/server/logic"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (*Grpc) GetSize(context.Context, *pbGameengine.GetSizeRequest) (*pbGameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in game engine called")
	return &pbGameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (*Grpc) SetScore(ctx context.Context, r *pbGameengine.SetScoreRequest) (*pbGameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in game engine called")
	ok := logic.SetScore(r.GetScore())
	return &pbGameengine.SetScoreResponse{
		Set: ok,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lst, err := net.Listen("tcp", g.address)
	if err != nil {
		return err
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)
	pbGameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Msg("Starting grpc server for m-game-engine on address: " + g.address)

	err = g.srv.Serve(lst)
	if err != nil {
		return err
	}
	return nil
}
