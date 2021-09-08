package demo02

import "testing"

func TestArrMethod(t *testing.T) {
	/* 初始化值为0*/
	var arr1 [3]int
	t.Log(arr1)
	arr2 := [3]int{1, 2, 3}
	t.Log(arr2)
	arr3 := [...]int{1, 2, 3, 4}
	t.Log(arr3)
}

func TestArrMethod1(t *testing.T) {
	arr := [...]int{1, 2, 3}

	for idx, val := range arr {
		t.Logf("idx: %T, val: %T", idx, val)
		t.Log(idx, val)

	}
	for _, val := range arr {

		t.Log(val)
	}

}
func TestArrMethod2(t *testing.T) {
	arr := [...]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
}

func TestSliceMethod(t *testing.T) {
	var slc1 []int
	t.Log("slc1: ", len(slc1), cap(slc1))

	slc2 := []int{1, 2, 3}
	t.Log("slc2: ", len(slc2), cap(slc2))

}
