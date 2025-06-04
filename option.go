package gredis

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
)

var (
	DefaultOption = func() Option {
		return Option{
			Network:               "tcp",
			Host:                  "127.0.0.1",
			Port:                  6379,
			MaxRetries:            3,
			ContextTimeoutEnabled: true,
			PoolSize:              16,
			MinIdleConns:          2,
			MaxIdleConns:          8,
			MaxActiveConns:        16,
		}
	}

	JsonOption = func(data []byte) (opt Option, err error) {
		opt = DefaultOption()
		err = json.Unmarshal(data, &opt)
		return
	}

	YamlOption = func(data []byte) (opt Option, err error) {
		opt = DefaultOption()
		err = yaml.Unmarshal(data, &opt)
		return
	}
)

type Option struct {
	Network               string `json:"network" yaml:"network"`
	Host                  string `json:"host" yaml:"host"`
	Port                  uint32 `json:"port" yaml:"port"`
	Username              string `json:"username" yaml:"username"`
	Password              string `json:"password" yaml:"password"`
	DB                    uint32 `json:"DB" yaml:"DB"`
	MaxRetries            uint32 `json:"maxRetries" yaml:"maxRetries"`
	DialTimeoutSecond     uint32 `json:"dialTimeoutSecond" yaml:"dialTimeoutSecond"`
	ReadTimeoutSecond     uint32 `json:"readTimeoutSecond" yaml:"readTimeoutSecond"`
	WriteTimeoutSecond    uint32 `json:"writeTimeoutSecond" yaml:"writeTimeoutSecond"`
	ContextTimeoutEnabled bool   `json:"contextTimeoutEnabled" yaml:"contextTimeoutEnabled"`
	PoolSize              uint32 `json:"poolSize" yaml:"poolSize"`
	PoolWaitTimeoutSecond uint32 `json:"poolWaitTimeoutSecond" yaml:"poolWaitTimeoutSecond"`
	MinIdleConns          uint32 `json:"minIdleConns" yaml:"minIdleConns"`
	MaxIdleConns          uint32 `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxActiveConns        uint32 `json:"maxActiveConns" yaml:"maxActiveConns"`
	ConnMaxIdleTimeSecond uint32 `json:"connMaxIdleTimeSecond" yaml:"connMaxIdleTimeSecond"`
	ConnMaxLifetimeSecond uint32 `json:"connMaxLifetimeSecond" yaml:"connMaxLifetimeSecond"`
	TLSConfig             *tls.Config
	id                    string
}

func (o *Option) NewClient() *redis.Client {
	return redis.NewClient(o.ToOptions())
}

func (o *Option) Id() string {
	if o.id != "" {
		return o.id
	}

	o.id = fmt.Sprintf("%s:%d:%d", o.Host, o.Port, o.DB)
	return o.id
}

func (o *Option) ToOptions() *redis.Options {
	return &redis.Options{
		Network:               o.Network,
		Addr:                  o.Addr(),
		Username:              o.Username,
		Password:              o.Password,
		DB:                    int(o.DB),
		MaxRetries:            int(o.MaxRetries),
		DialTimeout:           o.DialTimeout(),
		ReadTimeout:           o.ReadTimeout(),
		WriteTimeout:          o.WriteTimeout(),
		ContextTimeoutEnabled: o.ContextTimeoutEnabled,
		PoolSize:              int(o.PoolSize),
		PoolTimeout:           o.PoolWaitTimeout(),
		MinIdleConns:          int(o.MinIdleConns),
		MaxIdleConns:          int(o.MaxIdleConns),
		MaxActiveConns:        int(o.MaxActiveConns),
		ConnMaxIdleTime:       o.ConnMaxIdleTime(),
		ConnMaxLifetime:       o.ConnMaxLifetime(),
		TLSConfig:             o.TLSConfig,
	}
}

func (o *Option) Addr() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o *Option) DialTimeout() time.Duration {
	if o.DialTimeoutSecond < 1 {
		return time.Second * 3
	}
	return time.Second * time.Duration(o.DialTimeoutSecond)
}

func (o *Option) ReadTimeout() time.Duration {
	if o.ReadTimeoutSecond < 1 {
		return time.Second * 3
	}
	return time.Second * time.Duration(o.ReadTimeoutSecond)
}

func (o *Option) WriteTimeout() time.Duration {
	if o.WriteTimeoutSecond < 1 {
		return time.Second * 3
	}
	return time.Second * time.Duration(o.WriteTimeoutSecond)
}

func (o *Option) PoolWaitTimeout() time.Duration {
	if o.PoolWaitTimeoutSecond < 1 {
		return time.Second * 20
	}

	return time.Second * time.Duration(o.PoolWaitTimeoutSecond)
}

func (o *Option) ConnMaxIdleTime() time.Duration {
	if o.ConnMaxIdleTimeSecond < 1 {
		return time.Second * 60
	}
	return time.Second * time.Duration(o.ConnMaxIdleTimeSecond)
}

func (o *Option) ConnMaxLifetime() time.Duration {
	if o.ConnMaxLifetimeSecond < 1 {
		return time.Second * 60
	}

	return time.Second * time.Duration(o.ConnMaxLifetimeSecond)
}
