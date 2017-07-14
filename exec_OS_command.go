package comm

import (
	//	"fmt"

	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

//执行Linux系统命令
// //this method can get the err information about the os command stdout returns
// func ExecOScmd(cmdName string) ([]byte, error) {
// 	//c := "ps -ef | grep -v \"grep\" | grep \"docker\""
// 	cmdArgs := strings.Fields(cmdName)

// 	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
// 	//cmd := exec.Command("/bin/sh", "-c", cmdName)
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		//	fmt.Println("StdoutPipe: " + err.Error())
// 		return []byte(""), err
// 	}

// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		//	fmt.Println("StderrPipe: ", err.Error())
// 		return []byte(""), err
// 	}
// 	defer stdout.Close()
// 	if err := cmd.Start(); err != nil {
// 		//fmt.Println("Start: ", err.Error())
// 		return []byte(""), err
// 	}

// 	bytesErr, err := ioutil.ReadAll(stderr)
// 	if err != nil {
// 		//fmt.Println("ReadAll stderr: ", err.Error())
// 		return bytesErr, err
// 	}

// 	if len(bytesErr) != 0 {
// 		//	fmt.Printf("stderr is not nil: %s", bytesErr)
// 		return []byte(bytesErr), err
// 	}

// 	bytes, err := ioutil.ReadAll(stdout)
// 	if err != nil {
// 		//fmt.Println("ReadAll stdout: ", err.Error())
// 		return bytes, err
// 	}

// 	if err := cmd.Wait(); err != nil {
// 		//fmt.Println("Wait: ", err.Error())
// 		return []byte(""), err
// 	}

// 	//fmt.Printf("stdout: %s", bytes)

// 	return bytes, err

// }

// func ExecOScmd2(cmdName string) ([]byte, error) {
// 	log.Println("cmdnames:", cmdName)
// 	cmd := exec.Command("/usr/bin/sh", "-c", cmdName)
// 	stdoutStderr, err := cmd.Output()
// 	if err != nil {
// 		return []byte(""), err
// 	}

// 	//fmt.Printf("%s\n", stdoutStderr)

// 	// // 获取输出对象，可以从该对象中读取输出结果
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		return []byte(""), err
// 	}
// 	// 保证关闭输出流
// 	defer stdout.Close()
// 	// 运行命令
// 	if err := cmd.Start(); err != nil {
// 		return []byte(""), err
// 	}
// 	// 读取输出结果
// 	_, err1 := ioutil.ReadAll(stdout)
// 	if err != nil {
// 		return []byte(""), err1
// 	}
// 	return stdoutStderr, err
// }

// func ExecOScmd3(cmdName string) bool {
// 	cmdArgs := strings.Fields(cmdName)

// 	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)

// 	//显示运行的命令
// 	log.Println("cccccc->", cmd.Args)
// 	stdoutStderr, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s\n", stdoutStderr)

// 	stdout, err := cmd.StdoutPipe()

// 	if err != nil {
// 		log.Println(err)
// 		//return false
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		log.Printf("Command start with error: %v", err)
// 	}
// 	reader := bufio.NewReader(stdout)

// 	//实时循环读取输出流中的一行内容
// 	for {
// 		line, err2 := reader.ReadString('\n')
// 		if err2 != nil || io.EOF == err2 {
// 			break
// 		}
// 		log.Println(line)
// 	}

// 	err = cmd.Wait()
// 	if err != nil {
// 		log.Printf("Command finished with error: %v", err)
// 	}
// 	return true
// }

// func ExecOScmd4(cmdName string) (string, error) {
// 	res, err := exec.Command("/usr/bin/sh", "-c", cmdName).CombinedOutput()
// 	return string(res), err
// }

// func ExecOScmd5(cmdName string) (string, error) {
// 	cmdArgs := strings.Fields(cmdName)

// 	//cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)

// 	res, err := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...).CombinedOutput()
// 	return string(res), err

// 	//return string(""), err
// }

// func ExecOScmd6(cmdName string) (string, error) {
// 	cmdArgs := strings.Fields(cmdName)

// 	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := cmd.Start(); err != nil {
// 		log.Fatal(err)
// 	}

// 	slurp, _ := ioutil.ReadAll(stderr)
// 	log.Printf("%s\n", slurp)

// 	if err := cmd.Wait(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return string(slurp), err
// }

// func ExecOSCmd(cmdName string) (res []byte, errs string) {
// 	var stdoutBuf, stderrBuf bytes.Buffer
// 	cmdArgs := strings.Fields(cmdName)
// 	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
// 	stdoutIn, _ := cmd.StdoutPipe()
// 	stderrIn, _ := cmd.StderrPipe()

// 	var errStdout, errStderr error
// 	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
// 	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
// 	err := cmd.Start()
// 	if err != nil {
// 		//log.Fatalf("cmd.Start() failed with '%s'\n", err)
// 		return nil, err.Error()
// 	}

// 	go func() {
// 		_, errStdout = io.Copy(stdout, stdoutIn)
// 	}()

// 	go func() {
// 		_, errStderr = io.Copy(stderr, stderrIn)
// 	}()

// 	err = cmd.Wait()
// 	if err != nil {
// 		//log.Fatalf("cmd.Run() failed with %s\n", err)
// 		return nil, err.Error()
// 	}
// 	if errStdout != nil || errStderr != nil {
// 		//log.Fatal("failed to capture stdout or stderr\n")
// 		return nil, errors.New("failed to capture stdout or stderr").Error()
// 	}
// 	outStr, errStr := stdoutBuf.Bytes(), stderrBuf.Bytes()
// 	//	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
// 	return outStr, string(errStr)
// }

func ExecOSCmdNoReturn(cmdName string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmdArgs := strings.Fields(cmdName)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Println("cmd.Start() failed with '%s'\n", err)
		return
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Println("cmd.Run() failed with %s\n", err)
		return
	}
	if errStdout != nil || errStderr != nil {
		log.Println("failed to capture stdout or stderr\n")
		return
	}
	//	outStr, errStr := stdoutBuf.Bytes(), stderrBuf.Bytes()
	//log.Println("out=",string(outStr))
	//log.Println("err=",string(errStr))
	//log.Println("\nout:\n%s\nerr:\n%s\n", string(outStr), string(errStr))
	return
}

func ExecOSCmdForBash(cmdName string) (res []byte, errs error) {
	var stdoutBuf, stderrBuf bytes.Buffer
	//cmdArgs := strings.Fields(cmdName)
	cmd := exec.Command("/usr/bin/sh", "-c", cmdName)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		//log.Fatalf("cmd.Start() failed with '%s'\n", err)
		return nil, err
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		//log.Fatalf("cmd.Run() failed with %s\n", err)
		return nil, err
	}
	if errStdout != nil || errStderr != nil {
		//log.Fatal("failed to capture stdout or stderr\n")
		return nil, errors.New("failed to capture stdout or stderr")
	}
	outStr, errStr := stdoutBuf.Bytes(), stderrBuf.Bytes()
	//	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	if len(errStr) > 0 {
		errs = errors.New(string(errStr))
	}
	return outStr, errs
}
