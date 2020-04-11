package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skratchdot/open-golang/open"
)

//var logger service.Logger

const percent int = 5
const blankMinute = 10

type program struct{}

/*func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	t := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-t.C:
			if !isLucky() {
				fmt.Println("Unlucky...")
				break
			}
			openVideo()
		}
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}*/

func main() {
	/*	svcConfig := &service.Config{
			DisplayName: "Muscle Training Time!!!",
			Description: "Sudeenly Muscle Training service.",
			Name:        "SuddenlyMuscleTraining",
		}

		prg := &program{}
		s, err := service.New(prg, svcConfig)
		if err != nil {
			log.Fatal(err)
		}
		logger, err = s.Logger(nil)
		if err != nil {
			log.Fatal(err)
		}
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}*/
	fmt.Println("Muscle Training Runner Start...")
	t := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-t.C:
			if !isLucky() {
				fmt.Println("Unlucky...")
				break
			}

			openVideo()
		}
	}
}

func openVideo() {
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

	err := open.Run(templateURL + videoIdList[index])
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
