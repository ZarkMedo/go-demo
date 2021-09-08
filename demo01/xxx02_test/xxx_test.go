package xxx02_test

import (
	"testing"
)

func TestXxx(t *testing.T) {
	a := 1
	aAddr := &a
	t.Log(a, aAddr)
	t.Logf("a: %T aAddr: %T", a, aAddr)

}

func TestArr(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 3, 3, 4}
	t.Log(a == b)
	t.Log(b == c)
}

// for 循环测试

func TestForMethod(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i)

	}
}

func TestIfMethod(t *testing.T) {
	if i := 1; i == 0 {
		t.Log("print ok")

	} else if i := 2; i == 2 {

	}

}

func TestSwitchMethod(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1, 2:
			t.Log("第一名")
		case 3, 4:
			t.Log("第二名")
		default:
			t.Log("其他")

		}

	}
}
func TestSwitchMethod2(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i>>2 == 0 && i >= 4:
			t.Logf("case 1: %d", i)
		case i>>2 == 1:
			t.Logf("case 2: %d", i)
		default:
			t.Logf("default: %d", i)

		}

	}
}
