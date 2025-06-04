package gredis

import "context"

func (p *Pool) SAdd(ctx context.Context, key string, members ...any) (newNum int64, err error) {
	var (
		cmd = p.client.SAdd(ctx, key, members...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SRem(ctx context.Context, key string, members ...any) (delNum int64, err error) {
	var (
		cmd = p.client.SRem(ctx, key, members...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SPop(ctx context.Context, key string) (member string, err error) {
	var (
		cmd = p.client.SPop(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SPopN(ctx context.Context, key string, count int64) (members []string, err error) {
	var (
		cmd = p.client.SPopN(ctx, key, count)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SIsMember(ctx context.Context, key string, member any) (exists bool, err error) {
	var (
		cmd = p.client.SIsMember(ctx, key, member)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SMembers(ctx context.Context, key string) (list []string, err error) {
	var (
		cmd = p.client.SMembers(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SCard(ctx context.Context, key string) (length int64, err error) {
	var (
		cmd = p.client.SCard(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, newCursor uint64, err error) {
	var (
		cmd = p.client.SScan(ctx, key, cursor, match, count)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
