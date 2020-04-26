package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/kardianos/service"
	"github.com/skratchdot/open-golang/open"
)

const logName = "test.log"

type program struct {
	intervalTime int
	parcentage   int

	exit chan struct{}
}

var logger service.Logger

var mtlogger *MTLogger

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	if service.Interactive() {
		//logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	mtlogger.WriteStartLog()
	go p.run()
	return nil
}

func (p *program) run() {
	fmt.Println("Muscle Training Runner Start...")
	t := time.NewTicker(time.Duration(p.intervalTime) * time.Minute)
	for {
		select {
		case <-t.C:
			if !isLucky(p.parcentage) {
				mtlogger.WriteUnluckyLog()
				break
			}

			url := getRandomURL()
			openVideo(url)
			mtlogger.WriteVideoPlayedLog(url)
		}
	}
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	mtlogger.WriteStopLog()
	close(p.exit)
	return nil
}

func main() {
	mtlogger = NewLogger("muscletrainer.log")

	var (
		intervalTime int
		parcentage   int
	)
	flag.IntVar(&intervalTime, "t", 30, "Interval of Events")
	flag.IntVar(&parcentage, "p", 5, "Parcentage of Video Play")

	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "SuddenlyMuscleTrainer",
		DisplayName: "Suddenly Muscle Trainer",
		Description: "This service suddenly plays Muscle Training video.",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}

	prg := &program{
		intervalTime: intervalTime,
		parcentage:   parcentage,
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
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

func isLucky(percent int) bool {
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
