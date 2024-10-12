package main

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "张三",
		Age:   20,
		Skill: "打球",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，期望为=%v  实际为=%v", true, res)
	}
	t.Logf("monster.Store()测试成功!")
}

func TestReStore(t *testing.T) {
	var monster Monster
	monster.ReStore()
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，期望为=%v  实际为=%v", true, res)
	}

	if monster.Name != "张三" {
		t.Fatalf("monster.ReStore() 错误，期望为=%v  实际为=%v", "张三", monster.Name)
	}
	t.Logf("monster.ReStore()测试成功!")
}
