package utils

import (
	"fmt"
	"os"
	"time"
)

const (
	// "2020-02-11T14:06:42.520978+0800"
	layout string = "2006-01-02T15:04:05.000000+0800"
)

var location *time.Location

func init() {
	location, _ = time.LoadLocation("Asia/Shanghai")
}

// TimeStamp TimeParser
func TimeStamp(val string) int64 {
	t, err := time.ParseInLocation(layout, val, location)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%% %v\n", err)
	}
	return t.Unix()
}