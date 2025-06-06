package gredis

import "strconv"

type Param map[string]string

// Exists 是否存在
func (p Param) Exists(key string) bool {
	_, ok := p[key]
	return ok
}

// Get 获取字符串
func (p Param) Get(key string) string {
	value, _ := p[key]
	return value
}

// ToInt 获取Int
func (p Param) ToInt(key string) int {
	value, _ := strconv.Atoi(p.Get(key))
	return value
}

// ToInt64 获取Int64
func (p Param) ToInt64(key string) int64 {
	value, _ := strconv.ParseInt(p.Get(key), 10, 64)
	return value
}

// ToUint8 获取Uint8
func (p Param) ToUint8(key string) uint8 {
	return uint8(p.ToInt64(key))
}

// ToFloat64 获取Float64
func (p Param) ToFloat64(key string) float64 {
	value, _ := strconv.ParseFloat(p.Get(key), 64)
	return value
}

func (p Param) Clone() Param {
	res := make(Param, len(p))

	for k, v := range p {
		res[k] = v
	}
	return res
}
