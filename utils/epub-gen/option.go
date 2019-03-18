package epub_gen

import "time"

type Book struct {
	Title       string
	Description string
	Publisher   string
	Author      []string
	Date        time.Time
	Lang        string
	Fonts       []string
	TocTitle    string
	Cover       string
	Content     []*Content
	Version     int
}

type Content struct {
	Title  string
	Author []string
	Date   string
	Link   string
	Data   string

	Filename       string
	Href           string
	ExcludeFromToc bool
	BeforeToc      bool

	id  string
	dir string
}

type Control struct {
	Looking struct {
		AppendChapterTitles bool
	}
	HttpDL struct {
		Timeout    time.Time
		RetryTimes int
		Headers    map[string]string
	}
	Output struct {
		Path string
	}
	Debug struct {

	}
}