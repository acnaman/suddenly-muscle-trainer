package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
)

type program struct {
	*Setting
	exit chan struct{}
}

type Setting struct {
	Frequency `json:"frequency"`
	ValidTime `json:"validTime"`
}

type Frequency struct {
	IntervalTime int `json:"intervalTime"`
	Parcentage   int `json:"parcentage"`
}

type ValidTime struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
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

	// load setting file
	mtlogger.WriteStartLog()
	setting := &Setting{
		Frequency: Frequency{
			IntervalTime: 30,
			Parcentage:   5,
		},
		ValidTime: ValidTime{
			StartTime: "00:00",
			EndTime:   "00:00",
		},
	}
	settingFilePath := path.Join(getExecDir(), "/setting.json")
	raw, err := ioutil.ReadFile(settingFilePath)
	if err != nil {
		log.Printf("Warning: Cannot read setting.json. [%s]\n", settingFilePath)
	} else {
		json.Unmarshal(raw, setting)
	}
	p.Setting = setting

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
	mtlogger = NewLogger(path.Join(getExecDir(), "muscletrainer.log"))

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

func getExecDir() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exe)
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
