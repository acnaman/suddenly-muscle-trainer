package setting

import (
	"bytes"
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
	VideoList `json:"videoList"`
}

type Frequency struct {
	IntervalTime int `json:"intervalTime"`
	Percentage   int `json:"percentage"`
}

type ValidTime struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type VideoList []string

// NewSetting Create Setting struct
func NewSetting() *Setting {
	return &Setting{
		Frequency: Frequency{
			IntervalTime: 30,
			Percentage:   5,
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

func SaveToFile(s *Setting) {
	settingFilePath := path.Join(getExecDir(), "/setting.json")
	f, err := os.OpenFile(settingFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	out := new(bytes.Buffer)
	json.Indent(out, b, "", "    ")

	f.Write([]byte(out.String()))
}
