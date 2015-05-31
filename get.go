package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/daviddengcn/go-colortext"
)

type Color int

const (
	None Color = Color(ct.None)
	Red  Color = Color(ct.Red)
	Blue Color = Color(ct.Blue)
)

var (
	stdout = os.Stdout
	stderr = os.Stderr
	stdin  = os.Stdin
)

func (p *Package) Get() error {
	args := []string{"go", "get"}
	if !p.Noupdate {
		args = append(args, "-u")
	}
	if p.Noinstall {
		args = append(args, "-d")
	}
	args = append(args, p.Name)
	return run(args, Blue)
}

func run(args []string, c Color) error {
	if len(args) == 0 {
		return errors.New("too few arguments")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = stdout
	//cmd.Stderr = stderr
	cmd.Stdin = stdin
	ct.ChangeColor(ct.Color(c), true, ct.None, false)
	err := cmd.Run()
	ct.ResetColor()
	return err
}
