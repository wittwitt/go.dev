package os1

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

// ### os.mkdir权限问题

// os.mkdir(xxx,mode),,linux调用，会减去umask

// 一般2种方法

// ```golang
// mask := syscall.Umask(0)
// defer syscall.Umask(mask)
// err := os.MkdirAll("/tmp/gotest/", 0777)

// //注意，这个子文件夹要一个一个设置
// err := os.MkdirAll("g10guang/t1/t2", 0777)
// os.Chmod("g10guang", 0777)
// ```

func TestExec1(t *testing.T) {

	fn := func(ctx context.Context) error {
		localPath := ""
		scfsPath := ""
		bwlimit := 20000

		bwLimit := fmt.Sprintf("--bwlimit=%d", bwlimit)

		cmd := exec.Command("rsync", bwLimit, localPath, scfsPath)
		//	cmd := exec.CommandContext(ctx, "rsync", bwLimit, localPath, scfsPath)
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // 进程组
		var stdOut, stdErr bytes.Buffer
		cmd.Stdout = &stdOut
		cmd.Stderr = &stdErr

		waitDone := make(chan struct{})
		go func() {
			select {
			case <-ctx.Done():
				syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) //  负进程号为进程组号
			case <-waitDone:
			}
		}()

		if err := cmd.Start(); err != nil {
			return fmt.Errorf("cmd start: %v", err)
		}
		if err := cmd.Wait(); err != nil {
			stdErrStr := stdErr.String()
			return fmt.Errorf("run: stdErr(%s), %w", stdErrStr, err)
		}
		close(waitDone)

		stdOutStr := stdOut.String()
		if stdOutStr != "" {
			log.Printf("upload: stdOut: %s", stdOutStr)
		}
		stdErrStr := stdErr.String()
		if stdErrStr != "" {
			log.Printf("upload: stdErr: %s", stdErrStr)
		}
		return nil
	}

	fn(context.Background())
}

func TestGPU(t *testing.T) {
	// cmdStr := "lspci | grep  -i vga | grep NVIDIA | wc -l"

	cmd := exec.Command("lspci")
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	cmd.Run()

	t.Log(strings.Contains(stdOut.String(), "NVIDIA"))
}
