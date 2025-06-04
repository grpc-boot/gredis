package gredis

import "context"

func (p *Pool) HIncrBy(ctx context.Context, key, field string, incr int64) (value int64, err error) {
	var (
		cmd = p.client.HIncrBy(ctx, key, field, incr)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HIncrByFloat(ctx context.Context, key, field string, incr float64) (value float64, err error) {
	var (
		cmd = p.client.HIncrByFloat(ctx, key, field, incr)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HSet(ctx context.Context, key, field string, value any) (isNew int64, err error) {
	var (
		cmd = p.client.HSet(ctx, key, field, value)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HSetNX(ctx context.Context, key, field string, value any) (ok bool, err error) {
	var (
		cmd = p.client.HSetNX(ctx, key, field, value)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HGet(ctx context.Context, key, field string) (value string, err error) {
	var (
		cmd = p.client.HGet(ctx, key, field)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HGetBytes(ctx context.Context, key, field string) (value []byte, err error) {
	var (
		cmd = p.client.HGet(ctx, key, field)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Bytes()
}

func (p *Pool) HMSet(ctx context.Context, key string, values ...any) (ok bool, err error) {
	var (
		cmd = p.client.HMSet(ctx, key, values...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HSetMap(ctx context.Context, key string, fv map[string]any) (value bool, err error) {
	var (
		args  = make([]any, 2*len(fv))
		index int
	)

	for field, item := range fv {
		args[index] = field
		index++
		args[index] = item
		index++
	}

	return p.HMSet(ctx, key, args...)
}

func (p *Pool) HExists(ctx context.Context, key string, field string) (exists bool, err error) {
	var (
		cmd = p.client.HExists(ctx, key, field)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HDel(ctx context.Context, key string, fields ...string) (delNum int64, err error) {
	var (
		cmd = p.client.HDel(ctx, key, fields...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HMGetMap(ctx context.Context, key string, fields ...string) (valMap map[string]string, err error) {
	var list []string
	list, err = p.HMGet(ctx, key, fields...)
	if err != nil {
		return
	}

	valMap = make(map[string]string, len(fields))
	for index, _ := range fields {
		valMap[fields[index]] = list[index]
	}

	return
}

func (p *Pool) HMGet(ctx context.Context, key string, fields ...string) (values []string, err error) {
	var (
		cmd     = p.client.HMGet(ctx, key, fields...)
		valList = cmd.Val()
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	if len(valList) > 0 {
		values = make([]string, len(valList))
		for index, value := range valList {
			values[index], _ = value.(string)
		}
	}

	err = cmd.Err()
	return
}

func (p *Pool) HGetAll(ctx context.Context, key string) (value map[string]string, err error) {
	var (
		cmd = p.client.HGetAll(ctx, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, newCursor uint64, err error) {
	var (
		cmd = p.client.HScan(ctx, key, cursor, match, count)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
