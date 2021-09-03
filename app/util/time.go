package util

import "time"

var now = time.Now()
var fmtd = "2006/01/02"
var fmtf = "2006/01/02 15:04:05"

func GetToday () string {
  return now.Format(fmtd)
}

func GetYesterday() string {
  yt := time.Now().Add(-time.Duration(24) * time.Hour)
  return yt.Format(fmtd)
}

func GetDaily() (string, string) {
  tml := now.Add(time.Duration(24) * time.Hour)
  return now.Format(fmtd), tml.Format(fmtd)
}

