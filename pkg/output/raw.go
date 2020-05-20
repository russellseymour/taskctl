package output

import (
	"fmt"
	"io"
)

type rawOutputDecorator struct {
	w io.Writer
}

func newRawOutputWriter(w io.Writer) *rawOutputDecorator {
	return &rawOutputDecorator{w: w}
}

func (d *rawOutputDecorator) WriteHeader() error {
	return nil
}

func (d *rawOutputDecorator) Write(b []byte) (int, error) {
	return d.w.Write(b)
}

func (d *rawOutputDecorator) WriteFooter() error {
	_, err := fmt.Fprint(d.w, "\r\n")
	return err
}
