package model

import "time"

type Agent struct {
	Crc          string `json:"crc"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	SessionKey   []byte `json:"session_key"`
	Listener     string `json:"listener"`
	Async        bool   `json:"async"`
	ExternalIP   string `json:"external_ip"`
	InternalIP   string `json:"internal_ip"`
	GmtOffset    int    `json:"gmt_offset"`
	Sleep        uint   `json:"sleep"`
	Jitter       uint   `json:"jitter"`
	Pid          string `json:"pid"`
	Tid          string `json:"tid"`
	Arch         string `json:"arch"`
	Elevated     bool   `json:"elevated"`
	Process      string `json:"process"`
	Os           int    `json:"os"`
	OsDesc       string `json:"os_desc"`
	Domain       string `json:"domain"`
	Computer     string `json:"computer"`
	Username     string `json:"username"`
	Impersonated string `json:"impersonated"`
	OemCP        int    `json:"oemcp"`
	ACP          int    `json:"acp"`
	LastTick     int    `json:"last_tick"`
	KillDate     int    `json:"killdate"`
	WorkingTime  int    `json:"workingtime"`
}
