package util

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
@Description : execute a command in windows cmd
*/

func Command(exe string, args []string) (result string, err error) {

	cmd := exec.Command(exe, args...)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		XWarning(fmt.Sprintf("cmd.Start error : %v\n", err))
		return
	}

	var rets []string
	// 从管道中实时获取输出并打印到终端
	for {

		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			break
		}

		rets = append(rets, ConvertByte2String(tmp, "GB18030"))

	}

	result = strings.Join(rets, "\n")

	return result, nil
}
