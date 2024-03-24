package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"syscall"

	"github.com/urfave/cli/v2"
)

const (
	// blockSize      = 64 * 1024 // 64KB
	sendBufferSize = 4 * 1024 // 4KB
// concurrency    = 4         // 并发数
// channelSize    = 100       // 数据通道缓冲区大小
)

func sendFile(ctx context.Context, filePath string, serverAddr string, spath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("file stat: %w", err)
	}

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("dial server: %w", err)
	}
	defer conn.Close()

	// fileName := filepath.Base(filePath)
	fmt.Fprintf(conn, "%s %d\n", spath, fileInfo.Size())

	conn.(*net.TCPConn).SetWriteBuffer(sendBufferSize)

	dataChan := make(chan []byte, channelSize)

	go readFile(ctx, file, dataChan, fileInfo.Size())

	fmt.Println("abc")
	for data := range dataChan {
		fmt.Println("cccc", len(data))
		_, err = conn.Write(data)
		if err != nil {
			return fmt.Errorf("Failed to send data:", err)
		}
	}

	fmt.Println("File sent successfully!")
	return nil
}

func readFile(ctx context.Context, file *os.File, dataChan chan<- []byte, totalSize int64) {
	sendBuffer := make([]byte, blockSize)
	bytesRead := 0
	for {
		n, err := file.Read(sendBuffer)
		fmt.Println(n, err)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}

		dataChan <- sendBuffer[:n]
		bytesRead += n
		if int64(bytesRead) == totalSize {
			break
		}

		select {
		case <-ctx.Done():
			break
		default:
		}
	}

	close(dataChan)
}

func ClientCmd() *cli.Command {
	return &cli.Command{
		Name: "c",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Value: "127.0.0.1:17733",
			},
			&cli.StringFlag{
				Name: "file",
			},
			&cli.StringFlag{
				Name: "spath",
			},
		},
		Action: ClientMainAction,
	}
}

func ClientMainAction(cCtx *cli.Context) error {
	serverAddr := cCtx.String("addr")
	file := cCtx.String("file")
	spath := cCtx.String("spath")
	return sendFile(cCtx.Context, file, serverAddr, spath)
}

func SendHttp(ctx context.Context, filePath string) error {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting:", err)
	}
	defer conn.Close()

	// Create a file to write the received data
	outputFile, err := os.Create("received_file.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer outputFile.Close()

	fh, err := conn.(*net.TCPConn).File()

	if _, err := syscall.Sendfile(int(outputFile.Fd()), int(fh.Fd()), nil, 1024); err != nil {
		log.Fatal("Error sending file:", err)
	}

	fmt.Println("File received successfully")

	return nil
}
