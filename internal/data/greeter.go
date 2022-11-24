package data

import (
	"context"

	"github.com/go-leo/leo/log"
	"github.com/xiaoyan648/project-layout/internal/model"
)

type greeterRepo struct {
	data *Data
	log  log.Logger
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) model.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  logger.With(log.F{K: "layer", V: "data"}),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *model.Greeter) (*model.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *model.Greeter) (*model.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*model.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*model.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*model.Greeter, error) {
	return nil, nil
}
