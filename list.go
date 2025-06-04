package gredis

import "context"

func (p *Pool) LLen(ctx context.Context, key string) (length int64, err error) {
	var (
		cmd = p.client.LLen(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) LRange(ctx context.Context, key string, start, stop int64) (items []string, err error) {
	var (
		cmd = p.client.LRange(ctx, key, start, stop)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) LPush(ctx context.Context, key string, values ...any) (length int64, err error) {
	var (
		cmd = p.client.LPush(ctx, key, values...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) LPop(ctx context.Context, key string) (value string, err error) {
	var (
		cmd = p.client.LPop(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) RPush(ctx context.Context, key string, values ...any) (length int64, err error) {
	var (
		cmd = p.client.RPush(ctx, key, values...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) RPop(ctx context.Context, key string) (value string, err error) {
	var (
		cmd = p.client.RPop(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) RPopLPush(ctx context.Context, srcKey, dstKey string) (value string, err error) {
	var (
		cmd = p.client.RPopLPush(ctx, srcKey, dstKey)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) LRem(ctx context.Context, key string, count int64, val any) (items int64, err error) {
	var (
		cmd = p.client.LRem(ctx, key, count, val)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
