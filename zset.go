package gredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (p *Pool) ZAdd(ctx context.Context, key string, members ...redis.Z) (newNum int64, err error) {
	var (
		cmd = p.client.ZAdd(ctx, key, members...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZCard(ctx context.Context, key string) (length int64, err error) {
	var (
		cmd = p.client.ZCard(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZCount(ctx context.Context, key, min, max string) (count int64, err error) {
	var (
		cmd = p.client.ZCount(ctx, key, min, max)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZScore(ctx context.Context, key, member string) (score float64, err error) {
	var (
		cmd = p.client.ZScore(ctx, key, member)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRange(ctx context.Context, key string, start, stop int64) (list []string, err error) {
	var (
		cmd = p.client.ZRange(ctx, key, start, stop)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRangeWithScores(ctx context.Context, key string, start, stop int64) (list []redis.Z, err error) {
	var (
		cmd = p.client.ZRangeWithScores(ctx, key, start, stop)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRevRange(ctx context.Context, key string, start, stop int64) (list []string, err error) {
	var (
		cmd = p.client.ZRevRange(ctx, key, start, stop)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) (list []redis.Z, err error) {
	var (
		cmd = p.client.ZRevRangeWithScores(ctx, key, start, stop)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) (list []string, err error) {
	var (
		cmd = p.client.ZRangeByScore(ctx, key, opt)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) (list []redis.Z, err error) {
	var (
		cmd = p.client.ZRangeByScoreWithScores(ctx, key, opt)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRank(ctx context.Context, key, member string) (rank int64, err error) {
	var (
		cmd = p.client.ZRank(ctx, key, member)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZRevRank(ctx context.Context, key, member string) (rank int64, err error) {
	var (
		cmd = p.client.ZRevRank(ctx, key, member)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZIncrBy(ctx context.Context, key string, increment float64, member string) (score float64, err error) {
	var (
		cmd = p.client.ZIncrBy(ctx, key, increment, member)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, newCursor uint64, err error) {
	var (
		cmd = p.client.ZScan(ctx, key, cursor, match, count)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
