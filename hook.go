package simplelog

import (
	"github.com/tanzy2018/simplelog/encode"
)

// HookFunc ...
type HookFunc func() encode.Meta

// IHook ...
type IHook interface {
	Add(hfs ...HookFunc)
	Hooks() []encode.Meta
}

type hook struct {
	hfs []HookFunc
}

func (h *hook) Add(hfs ...HookFunc) {
	if len(h.hfs) == 0 {
		h.hfs = make([]HookFunc, 0, 1)
	}
	for _, hf := range hfs {
		h.hfs = append(h.hfs, hf)
	}
}

func (h *hook) Hooks() []encode.Meta {
	if h == nil {
		return nil
	}
	md := make([]encode.Meta, 0, len(h.hfs))
	for _, hf := range h.hfs {
		md = append(md, hf())
	}
	return md
}
