package model

import (
	"PublicClipboard/gobalConfig"
	"gorm.io/gorm"
)

var clipboard Clipboard
var Logs []Log

type Clipboard struct {
	gorm.Model
	Msg string `json:"msg"`
}
type Log struct {
	Id   int    `json:"id" gorm:"autoIncrement"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}

func UpdClipboard(msg string) (bool, error) {
	clipboard.Msg = msg
	return true, nil
}
func GetClipboard() Clipboard {
	return clipboard
}
func AddLog(log Log) {
	log.Id = len(Logs)
	tLogs := []Log{log}
	Logs = append(tLogs, Logs...)
	if len(Logs) > gobalConfig.HistoryNum {
		Logs = Logs[:gobalConfig.HistoryNum-1]
	}
	return
}

func ListLog() []Log {
	return Logs
}
