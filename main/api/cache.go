package api

import (
	"os"
	"strings"
	"time"
)

func readCacheFile(route string) (string, os.FileInfo, error) {
	split := strings.Split(route, "/")
	cacheFileName := "./" + split[len(split)-1]
	file, err := os.Stat(cacheFileName)
	return cacheFileName, file, err
}

func notOldEnough(file os.FileInfo) bool {
	duration := time.Duration(-2629743000000000) // one month
	//duration := time.Duration(0) // one month
	oneMonth := time.Now().Add(duration)
	return file.ModTime().Before(oneMonth)
}