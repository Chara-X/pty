# pty

```go
func ExamplePty() {
	var c = exec.Command("/bin/bash")
	var p, t = pt.Open()
	defer p.Close()
	defer t.Close()
	c.Stdin = t
	c.Stdout = t
	c.Stderr = t
	c.Start()
	defer c.Process.Kill()
	go io.Copy(os.Stdout, p)
	io.Copy(p, os.Stdin)
}
```

## Reference

[Build A Simple Terminal Emulator In 100 Lines of Golang](https://ishuah.com/2021/03/10/build-a-terminal-emulator-in-100-lines-of-go/)
