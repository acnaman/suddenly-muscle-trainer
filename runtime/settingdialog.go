package runtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

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

	a := app.New()
	window := a.NewWindow("Muscle Trainer Setting")

	intervalTime := widget.NewEntry()
	intervalTime.SetPlaceHolder("30")
	intervalTime.SetText(strconv.Itoa(setting.IntervalTime))

	percentage := widget.NewEntry()
	percentage.SetPlaceHolder("17:30")
	percentage.SetText(strconv.Itoa(setting.Parcentage))

	startTimeEntry := widget.NewEntry()
	startTimeEntry.SetPlaceHolder("09:00")
	startTimeEntry.SetText(setting.StartTime)

	endTimeEntry := widget.NewEntry()
	endTimeEntry.SetPlaceHolder("17:30")
	endTimeEntry.SetText(setting.EndTime)

	window.SetContent(widget.NewVBox(
		widget.NewLabel("Muscle Trainer Setting"),
		&widget.Form{
			Items: []*widget.FormItem{
				{"Interval Time (min)", intervalTime},
				{"Percentage", percentage},
				{"Start Time", startTimeEntry},
				{"End Time", endTimeEntry},
			},
		},
		widget.NewButton("Save", func() {
			setting.IntervalTime, _ = strconv.Atoi(intervalTime.Text)
			setting.Parcentage, _ = strconv.Atoi(percentage.Text)
			setting.StartTime = startTimeEntry.Text
			setting.EndTime = endTimeEntry.Text

			onSubmit(setting)
			a.Quit()
		}),
		widget.NewButton("Cancel", func() {
			a.Quit()
		}),
	))

	window.ShowAndRun()
}

func onSubmit(s *Setting) {
	fmt.Println("pushed ok")

	settingFilePath := path.Join(getExecDir(), "/setting.json")
	f, err := os.OpenFile(settingFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	f.Write(b)

}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
