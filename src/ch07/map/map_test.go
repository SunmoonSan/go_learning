package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Log("Len m1", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 3)
	t.Logf("Len m3=%d", len(m3))
	m3[0] = 0
	m3[1] = 1
	m3[2] = 2
	m3[3] = 3
	t.Log(m3, len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[4]; ok {
		t.Log(v)
	} else {
		t.Log("不存在...")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 4: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
