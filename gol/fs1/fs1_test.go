package fs1

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_t1(t *testing.T) {
	Statfs("/home/abc/asd")
	Statfs("/home/abc/asd/abcd")
	Statfs("/home/abc/asd/abce")
}

// 获取path所在的硬盘分区总容量、已使用量
func Statfs(path string) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Capacity: ", int64(stat.Blocks)*int64(stat.Bsize)/1024/1024/1024)
	fmt.Println("FSAvailable:   ", int64(stat.Bavail)*int64(stat.Bsize)/1024/1024/1024)
}

func Test_t2(t *testing.T) {
	require.Equal(t, filepath.Base("/abc/asd"), "asd")
	require.Equal(t, filepath.Base("/abc/asd.exe"), "asd.exe")
	require.Equal(t, filepath.Base("/abc/asd.exe.abc"), "asd.exe.abc")
}

func Test_tmp(t *testing.T) {
	dir, err := os.MkdirTemp(os.TempDir(), "go.dev.*")
	require.NoError(t, err)
	k2p := filepath.Join(dir, "2k.txt")
	t.Log(k2p)
	k2f, err := os.OpenFile(k2p, os.O_EXCL|os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	require.NoError(t, err)

	// 创建一个大小为2K的字节数组
	data := make([]byte, 2048)
	rand.Read(data)

	_, err = k2f.Write(data)
	require.NoError(t, err)
	k2f.Close()
}

func Test_stat(t *testing.T) {
	dir, err := os.MkdirTemp(os.TempDir(), "go.dev.*")
	require.NoError(t, err)

	statInfo, err := os.Stat(dir)
	require.NoError(t, err)

	t.Log(statInfo.IsDir())

	_, err = os.Stat(filepath.Join("some"))
	require.Equal(t, true, errors.Is(err, os.ErrNotExist))

}
