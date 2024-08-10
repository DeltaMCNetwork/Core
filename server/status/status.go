package status

import (
	"net/deltamc/server/component"
)

type Response struct {
	Version     *Version                 `json:"version"`
	Players     *Players                 `json:"players"`
	Description *component.TextComponent `json:"description"`
	Favicon     string                   `json:"favicon"`
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int             `json:"max"`
	Online int             `json:"online"`
	Sample []*PlayerSample `json:"sample"`
}

type PlayerSample struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

func NewResponse() *Response {
	return &Response{}
}

func NewPlayerSample(name string, id string) *PlayerSample {
	return &PlayerSample{Name: name, UUID: id}
}

func (r *Response) WithVersion(proto int, name string) *Response {
	r.Version = &Version{Name: name, Protocol: proto}
	return r
}

func (r *Response) WithInfo(max int, online int) *Response {
	r.Players = &Players{Max: max, Online: online}
	return r
}

func (r *Response) WithSamples(samples []*PlayerSample) *Response {
	r.Players.Sample = append(r.Players.Sample, samples...)
	return r
}

func (r *Response) WithDescription(component *component.TextComponent) *Response {
	r.Description = component
	return r
}

func (r *Response) WithFavicon(favicon string) *Response {
	r.Favicon = favicon
	return r
}
