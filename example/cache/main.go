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
