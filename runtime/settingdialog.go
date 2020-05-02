package runtime

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"fyne.io/fyne/widget"

	"fyne.io/fyne/app"
)

func ShowSettingDialog() {
	settingFilePath := path.Join(getExecDir(), "/setting.json")
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
	if exists(settingFilePath) {
		raw, err := ioutil.ReadFile(settingFilePath)
		if err != nil {
			log.Printf("Warning: Cannot read setting.json. [%s]\n", settingFilePath)
		} else {
			json.Unmarshal(raw, setting)
		}
	}

	/*f, err := os.OpenFile(settingFilePath, os.O_RDWR|os.O_CREATE|O_, 0755)
	if err != nil {
		log.Fatal(err)
	}*/

	a := app.New()
	window := a.NewWindow("Setting")
	window.SetContent(widget.NewVBox(
		widget.NewLabel("Muscle Trainer Setting"),
		widget.NewButton("Save", func() {
			a.Quit()
		}),
		widget.NewButton("Cancel", func() {
			a.Quit()
		}),
	))

	window.ShowAndRun()
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
