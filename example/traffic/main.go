package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
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
	)

	limitByIncr(red)
	time.Sleep(time.Second * 3)
	limitByDay(red)
	time.Sleep(time.Second * 3)
	limitByMinute(red)
	time.Sleep(time.Second * 3)
	limitBySecond(red)
}

func limitByIncr(red *gredis.Pool) {
	fmt.Println("limit by incr start")

	var (
		key     = fmt.Sprintf(`trafficIncr:%s`, time.Now().Format("2006-01-02"))
		num     = 128
		okCount = 0
		start   = time.Now()
	)

	for i := 0; i < num; i++ {
		var (
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			ok, err     = red.AcquireByIncr(ctx, key, 1, 100, 3600)
		)

		cancel()
		if err != nil {
			fmt.Printf("acquire redis failed, err:%v\n", err)
			return
		}

		if !ok {
			fmt.Println("limited by incr")
		} else {
			okCount++
			fmt.Printf("acquire redis success: %d cost: %s\n", okCount, time.Since(start))
		}
		time.Sleep(time.Millisecond * 5)
	}

	fmt.Println("limit by incr end")
}

// 每秒100个请求，突发可以突破到128个
func limitBySecond(red *gredis.Pool) {
	fmt.Println("limit by second start")

	var (
		key     = `trafficSecondsT`
		num     = 1000
		okCount = 0
		start   = time.Now()
	)

	for i := 0; i < num; i++ {
		var (
			current     = time.Now().Unix()
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			ok, err     = red.AcquireByLeakyBucket(ctx, key, current, 128, 100, 1, 3600)
		)

		cancel()
		if err != nil {
			fmt.Printf("acquire redis failed, err:%v\n", err)
			return
		}

		if !ok {
			fmt.Println("limited by seconds")
		} else {
			okCount++
			fmt.Printf("acquire redis success: %d cost: %s\n", okCount, time.Since(start))
		}
		time.Sleep(time.Millisecond * 5)
	}

	fmt.Println("limit by second end")
}

// 每分钟60个请求
func limitByMinute(red *gredis.Pool) {
	fmt.Println("limit by minute start")
	var (
		key     = `trafficMinuteT`
		num     = 1000
		okCount = 0
		start   = time.Now()
	)

	for i := 0; i < num; i++ {
		var (
			current, _  = strconv.ParseInt(time.Now().Format("200601021504"), 10, 64)
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			ok, err     = red.AcquireByLeakyBucket(ctx, key, current, 60, 60, 1, 3600)
		)

		cancel()
		if err != nil {
			fmt.Printf("acquire redis failed, err:%v\n", err)
			return
		}

		if !ok {
			fmt.Println("limited by minute")
		} else {
			okCount++
			fmt.Printf("acquire redis success: %d cost: %s\n", okCount, time.Since(start))
		}
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("limit by minute end")
}

func limitByDay(red *gredis.Pool) {
	fmt.Println("limit by day start")
	var (
		key     = `trafficDayT`
		num     = 100
		okCount = 0
		start   = time.Now()
	)

	for i := 0; i < num; i++ {
		var (
			current, _  = strconv.ParseInt(time.Now().Format("20060102"), 10, 64)
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			ok, err     = red.AcquireByLeakyBucket(ctx, key, current, 60, 60, 1, 3600*24*2)
		)

		cancel()
		if err != nil {
			fmt.Printf("acquire redis failed, err:%v\n", err)
			return
		}

		if !ok {
			fmt.Println("limited by day")
		} else {
			okCount++
			fmt.Printf("acquire redis success: %d cost: %s\n", okCount, time.Since(start))
		}
		time.Sleep(time.Millisecond * 50)
	}

	fmt.Println("limit by day end")
}
