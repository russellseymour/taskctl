package output

import (
	"fmt"
	"io"
	"regexp"

	"github.com/sirupsen/logrus"

	"github.com/Ensono/taskctl/pkg/task"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var ansiRegexp = regexp.MustCompile(ansi)

type prefixedOutputDecorator struct {
	t *task.Task
	w *SafeWriter
}

func NewPrefixedOutputWriter(t *task.Task, w io.Writer) *prefixedOutputDecorator {
	return &prefixedOutputDecorator{
		t: t,
		w: NewSafeWriter(w),
	}
}

func chunkByteSlice(items []byte, chunkSize int) (chunks [][]byte) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func (d *prefixedOutputDecorator) Write(p []byte) (int, error) {
	return d.w.Write([]byte(fmt.Sprintf("\x1b[18m%s\x1b[0m: %s\r\n", d.t.Name, p)))
}

func (d *prefixedOutputDecorator) WriteHeader() error {
	logrus.Infof("Running task %s...", d.t.Name)
	return nil
}

func (d *prefixedOutputDecorator) WriteFooter() error {
	logrus.Infof("%s finished. Duration %s", d.t.Name, d.t.Duration())
	return nil
}
