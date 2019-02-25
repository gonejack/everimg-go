package log

const (
	DEBUG LEVEL = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelTexts = map[LEVEL]string{
	DEBUG: "DEBUG",
	INFO: "INFO",
	WARN: "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

type LEVEL int

func (l LEVEL) String() string {
	return levelTexts[l]
}