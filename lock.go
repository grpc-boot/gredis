package gredis

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	lockFormat = `gored_L:%s`
)

var (
	delTokenScript = redis.NewScript(`if redis.call('get', KEYS[1]) == ARGV[1]
	then
		return redis.call('del', KEYS[1])
	end
	return 0`)
)

func (p *Pool) Acquire(ctx context.Context, key string, lockSeconds int64) (token int64, err error) {
	var (
		t  = rand.Int63()
		ok bool
	)

	ok, err = p.SetNx(ctx, fmt.Sprintf(lockFormat, key), t, time.Duration(lockSeconds)*time.Second)
	if ok {
		token = t
	}
	return
}

func (p *Pool) HasLock(ctx context.Context, key string) (hasLock bool, err error) {
	num, err := p.Exists(ctx, fmt.Sprintf(lockFormat, key))
	if err != nil {
		return
	}

	return num == 1, nil
}

func (p *Pool) Release(ctx context.Context, key string, token int64) (delNum int64, err error) {
	var res any
	res, err = p.RunScript(ctx, delTokenScript, []string{fmt.Sprintf(lockFormat, key)}, token)
	delNum, _ = res.(int64)
	return
}
