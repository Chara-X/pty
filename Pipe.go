package pty

import (
	"io"
	"os"
)

func Pipe() (io.ReadWriteCloser, io.ReadWriteCloser, error) {
	// return pty.Open()
	var pr, tw, _ = os.Pipe()
	var tr, pw, _ = os.Pipe()
	return &file{Reader: pr, WriteCloser: pw}, &file{Reader: tr, WriteCloser: tw}, nil
}

type file struct {
	io.Reader
	io.WriteCloser
}
