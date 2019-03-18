package epub_gen

import (
	"everimg-go/utils/kit"
)

type creator struct {
	book Book
	ctl Control

	wd string
}

func (c *creator) create() error {
	panic("")
}
func (c *creator) render() error {
	panic("")
}
func (c *creator) makeCover() error {
	panic("")
}
func (c *creator) downloadAllImage() error {
	panic("")
}
func (c *creator) GenEpub() error {
	panic("")
}

func Epub(book Book, ctl Control) (err error) {
	c := &creator{}

	c.book = book
	c.ctl = ctl
	c.wd = kit.Uuid()

	return c.create()
}
