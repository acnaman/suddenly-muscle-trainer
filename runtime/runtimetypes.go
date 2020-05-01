package runtime

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
