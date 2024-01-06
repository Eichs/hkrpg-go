package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Eichs/hkrpg-go/gdconf"
	"github.com/Eichs/hkrpg-go/internal"
	"github.com/Eichs/hkrpg-go/internal/Net"
	"github.com/Eichs/hkrpg-go/pkg/config"
	"github.com/Eichs/hkrpg-go/pkg/logger"
)

func main() {
	// 启动读取配置
	err := config.LoadConfig()
	if err != nil {
		if err == config.FileNotExist {
			p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
			cf, _ := os.Create("./config.json")
			cf.Write(p)
			cf.Close()
			fmt.Printf("找不到配置文件\n已生成默认配置文件 config.json \n")
			main()
		} else {
			panic(err)
		}
	}
	// 初始化日志
	logger.InitLogger()
	logger.SetLogLevel(strings.ToUpper(config.GetConfig().LogLevel))
	logger.Info("hkrpg-go")

	cfg := config.GetConfig()
	// 初始化
	newserver := internal.NewServer(cfg)
	if newserver == nil {
		logger.Error("服务器初始化失败")
		return
	}
	gdconf.InitGameDataConfig()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 启动SDK服务
	go func() {
		if err := newserver.Start(); err != nil {
			logger.Error("无法启动SDK服务器")
		}
	}()

	// 启动game服务
	go func() {
		if err := Net.Run(); err != nil {
			logger.Error("无法启动GAME服务器")
		}
	}()

	go func() {
		select {
		case <-done:
			logger.Info("SDK服务正在关闭")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := newserver.Shutdown(ctx); err != nil {
				logger.Error("无法正常关闭HTTP服务")
			}
			logger.Info("SDK服务已停止")
			logger.CloseLogger()
			os.Exit(0)
		}
	}()
	select {}
}
