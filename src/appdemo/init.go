package main

import (
	"appdemo/model"
	"bdlib/config"
	"fargo"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// Init 初始化资源（包括conf, client, load data ...）
func Init() (err error) {

	// 初始化local cahce
	if err = initLocalCache(fargo.GCfg); err != nil {
		fargo.Error(err)
		return
	}

	// 初始化file cahce
	if err = initFileCache(fargo.GCfg); err != nil {
		fargo.Error(err)
		return
	}

	// 初始化db
	if err = model.Init(fargo.GCfg); err != nil {
		fargo.Error(err)
		return
	}

	// pprof信息
	initProfile(fargo.GCfg)

	return
}

// 初始化profile信息
func initProfile(cfg config.Configer) {
	if profile, _ := cfg.GetBoolSetting("web", "enableCPUProfile", false); profile {
		// 加入memprofile追踪内存状态
		go memprofile()
	}

	if profile, _ := cfg.GetBoolSetting("web", "enableCPUProfile", false); profile {
		// cpu profile
		go cpuprofile()
	}
}

// TODO 初始化file cache
func initFileCache(cfg config.Configer) (err error) {
	return
}

// TODO 初始化localcache
func initLocalCache(cfg config.Configer) (err error) {
	return
}

func cpuprofile() {

	for i := 0; i < 60; i++ {
		filename := fmt.Sprintf("/tmp/appdemo_cpu_%d%d.pprof", i/10, i%10)
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fargo.Error(err)
			return
		}
		pprof.StartCPUProfile(file)
		time.Sleep(3 * time.Minute)
		pprof.StopCPUProfile()
		file.Close()
	}
}

func memprofile() {
	// 每10分钟写一次
	runtime.SetBlockProfileRate(1)
	for i := 0; i < 60; i++ {
		filename := fmt.Sprintf("/tmp/appdemo_mem_%d%d.pprof", i/10, i%10)
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fargo.Error(err)
			return
		}
		time.Sleep(3 * time.Minute)
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}
