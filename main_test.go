package main

import "testing"

func TestGenerateRandomInteger(t *testing.T) {
	var list [10]int
	for i := 0; i < len(list); i++ {
		rand := generateRandomInteger(100)
		if rand < 0 || 100 < rand {
			t.Error("over range number:", rand)
		}
		list[i] = rand
	}

	first := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] != first {
			break
		}
		if i == len(list)-1 {
			// 10^9の確率で失敗する
			t.Error("all number is equal:", first)
		}
	}
}
