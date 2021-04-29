package model

type Ip struct {
	PreviousAddr   string `json:"previous_ip"`
	CurrentAddr    string `json:"current_ip"`
	ChangeTime     string `json:"change_time"`
	NextChangeTime string `json:"next_change_time"`
}
