package kit

import (
	"strconv"
	"strings"
	"time"
)

func InArray(needle string, stack []string) bool  {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}

	return false
}

func SecondsToText(seconds int) (text string)  {
	units := []string{"天", "小时", "分钟", "秒"}
	numbers := []int{seconds / 86400, seconds / 3600 % 24, seconds / 60 % 60, seconds % 60}

	for idx, unit := range units {
		if numbers[idx] > 0 {
			text += strconv.Itoa(numbers[idx]) + unit
		}
	}

	if text == "" {
		text = "1秒"
	}

	return
}

func TimestampToText(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func MacroReplace(tpl string, info map[string]string) string {
	var oldNews []string

	for search, replace := range info {
		oldNews = append(oldNews, search, replace)
	}

	return strings.NewReplacer(oldNews...).Replace(tpl)
}