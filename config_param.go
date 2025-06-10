package gredis

import "strconv"

type ConfigParam map[string]string

// Exists 是否存在
func (p ConfigParam) Exists(key string) bool {
	_, ok := p[key]
	return ok
}

// Get 获取字符串
func (p ConfigParam) Get(key string) string {
	value, _ := p[key]
	return value
}

// ToBool 获取Bool
func (p ConfigParam) ToBool(key string) bool {
	value, _ := strconv.ParseBool(p.Get(key))
	return value
}

// ToInt 获取Int
func (p ConfigParam) ToInt(key string) int {
	value, _ := strconv.Atoi(p.Get(key))
	return value
}

// ToInt64 获取Int64
func (p ConfigParam) ToInt64(key string) int64 {
	value, _ := strconv.ParseInt(p.Get(key), 10, 64)
	return value
}

// ToUint8 获取Uint8
func (p ConfigParam) ToUint8(key string) uint8 {
	return uint8(p.ToInt64(key))
}

// ToFloat64 获取Float64
func (p ConfigParam) ToFloat64(key string) float64 {
	value, _ := strconv.ParseFloat(p.Get(key), 64)
	return value
}

func (p ConfigParam) Clone() ConfigParam {
	res := make(ConfigParam, len(p))

	for k, v := range p {
		res[k] = v
	}
	return res
}
