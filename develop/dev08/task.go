package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	gops "github.com/mitchellh/go-ps"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	invite()
	for sc.Scan() {

		fmt.Println(parseCmd(sc.Text()))
		invite()
	}
}

func parseCmd(arg string) string {
	cmds := strings.Split(arg, "|")
	if len(cmds) > 1 {
		var cmdArg []string = strings.Split(cmds[0], " ")
		var result string
		if len(cmdArg) == 1 {
			result = executeCmd(cmdArg[0], "")
		}
		for index, val := range cmds {
			if index == 0 {
				continue
			}
			cmdArg = strings.Split(val, " ")
			result = executeCmd(cmdArg[0], result)
		}
		return result
	} else if len(cmds) == 1 {
		var cmdArg []string = strings.Split(arg, " ")
		if len(cmdArg) == 2 {
			return executeCmd(cmdArg[0], cmdArg[1])
		} else if len(cmdArg) == 1 {
			return executeCmd(cmdArg[0], "")
		}
		return "Many arguments"
	} else {
		return "Wrong arguments"
	}
}

func executeCmd(cmd string, arg string) string {
	if cmd == "cd" {
		return cd(arg)
	} else if cmd == "pwd" {
		return pwd(arg)
	} else if cmd == "echo" {
		return echo(arg)
	} else if cmd == "kill" {
		return kill(arg)
	} else if cmd == "ps" {
		return ps(arg)
	} else {
		return ""
	}
}

func invite() {
	host, _ := os.Hostname()
	path := pwd("")
	fmt.Printf("%s@%s: ", host, path)
}

func cd(arg string) string {
	err := os.Chdir(arg)
	if err != nil {
		return err.Error()
	}
	return ""
}

func pwd(arg string) string {
	path, _ := os.Getwd()
	return path
}

func echo(arg string) string {
	return arg
}

func kill(arg string) string {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return "id process is wrong"
	}

	proc, err := os.FindProcess(id)
	if err != nil {
		return "Process not exist"
	}

	errKill := proc.Kill()
	if errKill != errKill {
		return "Error kill process"
	}

	return arg
}

func ps(arg string) string {
	procList, _ := gops.Processes()
	var head string = "PID PPID EXECUTABLE\n"
	var procListString []string = make([]string, 0)
	for _, v := range procList {
		var pid string = strconv.Itoa(v.Pid())
		var ppid string = strconv.Itoa(v.PPid())
		var ex string = v.Executable()
		procListString = append(procListString, pid+" "+ppid+" "+ex)
	}
	var result string = head + strings.Join(procListString, "\n")
	return result
}
