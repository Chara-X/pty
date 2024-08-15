package pty

import (
	"io"
	"os"
)

func Open() (pty, tty io.ReadWriteCloser) {
	var pr, tw, _ = os.Pipe()
	var tr, pw, _ = os.Pipe()
	pty = &file{pr, pw}
	tty = &file{tr, tw}
	return
}

type file struct {
	io.Reader
	io.WriteCloser
}
