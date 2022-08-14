package zlogr

import "github.com/go-logr/logr"

const (
	TimezoneAsiaJakarta = `Asia/Jakarta`
)

var (
	// Log Singleton Zerolog instance
	Log logr.Logger
)
