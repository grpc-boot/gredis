package gredis

import "context"

func (p *Pool) Info(ctx context.Context, sections ...string) (info string, err error) {
	var (
		cmd = p.client.Info(ctx, sections...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) InfoMap(ctx context.Context, sections ...string) (info map[string]map[string]string, err error) {
	var (
		cmd = p.client.InfoMap(ctx, sections...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ConfigGet(ctx context.Context, parameter string) (res Param, err error) {
	var (
		cmd = p.client.ConfigGet(ctx, parameter)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) ConfigSet(ctx context.Context, parameter, value string) (ok bool, err error) {
	var (
		cmd = p.client.ConfigSet(ctx, parameter, value)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	ok = cmd.Val() == OK
	err = cmd.Err()
	return
}
