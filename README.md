# gredis

- TOC
    - [SetLog](#SetLog)
    - [Traffic](#Traffic)
    - [Cache](#Cache)
    - [Geo](#Geo)

### SetLog

```go
package main

import (
	"fmt"

	"github.com/grpc-boot/gredis"
)

func init() {
	gredis.SetErrorLog(func(err error, cmd string, opt *gredis.Option) {
		fmt.Printf("exec [%s] failed at[%s]\n", cmd, opt.Addr())
	})
}
```

### Lock

```go
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
```

### Traffic

```go
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
```

### Cache

```go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/grpc-boot/gredis"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
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
		key = `cacheT`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ci, err := red.ComCache(ctx, key, 15, func() (value []byte, err error) {
		// todo
		time.Sleep(time.Second * 3)

		// 仅支持基本的数据类型，不支持嵌套切片和map等结构
		data, err := structpb.NewStruct(map[string]any{
			"id":      123,
			"age":     uint8(35),
			"name":    "Masco",
			"male":    true,
			"height":  1.73,
			"current": time.Now().Format("2006-01-02 15:04:05"),
		})

		if err != nil {
			return
		}

		return proto.Marshal(data)
	})

	if err != nil {
		fmt.Printf("get cache from redis failed, err:%v\n", err)
		return
	}

	data, err := ci.MapData()
	if err != nil {
		fmt.Printf("get map data failed, err:%v\n", err)
	}

	for k, v := range data {
		fmt.Printf("key:%s, type: %T value:%v\n", k, v, v)
	}
}
```

### Geo

```go
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/grpc-boot/gredis"

	"github.com/redis/go-redis/v9"
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
		key = `geoT`
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newCount, err := red.GeoAdd(ctx,
		key,
		&redis.GeoLocation{
			Name:      "王府井金街店",
			Longitude: 116.410256,
			Latitude:  39.909594,
		},
		&redis.GeoLocation{
			Name:      "西单明珠店",
			Longitude: 116.376018,
			Latitude:  39.909956,
		},
		&redis.GeoLocation{
			Name:      "密云鼓楼大街店",
			Longitude: 116.846254,
			Latitude:  40.375323,
		},
	)

	if err != nil {
		fmt.Printf("redis add failed, err:%v\n", err)
	}

	fmt.Printf("redis add success, count:%d\n", newCount)

	locs, err := red.GeoRadius(ctx, key, 116.376018, 39.909956, &redis.GeoRadiusQuery{
		Radius:      1000, // 搜索半径
		Unit:        "km", // 单位
		WithCoord:   true, // 返回坐标
		WithDist:    true, // 返回距离
		WithGeoHash: true, // 返回geo hash
		Count:       100,
		Sort:        "ASC",
	})

	if err != nil {
		fmt.Printf("redis query failed, err:%v\n", err)
	}
	fmt.Printf("geo radius: %+v\n", locs)

	dist, err := red.GeoDist(ctx, key, "西单明珠店", "密云鼓楼大街店", "km")
	if err != nil {
		fmt.Printf("redis query failed, err:%v\n", err)
	}
	fmt.Printf("geo dist: %+v\n", dist)

	pos, err := red.GeoPos(ctx, key, "西单明珠店", "密云鼓楼大街店")
	if err != nil {
		fmt.Printf("redis query failed, err:%v\n", err)
	}
	fmt.Printf("geo pos[0] lat:%v long:%v\n", pos[0].Latitude, pos[0].Longitude)
	fmt.Printf("geo pos[1] lat:%v long:%v\n", pos[1].Latitude, pos[1].Longitude)
}
```