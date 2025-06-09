package gredis

import (
	"context"
	"time"

	proto2 "github.com/grpc-boot/gredis/proto"

	"github.com/goccy/go-json"
	"google.golang.org/protobuf/proto"
)

var (
	DefaultCacheTimeout = time.Hour * 24
)

type CacheHandler func() (value []byte, err error)

func (p *Pool) CacheGetWithProto(ctx context.Context, key string, timeoutSeconds int64, handler CacheHandler, out proto.Message) error {
	value, err := p.CacheGet(ctx, key, timeoutSeconds, handler)
	if len(value) == 0 {
		if err != nil {
			return err
		}
	}

	return proto.Unmarshal(value, out)
}

func (p *Pool) CacheGetWithJson(ctx context.Context, key string, timeoutSeconds int64, handler CacheHandler, out any) error {
	value, err := p.CacheGet(ctx, key, timeoutSeconds, handler)
	if len(value) == 0 {
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(value, out)
}

func (p *Pool) CacheGet(ctx context.Context, key string, timeoutSeconds int64, handler CacheHandler) (value []byte, err error) {
	var (
		item proto2.CacheItem
		cur  = time.Now().Unix()
	)

	data, err := p.GetBytes(ctx, key)
	if err != nil {
		if IsNil(err) {
			return p.fetchAndSaveData(ctx, key, &item, handler)
		}
		return
	}

	err = proto.Unmarshal(data, &item)
	if err != nil {
		token, _ := p.Acquire(ctx, key, timeoutSeconds)
		if token > 0 {
			return p.fetchAndSaveData(ctx, key, &item, handler)
		}
		return
	}

	if !item.Expired(cur, timeoutSeconds) {
		return item.Data, nil
	}

	token, _ := p.Acquire(ctx, key, timeoutSeconds)
	if token == 0 {
		if item.IsEmpty() {
			return nil, nil
		}

		return item.GetData(), nil
	}

	return p.fetchAndSaveData(ctx, key, &item, handler)
}

func (p *Pool) fetchAndSaveData(ctx context.Context, key string, ci *proto2.CacheItem, handler CacheHandler) (value []byte, err error) {
	value, err = handler()
	if err != nil {
		return
	}

	if len(value) == 0 {
		ci.SaveData(proto2.Empty)
	} else {
		ci.SaveData(value)
	}

	data, err := proto.Marshal(ci)
	if err != nil {
		return
	}

	_, err = p.SetEx(ctx, key, data, DefaultCacheTimeout)
	return
}
