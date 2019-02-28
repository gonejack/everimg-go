package request

import "encoding/json"

type Request struct {
	Name string `json:"name"`
}

func (r *Request) String() string {
	bs, _ := r.IntoJSON()

	return string(bs)
}

func (r *Request) FromJSON(bs []byte) (*Request, error)  {
	return r, json.Unmarshal(bs, r)
}

func (r *Request) IntoJSON() ([]byte, error) {
	return json.Marshal(r)
}

func New() *Request {
	return new(Request)
}