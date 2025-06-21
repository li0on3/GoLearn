package main

import (
	"testing"
	"time"
)

func Test_getTime(t *testing.T) {

	if ans := getTime(); ans != time.DateTime {
		t.Errorf("获取当前时间失败 %s", ans)
	}

}

func Test_getVersion(t *testing.T) {
	if ans := getVersion(); ans != "go1.24.3" {
		t.Errorf("获取当前Golang版本失败 %s", ans)
	}
}
