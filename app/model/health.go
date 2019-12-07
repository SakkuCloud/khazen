package model

import (
	"khazen/config"
	"time"
)

type Health struct {
	Status     bool   `json:"status"`
	Uptime     string `json:"uptime"`
	ServerTime string `json:"server_time"`
}

func (h *Health) SetUptime() {
	h.Uptime = time.Since(config.StartTime).String()
}

func (h *Health) SetServerTime() {
	h.ServerTime = time.Now().Format(time.RFC850)
}
