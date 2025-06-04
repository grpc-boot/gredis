package gredis

import (
	"context"
	"time"
)

func (p *Pool) MGet(ctx context.Context, keys ...string) (values []string, err error) {
	var (
		cmd         = p.client.MGet(ctx, keys...)
		valList, er = cmd.Result()
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	if len(valList) > 0 {
		values = make([]string, len(valList))
		for index, value := range valList {
			val, _ := value.(string)
			values[index] = val
		}
	}

	return values, er
}

func (p *Pool) MGetBytes(ctx context.Context, keys ...string) (values [][]byte, err error) {
	var (
		cmd         = p.client.MGet(ctx, keys...)
		valList, er = cmd.Result()
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	if len(valList) > 0 {
		values = make([][]byte, len(valList))
		for index, value := range valList {
			val, _ := value.(string)
			values[index] = String2Bytes(val)
		}
	}

	return values, er
}

func (p *Pool) MSet(ctx context.Context, values ...any) (ok bool, err error) {
	var (
		cmd     = p.client.MSet(ctx, values...)
		res, er = cmd.Result()
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}
	return res == OK, er
}

func (p *Pool) Get(ctx context.Context, key string) (value string, err error) {
	var (
		cmd = p.client.Get(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) GetBytes(ctx context.Context, key string) (value []byte, err error) {
	var (
		cmd = p.client.Get(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Bytes()
}

func (p *Pool) Set(ctx context.Context, key string, value any, expiration time.Duration) (ok bool, err error) {
	var (
		cmd     = p.client.Set(ctx, key, value, expiration)
		res, er = cmd.Result()
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return res == OK, er
}

func (p *Pool) SetEx(ctx context.Context, key string, value any, timeout time.Duration) (ok bool, err error) {
	var (
		cmd     = p.client.SetEx(ctx, key, value, timeout)
		res, er = cmd.Result()
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return res == OK, er
}

func (p *Pool) SetNx(ctx context.Context, key string, value any, timeout time.Duration) (ok bool, err error) {
	var (
		cmd = p.client.SetNX(ctx, key, value, timeout)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) IncrBy(ctx context.Context, key string, value int64) (newValue int64, err error) {
	var (
		cmd = p.client.IncrBy(ctx, key, value)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) IncrByFloat(ctx context.Context, key string, value float64) (newValue float64, err error) {
	var (
		cmd = p.client.IncrByFloat(ctx, key, value)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
