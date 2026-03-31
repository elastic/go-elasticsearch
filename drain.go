package elasticsearch

import (
	"io"
)

type autoDrainingReader struct {
	io.ReadCloser
}

func (a *autoDrainingReader) Close() error {
	_, _ = io.Copy(io.Discard, a.ReadCloser)
	return a.ReadCloser.Close()
}

type autoDrainingReadWriter struct {
	io.ReadWriteCloser
}

func (a *autoDrainingReadWriter) Close() error {
	_, _ = io.Copy(io.Discard, a.ReadWriteCloser)
	return a.ReadWriteCloser.Close()
}

func makeAutoDraining(r io.ReadCloser) io.ReadCloser {
	if rw, ok := r.(io.ReadWriteCloser); ok {
		return &autoDrainingReadWriter{rw}
	}
	return &autoDrainingReader{r}
}
