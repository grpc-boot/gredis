package gredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (p *Pool) GetBit(ctx context.Context, key string, offset int64) (value int64, err error) {
	var (
		cmd = p.client.GetBit(ctx, key, offset)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) (num int64, err error) {
	var (
		cmd = p.client.BitCount(ctx, key, bitCount)
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

func (p *Pool) BitOpAnd(ctx context.Context, destKey string, keys ...string) (bytesLen int64, err error) {
	var (
		cmd = p.client.BitOpAnd(ctx, destKey, keys...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) BitOpOr(ctx context.Context, destKey string, keys ...string) (bytesLen int64, err error) {
	var (
		cmd = p.client.BitOpOr(ctx, destKey, keys...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) BitOpNot(ctx context.Context, destKey string, key string) (bytesLen int64, err error) {
	var (
		cmd = p.client.BitOpNot(ctx, destKey, key)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) BitOpXor(ctx context.Context, destKey string, keys ...string) (bytesLen int64, err error) {
	var (
		cmd = p.client.BitOpXor(ctx, destKey, keys...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

// BitField is an API after Redis version 3.2
// type:	位宽格式：u（无符号）/i（有符号）+ 位数（1-64）	u8（8位无符号整数）
// offset:	位段起始位置（比特偏移量），支持#前缀指定字段序（如#0第1个字段）	0 或 #1
// 原子增加第2个字段（4位无符号计数器）
// BITFIELD api_count OVERFLOW SAT INCRBY u4 #1 1
// OVERFLOW: WRAP	回绕（默认）：溢出时从最小值重新计数	循环计数器（如ID生成）
// OVERFLOW: SAT	饱和：超出上限取最大值，低于下限取最小值	限制数值范围（如温度值）
// OVERFLOW: FAIL	失败：溢出时返回 nil，不执行操作	严格数值控制（如余额）
func (p *Pool) BitField(ctx context.Context, key string, values ...any) (res []int64, err error) {
	var (
		cmd = p.client.BitField(ctx, key, values...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
