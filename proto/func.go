package proto

import (
	"errors"
	"time"
	"unsafe"

	"github.com/grpc-boot/gredis/base"

	"github.com/goccy/go-json"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	EmptyTag = `result:EmPty`
	Forever  = 0
)

var (
	ErrEmpty = errors.New(`value is empty`)
	Empty    = []byte(EmptyTag)
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

func (ci *CacheItem) EffectiveData() []byte {
	if ci.IsEmpty() {
		return []byte{}
	}

	return ci.GetData()
}

func (ci *CacheItem) UnmarshalProtoData(out proto.Message) error {
	if ci.IsEmpty() {
		return ErrEmpty
	}

	return proto.Unmarshal(ci.GetData(), out)
}

func (ci *CacheItem) MapData() (data base.MapData, err error) {
	if ci.IsEmpty() {
		err = ErrEmpty
		return
	}

	var out = &structpb.Struct{}
	err = proto.Unmarshal(ci.GetData(), out)
	if err != nil {
		return
	}

	data = out.AsMap()
	return
}

func (ci *CacheItem) UnmarshalJsonData(out any) error {
	if ci.IsEmpty() {
		return ErrEmpty
	}

	return json.Unmarshal(ci.GetData(), out)
}

func (ci *CacheItem) SaveData(data []byte) *CacheItem {
	if len(data) == 0 {
		ci.Data = Empty
	} else {
		ci.Data = data
	}

	ci.UpdatedAt = time.Now().Unix()
	if ci.CreatedAt == 0 {
		ci.CreatedAt = ci.GetUpdatedAt()
	}
	ci.UpdatedCount++
	return ci
}

func (ci *CacheItem) Marshal() ([]byte, error) {
	return proto.Marshal(ci)
}

func bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
