package main

import (
	"context"
	"github.com/dreampuf/evernote-sdk-golang/client"
	"testing"
	"time"
)

const (
	EvernoteKey = "abc"
	EvernoteSecret = "abc"
	EvernoteAuthorToken = "abc"
)

func Test_main(t *testing.T) {
	clientCtx, _ := context.WithTimeout(context.Background(), time.Duration(15) * time.Second)

	c := client.NewClient(EvernoteKey, EvernoteSecret, client.SANDBOX)

	us, err := c.GetUserStore()
	if err != nil {
		t.Fatal(err)
	}


	userUrls, err := us.GetUserUrls(clientCtx, EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}

	ns, err := c.GetNoteStoreWithURL(userUrls.GetNoteStoreUrl())
	if err != nil {
		t.Fatal(err)
	}

	note, err := ns.GetDefaultNotebook(clientCtx, EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	if note == nil {
		t.Fatal("Invalid Note")
	}
}
