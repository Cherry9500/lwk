package main

import (
	"testing"
)

// go test -v -test.run TestAddUpper  指定测试单个函数
// go test -v cal_test.go cal.go   指定测试单个文件
// 测试用例文件必须以_test.go结尾
func TestGetSub(t *testing.T) { //形参必须是*testing.T
	res := getsub(10, 3)
	if res != 7 {
		t.Fatalf("getsub(10, 3)执行错误，期望值=%v  实际值=%v", 7, res)
	}
	t.Logf("getsub(10, 3)执行正确")
}
