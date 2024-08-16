package pty

import (
	"io"

	"github.com/creack/pty"
)

func Pipe() (io.ReadWriteCloser, io.ReadWriteCloser, error) { return pty.Open() }
