package model

import (
	"context"

	"github.com/go-leo/leo/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NewWithCode(code.)
)

// Greeter is a Greeter
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  log.Logger
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: logger.With(log.F{K: "layer", V: "model"})}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	// uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello) 需要设置logger context功能
	return uc.repo.Save(ctx, g)
}
