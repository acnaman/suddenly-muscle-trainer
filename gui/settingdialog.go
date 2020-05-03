package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/acnaman/suddenly-muscle-trainer/runtime/setting"
)

func main() {
	ShowSettingDialog()
}

func ShowSettingDialog() {
	settingFilePath := path.Join(getExecDir(), "/setting.json")
	setting := setting.GetCurrentSetting()

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

func onSubmit(s *setting.Setting) {
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
