package main

import (
	"bytes"
	"errors"
	"io"
	"os"
)

// protocol enum
const (
	Default = Local
	Local   = "local"
)

var (
	streambuilder = map[string]func(s *Stream) (r io.WriteCloser, err error){
		"local": localbuilder,
	}
)

// stream implement
func (s *Stream) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}
	index := bytes.Index(text, []byte("://"))
	if index < 0 {
		s.Protocol = Default
		s.Address = string(text)
		return nil
	} else if index == 0 {
		return errors.New("error format at ://")
	}
	s.Protocol = string(text[:index])
	text = text[index+3:]
	index = bytes.Index(text, []byte("||"))
	if index < 0 {
		s.Address = string(text)
		return nil
	} else if index == 0 {
		return errors.New("error format at ||")
	}
	s.Address = string(text[:index])
	s.Token = string(text[index+2:])
	return nil
}

func buildLogger(s *Stream) (r io.WriteCloser, err error) {
	if len(s.Protocol) == 0 {
		s.Protocol = Default
	}
	if rbuilder, ok := streambuilder[s.Protocol]; !ok {
		err = errors.New("not support protocol " + s.Protocol)
		return
	} else {
		r, err = rbuilder(s)
	}
	return
}

func localbuilder(s *Stream) (r io.WriteCloser, err error) {
	if len(s.Address) == 0 {
		r = os.Stdout
		return
	}
	r, err = os.OpenFile(s.Address, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return
}
