package handler

import (
	"net/http"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo/api/result"
)

// buildInfo 缓存构建信息，避免每次请求都读取
var buildInfo = struct {
	commit    string
	goVersion string
	buildTime string
	initTime  time.Time
}{
	initTime: time.Now(),
}

// init 在包初始化时读取构建信息
func init() {
	if info, ok := debug.ReadBuildInfo(); ok {
		// 获取提交哈希
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				buildInfo.commit = setting.Value
			case "vcs.time":
				buildInfo.buildTime = setting.Value
			}
		}
		buildInfo.goVersion = info.GoVersion
	}
}

// getBuildMode 检测当前运行模式
func getBuildMode() string {
	// 如果没有commit信息，很可能是使用go run运行的
	if buildInfo.commit == "" {
		return "development (go run)"
	}

	// 检查是否是调试模式
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "-ldflags" {
				// 如果有ldflags，说明是正式构建
				return "release"
			}
		}
	}

	return "unknown"
}

// Health 处理健康检查请求，返回服务器状态信息
func Health(c *gin.Context) {
	// 获取运行时信息
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// 构建响应数据
	data := gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"uptime":    time.Since(buildInfo.initTime).String(),
		"version": gin.H{
			"commit":     buildInfo.commit,
			"go_version": buildInfo.goVersion,
			"build_time": buildInfo.buildTime,
			"build_mode": getBuildMode(),
		},
		"runtime": gin.H{
			"go_os":         runtime.GOOS,
			"go_arch":       runtime.GOARCH,
			"num_goroutine": runtime.NumGoroutine(),
			"num_cpu":       runtime.NumCPU(),
		},
		"memory": gin.H{
			"alloc":       m.Alloc,
			"total_alloc": m.TotalAlloc,
			"sys":         m.Sys,
			"num_gc":      m.NumGC,
		},
	}

	c.JSON(http.StatusOK, result.BaseResult{
		Code:    http.StatusOK,
		Success: true,
		Message: "Server Online",
		Data:    data,
	})
}
