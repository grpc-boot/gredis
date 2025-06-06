package gredis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	redKey = `redis`
)

func init() {
	SetErrorLog(func(err error, cmd string, opt *Option) {
		fmt.Printf("exec [%s] failed at[%s]\n", cmd, opt.Addr())
	})

	Put(redKey, DefaultOption())
}

func TestPool_Info(t *testing.T) {
	infoStr, err := Get(redKey).Info(context.Background())
	if err != nil {
		t.Fatalf("get info failed: %v", err)
	}
	t.Logf("info: %s\n", infoStr)

	// top cmd
	info, err := Get(redKey).InfoMap(context.Background(), "COMMANDSTATS")
	if err != nil {
		t.Fatalf("get info failed: %v", err)
	}
	t.Logf("info: %+v\n", info)
}

func TestPool_BitOp(t *testing.T) {
	var (
		LabelSport  = `sportL`
		LabelMusic  = `musicL`
		LabelDance  = `danceL`
		LabelResult = `resultL`
		red         = Get(redKey)
		userA       = int64(100)
		userB       = int64(110)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	)

	defer func() {
		_, _ = red.Del(ctx, LabelSport, LabelMusic, LabelDance, LabelResult)
		cancel()
	}()

	bytesLen, err := red.BitOpAnd(ctx, LabelResult, LabelSport, LabelMusic, LabelDance)
	if err != nil {
		t.Fatalf("BitOpAnd failed: %v", err)
	}

	t.Logf("bytesLen: %d", bytesLen)

	_, err = red.SetBit(ctx, LabelSport, userA, 1)
	if err != nil {
		t.Fatalf("SetBit failed: %v", err)
	}

	_, _ = red.SetBit(ctx, LabelMusic, userA, 1)
	_, _ = red.SetBit(ctx, LabelDance, userA, 1)
	_, _ = red.SetBit(ctx, LabelDance, userB, 1)

	bytesLen, err = red.BitOpAnd(ctx, LabelResult, LabelSport, LabelMusic, LabelDance)
	if err != nil {
		t.Fatalf("BitOpAnd failed: %v", err)
	}

	t.Logf("bytesLen: %d", bytesLen)

	res, err := red.BitCount(ctx, LabelResult, &redis.BitCount{
		Start: 0,
		End:   bytesLen,
	})
	if err != nil {
		t.Fatalf("BitCount failed: %v", err)
	}
	t.Logf("num: %d", res)

	bytesLen, err = red.BitOpOr(ctx, LabelResult, LabelSport, LabelMusic, LabelDance, LabelResult)
	if err != nil {
		t.Fatalf("BitOpOr failed: %v", err)
	}

	t.Logf("bytesLen: %d", bytesLen)
	res, err = red.BitCount(ctx, LabelResult, &redis.BitCount{
		Start: 0,
		End:   bytesLen,
	})
	if err != nil {
		t.Fatalf("BitCount failed: %v", err)
	}
	t.Logf("num: %d", res)
}

func TestPool_GetBit(t *testing.T) {
	var (
		newValue    = int64(0)
		offset      = int64(8)
		key         = `bitmapT`
		red         = Get(redKey)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	)

	defer func() {
		_, _ = red.Del(ctx, key)
		cancel()
	}()

	oldV, err := red.SetBit(ctx, key, offset, newValue)
	if err != nil {
		t.Fatalf("SetBit(%v) failed: %v", key, err)
	}

	t.Logf("SetBit(%v, %v, %v) => %d", key, offset, oldV, newValue)

	newV, err := red.GetBit(ctx, key, offset)
	if err != nil {
		t.Fatalf("GetBit(%v) failed: %v", key, err)
	}
	t.Logf("GetBit(%v, %v) => %d", key, offset, newV)

	num, err := red.BitCount(ctx, key, &redis.BitCount{
		Start: 0,
		End:   100,
	})
	if err != nil {
		t.Fatalf("BitCount(%v) failed: %v", key, err)
	}
	t.Logf("BitCount(%v, %v) => %d", key, offset, num)
}
