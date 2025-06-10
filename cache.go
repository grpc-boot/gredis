package gredis

import (
	"context"
	"time"

	proto2 "github.com/grpc-boot/gredis/proto"

	"google.golang.org/protobuf/proto"
)

var (
	DefaultCacheTimeout = time.Hour * 24
)

type CacheHandler func() (value []byte, err error)

func (p *Pool) ComCache(ctx context.Context, key string, timeoutSeconds int64, handler CacheHandler) (*proto2.CacheItem, error) {
	var (
		cur  = time.Now().Unix()
		item = &proto2.CacheItem{}
	)

	data, err := p.GetBytes(ctx, key)
	if err != nil {
		if IsNil(err) {
			if p.fetchAndSaveData(ctx, key, item, handler) == nil {
				return item, nil
			}
		}

		return nil, err
	}

	err = proto.Unmarshal(data, item)
	if err != nil {
		token, _ := p.Acquire(ctx, key, timeoutSeconds)
		if token > 0 {
			if p.fetchAndSaveData(ctx, key, item, handler) == nil {
				return item, nil
			}
		}

		return nil, err
	}

	if !item.Expired(cur, timeoutSeconds) {
		return item, nil
	}

	token, _ := p.Acquire(ctx, key, timeoutSeconds)
	if token == 0 {
		return item, nil
	}

	_ = p.fetchAndSaveData(ctx, key, item, handler)
	return item, nil
}

func (p *Pool) fetchAndSaveData(ctx context.Context, key string, ci *proto2.CacheItem, handler CacheHandler) error {
	value, err := handler()
	if err != nil {
		return err
	}

	ci.SaveData(value)
	data, err := ci.Marshal()
	if err != nil {
		WriteLog(err, "Marshal CacheItem", p.opt)
		return nil
	}

	_, err = p.SetEx(ctx, key, data, DefaultCacheTimeout)
	return nil
}
