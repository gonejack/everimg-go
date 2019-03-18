package kit

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func InArray(needle string, stack []string) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}

	return false
}

func SecondsToText(seconds int) (text string) {
	units := []string{"天", "小时", "分钟", "秒"}
	numbers := []int{seconds / 86400, seconds / 3600 % 24, seconds / 60 % 60, seconds % 60}

	for idx, unit := range units {
		if numbers[idx] > 0 {
			text += strconv.Itoa(numbers[idx]) + unit
		}
	}

	if text == "" {
		text = "0秒"
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

func Uuid() string {
	bs := make([]byte, 16)

	if _, err := rand.Read(bs); err == nil {
		bs[8] = (bs[8] | 0x80) & 0xBF
		bs[6] = (bs[6] | 0x40) & 0x4F

		return fmt.Sprintf("%x-%x-%x-%x-%x", bs[0:4], bs[4:6], bs[6:8], bs[8:10], bs[10:])
	} else {
		panic(err)
	}
}

func RemoveDiacritics(str string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	s, _, _ := transform.String(t, str)

	return s
}