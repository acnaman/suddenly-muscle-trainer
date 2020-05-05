package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/service"
	"github.com/skratchdot/open-golang/open"

	"github.com/acnaman/suddenly-muscle-trainer/runtime/setting"
)

type program struct {
	*setting.Setting
	exit chan struct{}
}

var logger service.Logger

var mtlogger *MTLogger

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	mtlogger = NewLogger(path.Join(getExecDir(), "muscletrainer.log"))

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

	prg := &program{}

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

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	if service.Interactive() {
		//logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	mtlogger.WriteStartLog()

	// load setting file
	p.Setting = setting.GetCurrentSetting()

	var validTimestr string
	if p.StartTime == p.EndTime {
		validTimestr = "All Time"
	} else {
		validTimestr = p.StartTime + "〜" + p.EndTime
	}
	mtlogger.WriteString(fmt.Sprintf("Setting information: Interval=%dmin, Percentage=%d％, validTiume=%s", p.IntervalTime, p.Parcentage, validTimestr))

	go p.run()
	return nil
}

func (p *program) run() {
	fmt.Println("Muscle Training Runner Start...")
	t := time.NewTicker(time.Duration(p.IntervalTime) * time.Minute)
	for {
		select {
		case <-t.C:
			if !isValidTime(p.StartTime, p.EndTime) {
				mtlogger.WriteInvalidTimeLog()
				break
			}
			if !isLucky(p.Parcentage) {
				mtlogger.WriteUnluckyLog()
				break
			}

			url := getRandomURL(p.VideoList)
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

func getExecDir() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exe)
}

func getRandomURL(videoList setting.VideoList) string {
	var index int

	if videoList == nil || len(videoList) == 0 {
		videoIdList := []string{
			"https://www.youtube.com/watch?v=HF7H6M4nzNY",
			"https://www.youtube.com/watch?v=s4jzFWoRRA0",
			"https://www.youtube.com/watch?v=jK_8IgcgBHo",
			"https://www.youtube.com/watch?v=vJ_NBi0YuPM",
			"https://www.youtube.com/watch?v=yqQM3qPoQsk",
			"https://www.youtube.com/watch?v=MHwzwXPzIzI",
			"https://www.youtube.com/watch?v=MByVZoPO6Ds",
		}
		index = generateRandomInteger(len(videoIdList))

		return videoIdList[index]
	}

	index = generateRandomInteger(len(videoList))

	return videoList[index]

}

func openVideo(url string) {

	err := open.Run(url)
	if err != nil {
		panic(err)
	}
}

func isLucky(percent int) bool {
	rand := generateRandomInteger(100)
	if rand < percent {
		return true
	}
	return false
}

func generateRandomInteger(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func isValidTime(startTime, endTime string) bool {
	if startTime == endTime {
		return true
	}
	sTime := strings.Split(startTime, ":")
	eTime := strings.Split(endTime, ":")
	if len(sTime) != 2 || len(eTime) != 2 {
		log.Fatal("invalid time format")
	}
	now := time.Now()
	sh, _ := strconv.Atoi(sTime[0])
	sm, _ := strconv.Atoi(sTime[1])
	eh, _ := strconv.Atoi(eTime[0])
	em, _ := strconv.Atoi(eTime[1])
	// TODO: Error Check
	start := time.Date(now.Year(), now.Month(), now.Day(), sh, sm, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), eh, em, 0, 0, now.Location())

	//fmt.Println(start.Format("2006/01/02 15:04:05"), end.Format("2006/01/02 15:04:05"), now.Format("2006/01/02 15:04:05"))
	return now.After(start) && now.Before(end)
}
