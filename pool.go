package gredis

import (
	"github.com/redis/go-redis/v9"
)

type Pool struct {
	client *redis.Client
	opt    *Option
}

func NewPool(opt *Option) *Pool {
	return &Pool{
		client: opt.NewClient(),
		opt:    opt,
	}
}

func (p *Pool) Option() *Option {
	return p.opt
}
