package pty

import "io"

func Redirect(src io.Reader, dst io.Writer) (int64, error) { return io.Copy(dst, src) }
