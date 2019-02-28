package writer

type Interface interface {
	WriteString(s string)
	WriteBytes(bs []byte)
	Start()
	Stop()
}