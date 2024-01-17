package main

import (
	"bufio"
	"os"
)

// WriteFile ...
type WriteFile struct {
	f *os.File
	w *bufio.Writer
}

// NewWriteFile ...
func NewWriteFile(fileName string) (*WriteFile, error) {
	wf := &WriteFile{}
	// 覆盖写
	err := wf.Open(fileName)
	if err != nil {
		return nil, err
	}
	return wf, nil
}

// Open ...
func (c *WriteFile) Open(fileName string) error {
	var err error
	c.f, err = os.Create(fileName)
	if err != nil {
		return err
	}
	c.w = bufio.NewWriter(c.f)
	return nil
}

// Close ...
func (c *WriteFile) Close() {
	c.w.Flush()
	c.f.Sync()
	c.f.Close()
}

// WriteString ...
func (c *WriteFile) WriteString(str string) error {
	var err error
	_, err = c.w.WriteString(str)
	if err != nil {
		return err
	}
	c.w.Flush()
	return nil
}
