package main

import (
	"testing"
)

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

func TestIsValidTime(t *testing.T) {
	actual := isValidTime("00:00", "00:00")
	expect := true
	if actual != expect {
		t.Error("unavailable")
	}
}

func TestGetRondomURL(t *testing.T) {
	actual := getRandomURL(nil)
	if actual == "" {
		t.Error("getRondomURL returned \"\"")
	}

	actual2 := getRandomURL([]string{})
	if actual2 == "" {
		t.Error("getRondomURL returned \"\"")
	}

	actual3 := getRandomURL([]string{"http://test.com"})
	if actual3 != "http://test.com" {
		t.Error("getRondomURL returned not listed item")
	}
}

func TestIsLucky(t *testing.T) {
	testfunc := func(parcent int) {
		var luckyNum float64 = 0
		for i := 0; i < 100000; i++ {
			if isLucky(parcent) {
				luckyNum++
			}
		}
		expect := float64(parcent)
		actual := luckyNum / 1000.0

		if actual < expect-1 || expect+1 < actual {
			t.Errorf("expect:%v％, actual:%v％\n", parcent, actual)
		}
	}

	testfunc(5)
	testfunc(0)
	testfunc(1)
	testfunc(100)
	testfunc(30)

}
