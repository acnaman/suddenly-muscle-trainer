package main

import (
	"errors"
	"os"
	"regexp"
	"strconv"

	"fyne.io/fyne"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/acnaman/suddenly-muscle-trainer/runtime/setting"
)

func main() {
	ShowSettingDialog()
}

func ShowSettingDialog() {
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
			var err error
			setting.IntervalTime, err = checkIntervalTime(intervalTime.Text)
			if err != nil {
				showErrorMessage(a, err)
				return
			}
			setting.Parcentage, err = checkPercentage(percentage.Text)
			if err != nil {
				showErrorMessage(a, err)
				return
			}
			setting.StartTime, err = checkTimeFormat(startTimeEntry.Text)
			if err != nil {
				showErrorMessage(a, err)
				return
			}
			setting.EndTime, err = checkTimeFormat(endTimeEntry.Text)
			if err != nil {
				showErrorMessage(a, err)
				return
			}
			onSubmit(setting)
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	window.ShowAndRun()
}

func checkIntervalTime(text string) (int, error) {
	i, err := strconv.Atoi(text)
	const maxminute = 1440 // 60min * 24hour

	if err != nil {
		return 0, errors.New("Interval Time shoud be an integer from 0 to 1440")
	}
	if i < 0 || maxminute < i {
		return 0, errors.New("Parcentage Time shoud be an integer from 0 to 1440")
	}

	return i, nil
}

func checkPercentage(text string) (int, error) {
	i, err := strconv.Atoi(text)

	if err != nil {
		return 0, errors.New("Parcentage shoud be an integer from 0 to 100")
	}
	if i < 0 || 100 < i {
		return 0, errors.New("Parcentage Time shoud be an integer from 0 to 100")
	}

	return i, nil
}

var re = regexp.MustCompile(`^(\d\d):(\d\d)$`)

func checkTimeFormat(text string) (string, error) {
	b := re.MatchString(text)
	errMsg := "Time should be in the format 'hh:mm' (from 00:00 to 23:59)"
	if !b {
		return "", errors.New(errMsg)
	}

	s := re.FindAllStringSubmatch(text, -1)
	hour := s[0][1]
	ihour, err := strconv.Atoi(hour)
	if err != nil {
		return "", errors.New(errMsg)
	}
	if ihour < 0 || 23 < ihour {
		return "", errors.New(errMsg)
	}
	min := s[0][2]
	imin, err := strconv.Atoi(min)
	if err != nil {
		return "", errors.New(errMsg)
	}
	if imin < 0 || 59 < imin {
		return "", errors.New(errMsg)
	}

	return text, nil
}

func showErrorMessage(a fyne.App, err error) {
	w := a.NewWindow("Error!")
	w.SetContent(widget.NewVBox(
		widget.NewLabel(err.Error()),
		widget.NewButton("OK", func() {
			w.Hide()
		}),
	))
	w.RequestFocus()

	w.Show()

}

func onSubmit(s *setting.Setting) {
	setting.SaveToFile(s)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
