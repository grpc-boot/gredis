package base

type MapData map[string]any

func (md MapData) Exists(key string) bool {
	_, ok := md[key]
	return ok
}

func (md MapData) Get(key string) string {
	value, _ := md[key].(string)
	return value
}

func (md MapData) GetInt(key string) int {
	return int(md.GetInt64(key))
}

func (md MapData) GetInt64(key string) int64 {
	value, _ := md[key].(float64)
	return int64(value)
}

func (md MapData) GetUint8(key string) uint8 {
	return uint8(md.GetInt64(key))
}

func (md MapData) GetBool(key string) bool {
	value, _ := md[key].(bool)
	return value
}

func (md MapData) GetFloat64(key string) float64 {
	value, _ := md[key].(float64)
	return value
}

func (md MapData) Clone() MapData {
	res := make(MapData, len(md))

	for k, v := range md {
		res[k] = v
	}
	return res
}
