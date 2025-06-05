package gredis

import (
	"context"
	"fmt"
	"testing"
	"time"
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

func TestPool_GetBit(t *testing.T) {
	var (
		newValue    = int64(1)
		offset      = int64(8)
		key         = `bitmapT`
		red         = Get(redKey)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	)

	defer cancel()

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
}
