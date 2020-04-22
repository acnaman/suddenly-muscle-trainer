package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skratchdot/open-golang/open"
)

const percent int = 5
const blankMinute = 10
const logName = "test.log"

type program struct{}

func main() {
	logger := NewLogger(logName)

	fmt.Println("Muscle Training Runner Start...")
	logger.WriteStartLog()
	t := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-t.C:
			if !isLucky() {
				logger.WriteUnluckyLog()
				break
			}

			url := getRandomURL()
			openVideo(url)
			logger.WriteVideoPlayedLog(url)
		}
	}
}

func getRandomURL() string {
	videoIdList := []string{
		"HF7H6M4nzNY",
		"s4jzFWoRRA0",
		"jK_8IgcgBHo",
		"vJ_NBi0YuPM",
		"yqQM3qPoQsk",
		"MHwzwXPzIzI",
		"MByVZoPO6Ds",
	}

	index := generateRandomInteger(len(videoIdList))

	templateURL := "https://www.youtube.com/watch?v="

	return templateURL + videoIdList[index]
}

func openVideo(url string) {

	err := open.Run(url)
	if err != nil {
		panic(err)
	}
}

func isLucky() bool {
	rand := generateRandomInteger(100 / percent)
	if rand == 1 {
		return true
	}
	return false
}

func generateRandomInteger(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
