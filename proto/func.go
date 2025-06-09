package proto

import (
	"time"
	"unsafe"
)

const (
	EmptyTag = `result:EmPty`
	Forever  = 0
)

var (
	Empty = []byte(EmptyTag)
)

func (ci *CacheItem) Expired(cur, timeoutSeconds int64) bool {
	if timeoutSeconds == Forever {
		return false
	}

	return cur-ci.GetUpdatedAt() > timeoutSeconds
}

func (ci *CacheItem) IsEmpty() bool {
	return len(ci.GetData()) == 0 || bytes2String(ci.GetData()) == EmptyTag
}

func (ci *CacheItem) SaveData(data []byte) *CacheItem {
	ci.Data = data
	ci.UpdatedAt = time.Now().Unix()
	if ci.CreatedAt == 0 {
		ci.CreatedAt = ci.GetUpdatedAt()
	}
	ci.UpdatedCount++
	return ci
}

func bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
