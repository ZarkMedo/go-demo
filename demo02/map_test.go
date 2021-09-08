package demo02

import "testing"

func TestCreateMap(t *testing.T) {
	m1 := map[string]int{"key1": 1, "key2": 2}
	t.Log(m1)
	t.Logf("len: %d, ", len(m1))
	m2 := map[string]int{}
	t.Logf("len: %d, ", len(m2))
	m2["key1"] = 2
	t.Logf("len: %d, ", len(m2))
	// make
	m3 := make(map[string]string, 1)
	t.Logf("len: %d, ", len(m3))
	t.Log(m3["test"])
	m3["test2"] = "t2"
	m3["test1"] = "t1"
	t.Log(m3["test1"])
	t.Logf("len: %d, ", len(m3))

}

func TestMapKeyExist(t *testing.T) {
	m1 := map[string]int{"key1": 1}
	key := "key2"
	if value, isExist := m1[key]; isExist {
		t.Logf("key: %s, value: %d", key, value)
	} else {
		t.Logf("key: %s, value: 不存在。。。", key)
	}

}
func TestMapLoop(t *testing.T) {
	m := make(map[string]int, 10)
	for i := 1; i <= 10; i++ {
		m[string(i)] = i
	}
	t.Log(m)
	for key, val := range m {
		t.Logf("key: %s, value: %d", key, val)
	}
}
