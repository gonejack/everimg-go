package noteService

import (
	"context"
	"everimg-go/app/log"
	"github.com/dreampuf/evernote-sdk-golang/client"
	"github.com/dreampuf/evernote-sdk-golang/edam"
	"github.com/spf13/viper"
	"time"
)

type noteService struct {
	logger log.Logger
	conf *viper.Viper

	userStore edam.UserStore
}

func (s *noteService) GetRecentUpdateNotes() (notes []edam.Note) {
	metaList, err := s.getRecentUpdateNoteMetas()

	if err == nil {
		for _, meta := range metaList.GetNotes() {
			_ = s.getNote(meta)
		}
	}

	panic("todo")
}

func (s *noteService) getRecentUpdateNoteMetas() (metaList edam.NotesMetadataList, err error) {
	panic("todo")
}

func (s *noteService) Start() {
	clientCtx, _ := context.WithTimeout(context.Background(), time.Duration(15) * time.Second)
	c := client.NewClient("", "", client.SANDBOX)
	us, err := c.GetUserStore()
	if err != nil {
		panic(err)
	}
	userUrls, err := us.GetUserUrls(clientCtx, s.conf.GetString("token"))
	if err != nil {
		panic(err)
	}
	ns, err := c.GetNoteStoreWithURL(userUrls.GetNoteStoreUrl())
	if err != nil {
		panic(err)
	}
	note, err := ns.GetDefaultNotebook(clientCtx, s.conf.GetString("token"))
	if err != nil {
		panic(err)
	}
	if note == nil {
		panic(err)
	}
}

func (*noteService) Stop() {
	panic("implement me")
}

func (s *noteService) getNote(metadata *edam.NoteMetadata) (note edam.Note) {
	panic("todo")
}

func New(name string, conf string) *noteService {
	return new(noteService)
}

func NewDefault() *noteService {
	return new(noteService)
}