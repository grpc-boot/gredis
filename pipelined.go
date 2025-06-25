package gredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (p *Pool) TxPipelined(ctx context.Context, fn func(pipe redis.Pipeliner) error) ([]redis.Cmder, error) {
	cmdList, err := p.client.TxPipelined(ctx, fn)
	if !IsNil(err) {
		WriteLog(err, "MULTI EXEC", p.opt)
	}

	return cmdList, err
}

func (p *Pool) Pipelined(ctx context.Context, fn func(pipe redis.Pipeliner) error) ([]redis.Cmder, error) {
	cmdList, err := p.client.Pipelined(ctx, fn)
	if !IsNil(err) {
		WriteLog(err, "MULTI PIPELINE", p.opt)
	}

	return cmdList, err
}
