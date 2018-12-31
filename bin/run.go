package bin

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
)

func goRunFile(file string) error {

	var stdoutBuf, stderrBuf bytes.Buffer

	cmdB := []string{"run", file}
	cmd := exec.Command("go", cmdB...)
	//cmd.Dir = dir

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()

	if err != nil {
		return err
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		return err
	}
	if errStdout != nil || errStderr != nil {
		return errors.New("failed to capture stdout or stderr\n")
	}

	return nil

}
