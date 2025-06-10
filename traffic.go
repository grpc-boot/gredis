package gredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	leakyBucketScript = redis.NewScript(`
		local redKey      = KEYS[1]
		local capacity    = tonumber(ARGV[1])
		local current     = tonumber(ARGV[2])
		local speed       = tonumber(ARGV[3])
		local reqNum      = tonumber(ARGV[4])
		local keyTimeout  = tonumber(ARGV[5])

        if reqNum > capacity then
			return 0
		end
		
		local bucketInfo = redis.call('HMGET', redKey, 'laTime', 'rtCount')
		if not bucketInfo[1] then
			redis.call('HMSET', redKey, 'laTime', current, 'rtCount', capacity - reqNum)
			redis.call('EXPIRE', redKey, keyTimeout)
			return 1
		end
		
		local lastAddTime    = tonumber(bucketInfo[1])
		local remainTokenNum = tonumber(bucketInfo[2])
		local addTokenNum    = (current - lastAddTime) * speed
		if addTokenNum > 0 then
			lastAddTime    = current
			remainTokenNum = math.min(addTokenNum + remainTokenNum, capacity)
		end
		
		if reqNum > remainTokenNum then
			return 0
		end

		redis.call('HMSET', redKey, 'laTime', lastAddTime, 'rtCount', remainTokenNum - reqNum)
		redis.call('EXPIRE', redKey, keyTimeout)
		return 1`)

	incrLimitScript = redis.NewScript(`
		local key = KEYS[1]
		local increment = tonumber(ARGV[1])
		local ttl = tonumber(ARGV[2])
		 
		local newVal = redis.call("INCRBY",  key, increment)
		if newVal == increment then 
			redis.call("EXPIRE",  key, ttl)
		end 
		return newVal`)
)

func (p *Pool) AcquireByLeakyBucket(ctx context.Context, key string, current int64, capacity, speed, acqNum, keyTimeoutSecond int) (ok bool, err error) {
	cmd := leakyBucketScript.Run(ctx, p.client, []string{key}, capacity, current, speed, acqNum, keyTimeoutSecond)
	val, err := cmd.Int64()
	if err != nil {
		if !IsNil(err) {
			WriteLog(cmd.Err(), cmd.String(), p.opt)
		}
		return
	}
	return val == 1, err
}

func (p *Pool) AcquireByIncr(ctx context.Context, key string, acqNum, limitNum int64, keyTimeoutSecond int) (ok bool, err error) {
	cmd := incrLimitScript.Run(ctx, p.client, []string{key}, acqNum, keyTimeoutSecond)
	val, err := cmd.Int64()
	if err != nil {
		if !IsNil(err) {
			WriteLog(cmd.Err(), cmd.String(), p.opt)
		}
		return
	}

	return val <= limitNum, err
}
