package archive1

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"syscall"
	"time"
)

const tarModTime = "2023-04-01 07:04:05" // UTC time

// Dir2tar tar dir to tar file
func dir2tar(ctx context.Context, srcDir string, tarPath string, isCompress bool) (err error) {

	// 注意：
	// tar会打包文件的时间戳、所有者、反问权限等信息造成tar md5不同
	// tar最小10k？？

	// 打包原则：
	// 1. 不要外带文件夹，直接打包文件
	// 2. 所有文件权限0644
	// 3. 打包命令去各种附带信息，排序按字母，统一文件一个时间（UTC)： 2023-04-01 07:04:05

	// 打包命令：
	// tar --no-acls --no-selinux --no-xattrs --sort=name --owner=0 --group=0 --numeric-owner --mtime='2023-04-01 15:04:05' -cvf c1.tar  ./c1

	// 解包命令：
	// tar --no-same-owner --no-same-permissions -xvf c1.tar

	defer func() {
		if rev := recover(); rev != nil {
			err = fmt.Errorf("dir2tar recover: %v, %v", rev, err)
			return
		}
	}()

	filelist := []string{}

	if err := os.Chmod(srcDir, 0755); err != nil {
		return fmt.Errorf("0755: %s,%v", srcDir, err)
	}

	filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		filelist = append(filelist, d.Name())
		return os.Chmod(path, 0644)
	})

	sort.Slice(filelist, func(i, j int) bool {
		return filelist[i] < filelist[j]
	})

	tarArg := "-cvPf"
	if isCompress {
		tarArg = "-czvPf"
	}

	args := []string{
		"--blocking-factor=8", // 不懂
		"--no-acls",
		"--no-selinux",
		"--no-xattrs",
		"--sort=name",
		"--owner=0",
		"--group=0",
		"--numeric-owner",
		"--mtime=2023-04-01 15:04:05",
		"--format=ustar",
		tarArg,
		tarPath,
		"-C", srcDir}

	args = append(args, filelist...)

	cmd := exec.Command("tar", args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	waitDone := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		case <-waitDone:
		}
	}()

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("cmd.Start: %v", err)
	}
	if err := cmd.Wait(); err != nil {
		stdErrStr := stdErr.String()
		return fmt.Errorf("cmd.Wait: stdErr(%s), %v", stdErrStr, err)
	}
	close(waitDone)

	stdOutStr := stdOut.String()
	if stdOutStr != "" {
		log.Printf("stdOut: %s", stdOutStr)
	}
	stdErrStr := stdErr.String()
	if stdErrStr != "" {
		log.Printf("stdErr: %s", stdErrStr)
	}
	return nil
}

func dir2tar2(ctx context.Context, srcDir string, tarPath string, is2k bool) (err error) {

	defer func() {
		if rev := recover(); rev != nil {
			err = fmt.Errorf("dir2tar recover: %v, %v", rev, err)
			return
		}
	}()

	dir, err := os.Open(srcDir)
	if err != nil {
		return fmt.Errorf("os.Open(srcDir): %v", err)
	}
	defer dir.Close()

	tarfile, err := os.Create(tarPath)
	if err != nil {
		return fmt.Errorf("os.Create(tarPath): %v", err)
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		return fmt.Errorf("dir.Readdir(0): %v", err)
	}

	// 按照文件名字典序排序
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	modTime, _ := time.Parse("2006-01-02 15:04:05", tarModTime)
	for _, fi := range files {
		if !fi.Mode().IsRegular() {
			continue
		}

		hdr := &tar.Header{
			Typeflag: tar.TypeReg,
			Name:     fi.Name(),
			Mode:     0644,
			Size:     fi.Size(),
			ModTime:  modTime,
			Uid:      0,
			Gid:      0,
			Uname:    "",
			Gname:    "",
			Format:   tar.FormatUSTAR,
		}

		if err := tarball.WriteHeader(hdr); err != nil {
			return fmt.Errorf("WriteHeader: %v", err)
		}
		file, err := os.Open(filepath.Join(srcDir, fi.Name()))
		if err != nil {
			return fmt.Errorf("os.Open: %v", err)
		}
		defer file.Close()

		if _, err := io.Copy(tarball, file); err != nil {
			return fmt.Errorf("io.Copy: %v", err)
		}
	}
	return nil
}
