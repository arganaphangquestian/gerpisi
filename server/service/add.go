package service

import (
	"context"
	"log"

	"github.com/arganaphangquestian/gerpisi/server/data"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, req *data.AddRequest) (*data.AddResponse, error) {
	log.Printf("Hey hey hey, you add 2 number right?, here's the result %d\n", req.A+req.B)
	return &data.AddResponse{
		Res: req.A + req.B,
	}, nil
}
