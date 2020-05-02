package runtime

import (
	"encoding/json"
	"fmt"
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

	intervalTime := widget.NewEntry()
	intervalTime.SetPlaceHolder("30")
	percentage := widget.NewEntry()
	percentage.SetPlaceHolder("17:30")
	startTimeEntry := widget.NewEntry()
	startTimeEntry.SetPlaceHolder("09:00")
	endTimeEntry := widget.NewEntry()
	endTimeEntry.SetPlaceHolder("17:30")

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
			a.Quit()
		}),
		widget.NewButton("Cancel", func() {
			a.Quit()
		}),
	))

	window.ShowAndRun()
}

func onSubmit() {
	fmt.Println("pushed ok")
}

func onCancel() {
	fmt.Println("pushed cancel")
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
