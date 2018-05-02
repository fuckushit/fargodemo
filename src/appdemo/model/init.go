package model

import (
	"appdemo/appkey"
	"bdlib/config"
	"bdlib/crypt"
	"bdlib/mysql"
	"bdlib/redis"
	"bdlib/util"
	"fargo"
	"fmt"
)

// DefaultDBManager mysql manger
var DefaultDBManager *util.DBManager

// DefaultCache default redis cache
var DefaultCache *redis.RedisManager

// Init 初始化
func Init(cfg config.Configer) (err error) {

	// 初始化mola
	if err = initMolaDB(cfg); err != nil {
		fargo.Error(err)
		return
	}

	// 初始化redis
	if err = initRedisCache(cfg); err != nil {
		fargo.Error(err)
		return
	}

	if err = initMysql(cfg); err != nil {
		fargo.Error(err)
		return
	}

	return
}

// 初始化mysql
func initMysql(cfg config.Configer) (err error) {
	DefaultDBManager = &util.DBManager{}

	poolsze, err := cfg.GetIntSetting("appdatabase", "poolsize", int64(10))
	if err != nil {
		fargo.Error(err)
		return
	}
	enPass, err := cfg.GetSetting("appdatabase", "pass")
	if err != nil {
		fargo.Error(err)
		return
	}

	pass, err := crypt.Decrypt(enPass, appkey.DBKey)
	if err != nil {
		fargo.Error(err)
		return
	}

	section, err := cfg.GetSection("appdatabase")
	if err != nil {
		fargo.Error(err)
		return
	}
	section["pass"] = pass

	DefaultDBManager.Pool = &util.DBStore{DBPool: make(chan *mysql.DB, poolsze)}
	if DefaultDBManager.Pool.DBPool, err = mysql.NewConnectionPool(section, int(poolsze)); err != nil {
		fargo.Error(err)
		return
	}

	return
}

// 初始化redis cache
func initRedisCache(cfg config.Configer) (err error) {
	sec, err := cfg.GetSection("appredis")
	if err != nil {
		return
	}
	auth, err := sec.GetValue("auth")
	if err != nil {
		return
	}
	if auth != "" {
		if auth, err = crypt.Decrypt(auth, appkey.RedisKey); err != nil {
			return
		}
	}
	if DefaultCache, err = util.NewRedisManagerTimeout(sec); err != nil {
		err = fmt.Errorf("init redis cache err: %v", err)
		return
	}
	return
}

// TODO 初始化mola
func initMolaDB(cfg config.Configer) (err error) {
	return
}

// GetDefaultDB _
func GetDefaultDB() (db *mysql.DB) {
	db = DefaultDBManager.GetDB()
	return
}

// PutDefaultDB _
func PutDefaultDB(db *mysql.DB) {
	DefaultDBManager.PutDB(db)
	return
}
