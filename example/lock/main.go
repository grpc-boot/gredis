package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/grpc-boot/gredis"
)

const (
	redKey = `redis`
)

func init() {
	gredis.SetErrorLog(func(err error, cmd string, opt *gredis.Option) {
		fmt.Printf("exec [%s] failed at[%s]\n", cmd, opt.Addr())
	})

	opt, err := gredis.JsonOption([]byte(`{}`))
	if err != nil {
		fmt.Printf("init redis option failed, err:%v\n", err)
		os.Exit(1)
	}

	gredis.Put(redKey, opt)
}

func main() {
	var (
		red = gredis.Get(redKey)
		key = `lockT`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token, err := red.Acquire(ctx, key, 8)
	if err != nil {
		fmt.Printf("acquire redis failed, err:%v\n", err)
		return
	}

	fmt.Printf("got lock token:%v\n", token)

	// todo somethings
	time.Sleep(time.Second)

	// release lock
	_, _ = red.Release(context.Background(), key, token)
}
