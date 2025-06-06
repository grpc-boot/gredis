package gredis

import (
	"context"
	"time"
)

func (p *Pool) Scan(ctx context.Context, cursor uint64, match string, count int64) (keys []string, newCursor uint64, err error) {
	var (
		cmd = p.client.Scan(ctx, cursor, match, count)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) Type(ctx context.Context, key string) (t string, err error) {
	var (
		cmd = p.client.Type(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) Exists(ctx context.Context, keys ...string) (existsNum int64, err error) {
	var (
		cmd = p.client.Exists(ctx, keys...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) Ttl(ctx context.Context, key string) (duration time.Duration, err error) {
	var (
		cmd = p.client.TTL(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) Del(ctx context.Context, keys ...string) (delNum int64, err error) {
	var (
		cmd = p.client.Del(ctx, keys...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) Expire(ctx context.Context, key string, timeout time.Duration) (ok bool, err error) {
	var (
		cmd = p.client.Expire(ctx, key, timeout)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ExpireAt(ctx context.Context, key string, tm time.Time) (exists bool, err error) {
	var (
		cmd = p.client.ExpireAt(ctx, key, tm)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
