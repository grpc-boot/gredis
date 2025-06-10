package gredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (p *Pool) GeoPos(ctx context.Context, key string, members ...string) ([]*redis.GeoPos, error) {
	var (
		cmd = p.client.GeoPos(ctx, key, members...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) GeoAdd(ctx context.Context, key string, geoLocations ...*redis.GeoLocation) (newCount int64, err error) {
	var (
		cmd = p.client.GeoAdd(ctx, key, geoLocations...)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) GeoDist(ctx context.Context, key string, member1, member2, unit string) (float64, error) {
	var (
		cmd = p.client.GeoDist(ctx, key, member1, member2, unit)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	var (
		cmd = p.client.GeoRadius(ctx, key, longitude, latitude, query)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}

func (p *Pool) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) ([]string, error) {
	var (
		cmd = p.client.GeoSearch(ctx, key, q)
	)

	if !IsNil(cmd.Err()) {
		WriteLog(cmd.Err(), cmd.String(), p.opt)
	}

	return cmd.Result()
}
