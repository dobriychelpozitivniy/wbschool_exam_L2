package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то) +
- pwd - показать путь до текущего каталога +
- echo <args> - вывод аргумента в STDOUT +
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример) +
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат* +


Так же требуется поддерживать функционал fork/exec-команд
*/

type ShellFunc func(args []string) error

type ShellCMD map[string]ShellFunc

func main() {
	shell()
}

func shell() {
	in := os.Stdin

	sh := initShellCMDS()

	sc := bufio.NewScanner(in)
	for sc.Scan() {
		args := parseArgs(sc.Text())

		if f, ok := sh[args[0]]; ok {
			err := f(args[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Println(fmt.Sprintf("Комманда %s не распознана", sc.Text()))
	}
}

func initShellCMDS() ShellCMD {
	var shell ShellCMD = make(ShellCMD)

	shell["cd"] = cd
	shell["pwd"] = pwd
	shell["echo"] = echo
	shell["kill"] = kill
	shell["ps"] = ps
	shell["ls"] = ls

	return shell
}

func ls(args []string) error {
	cArgs := []string{
		"ls",
	}

	cArgs = append(cArgs, args...)

	cmd := exec.Command("powershell", cArgs...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func cd(args []string) error {
	// 'cd' to home dir with empty path not yet supported.
	if len(args) < 1 {
		return errors.New("path required")
	}
	// Change the directory and return the error.
	return os.Chdir(args[0])
}

func pwd(args []string) error {
	cArgs := []string{
		"pwd",
	}

	cArgs = append(cArgs, args...)

	cmd := exec.Command("powershell", cArgs...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func echo(args []string) error {
	cArgs := []string{
		"echo",
	}

	cArgs = append(cArgs, args...)

	cmd := exec.Command("powershell", cArgs...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func kill(args []string) error {
	cArgs := []string{
		"Stop-Process",
	}

	cArgs = append(cArgs, args...)

	cmd := exec.Command("powershell", cArgs...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func ps(args []string) error {
	cArgs := []string{
		"ps",
	}

	cArgs = append(cArgs, args...)

	out, err := exec.Command("powershell", cArgs...).Output()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println(string(out))

	return nil
}

func parseArgs(s string) []string {
	args := strings.Split(s, " ")

	return args
}
