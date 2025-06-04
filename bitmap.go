package gredis

import "context"

func (p *Pool) GetBit(ctx context.Context, key string, offset int64) (value int64, err error) {
	var (
		cmd = p.client.GetBit(ctx, key, offset)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) SetBit(ctx context.Context, key string, offset int64, value int64) (oldValue int64, err error) {
	var (
		cmd = p.client.SetBit(ctx, key, offset, int(value))
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
