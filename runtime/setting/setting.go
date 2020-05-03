package setting

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

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

// NewSetting Create Setting struct
func NewSetting() *Setting {
	return &Setting{
		Frequency: Frequency{
			IntervalTime: 30,
			Parcentage:   5,
		},
		ValidTime: ValidTime{
			StartTime: "00:00",
			EndTime:   "00:00",
		},
	}
}

// GetCurrentSetting returns new Setting struct from setting.json.
func GetCurrentSetting() *Setting {
	setting := NewSetting()
	settingFilePath := path.Join(getExecDir(), "/setting.json")
	raw, err := ioutil.ReadFile(settingFilePath)
	if err != nil {
		log.Printf("Warning: Cannot read setting.json. [%s]\n", settingFilePath)
	} else {
		json.Unmarshal(raw, setting)
	}
	return setting
}

func getExecDir() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exe)
}
