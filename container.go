package gredis

import (
	"hash/crc32"
	"math"
	"math/rand"
	"sync"
)

type Key interface {
	~string | ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint | ~int | ~uint64 | ~int64
}

var (
	Uint32Hash = crc32.ChecksumIEEE
	_container sync.Map
)

func KeyHash(key interface{}) uint32 {
	switch value := key.(type) {
	case int:
		return uint32(value & math.MaxUint32)
	case int8:
		return uint32(value) & math.MaxUint32
	case int16:
		return uint32(value) & math.MaxUint32
	case int32:
		return uint32(value) & math.MaxUint32
	case int64:
		return uint32(value) & math.MaxUint32
	case uint:
		return uint32(value) & math.MaxUint32
	case uint8:
		return uint32(value) & math.MaxUint32
	case uint16:
		return uint32(value) & math.MaxUint32
	case uint32:
		return uint32(value) & math.MaxUint32
	case uint64:
		return uint32(value) & math.MaxUint32
	case string:
		return Uint32Hash(String2Bytes(value))
	}
	return 0
}

func Put(key string, opts ...Option) {
	pl := make([]*Pool, len(opts))

	for index, opt := range opts {
		pl[index] = NewPool(&opt)
	}

	_container.Store(key, pl)
}

func Get(key string) (p *Pool) {
	var (
		value, _ = _container.Load(key)
		list, _  = value.([]*Pool)
	)

	if len(list) == 0 {
		return nil
	}

	if len(list) == 1 {
		return list[0]
	}

	p = list[rand.Intn(len(list))]
	return
}

func GetWithIndex(key string, index int) (p *Pool) {
	var (
		value, _ = _container.Load(key)
		list, _  = value.([]*Pool)
	)

	p = list[index]
	return
}

func GetWithShard[K Key](containerKey string, shardKey K) (p *Pool) {
	var (
		value, _ = _container.Load(containerKey)
		list, _  = value.([]*Pool)
	)

	if len(list) == 0 {
		return nil
	}

	if len(list) == 1 {
		return list[0]
	}

	index := int(KeyHash(shardKey) % uint32(len(list)))
	return list[index]
}
