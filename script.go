package gredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (p *Pool) RunScript(ctx context.Context, script *redis.Script, keys []string, args ...any) (res string, err error) {
	var (
		cmd = script.Run(ctx, p.client, keys, args)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.String(), cmd.Err()
}
