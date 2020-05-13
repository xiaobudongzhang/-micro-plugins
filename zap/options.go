package zap

import "go.uber.org/zap"

type Options struct {
	zap.Config
	LogFileDir string `json:"logFileDir"`
	AppName string `json:"appName"`
	ErrorFileName string `json:"errorFileName"`
	WarnFileName string `json:"warnFileName"`
	InfoFileName string `json:"infoFileName"`
	DebugFileName string `json:"debugFileName"`
	MaxSize int `json:"maxSize"`
	MaxBackups int `json:"maxBackups"`
	MaxAge int `json:"maxAge"`
}