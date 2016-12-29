package main

import (
	"log"
	"time"

	"gopkg.in/redis.v2"

	"github.com/seefan/gossdb"
	"github.com/xlvector/dlog"
)

const (
	SAVE_TYPE_SSDB  = 1
	SAVE_TYPE_REDIS = 2
)

type Model struct {
	pool     *gossdb.Connectors
	client   *gossdb.Client
	rediscli *redis.Client

	conf     *Config
	saveType int
}

func NewModel(conf *Config) *Model {
	if conf == nil {
		log.Fatal("config fail!")
		return nil
	}

	if conf.Type == "redis" {
		client := redis.NewTCPClient(&redis.Options{
			Addr:        conf.RedisAddr,
			DB:          conf.Db,
			DialTimeout: time.Duration(60) * time.Second,
		})
		return &Model{
			rediscli: client,
			conf:     conf,
			saveType: SAVE_TYPE_REDIS,
		}

	} else if conf.Type == "ssdb" {

		pool, err := gossdb.NewPool(&gossdb.Config{
			Host:             conf.Host,
			Port:             conf.Port,
			MinPoolSize:      5,
			MaxPoolSize:      50,
			AcquireIncrement: 5,
		})
		if err != nil {
			log.Fatal(err)
			return nil
		}

		c, err := pool.NewClient()
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &Model{
			pool:     pool,
			client:   c,
			conf:     conf,
			saveType: SAVE_TYPE_SSDB,
		}
	} else {
		log.Fatal("config type fail!")
		return nil
	}

}

func (p *Model) SavePx(px string) (int64, error) {
	switch p.saveType {
	case SAVE_TYPE_SSDB:
		sz, err := p.client.Qpush(p.conf.Key, px)
		dlog.Info("savepx: %d %v %v", sz, err, px)
		return sz, err
	case SAVE_TYPE_REDIS:
		cmd := p.rediscli.LPush(p.conf.Key, px)
		sz, err := cmd.Result()
		dlog.Info("savepx: %d %v %v", sz, err, px)
		return sz, err
	}
	return 0, nil
}
