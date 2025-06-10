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
