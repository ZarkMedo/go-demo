package demo02

import "testing"

func TestMapFactory(t *testing.T) {
	mapFactory := make(map[int]func(p int) int, 3)
	mapFactory[1] = func(p int) int { return p }
	mapFactory[2] = func(p int) int { return p * p }
	mapFactory[3] = func(p int) int { return p * p * p }
	t.Logf("k1: %d; k2: %d; k3: %d", mapFactory[1](3), mapFactory[2](3), mapFactory[3](3))
	t.Logf("k1: %T; k2: %T; k3: %T", mapFactory[1], mapFactory[2], mapFactory[3])

}
func TestMapList(t *testing.T) {
	mapDemo := map[int]int{}
	mapDemo[1] = 1

	t.Log(len(mapDemo))
	delete(mapDemo, 1)
	t.Log(len(mapDemo))

}
