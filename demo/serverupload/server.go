package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"
)

const (
	blockSize       = 64 * 1024 // 64KB
	receiveBuffSize = 4 * 1024  // 4KB
	channelSize     = 100       // 缓冲通道大小
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取文件元信息
	metaInfo, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read file metadata:", err)
		return
	}

	parts := strings.Split(strings.TrimSpace(metaInfo), " ")
	fileName, fileSize, err := parts[0], mustParseInt(parts[1]), nil
	if err != nil {
		fmt.Println("Invalid file metadata:", err)
		return
	}

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	n, err := io.CopyN(file, conn, int64(fileSize))

	fmt.Println(fileSize, n, err)

	// // 创建缓冲通道
	// dataChan := make(chan []byte, channelSize)

	// // 启动读取goroutine
	// go receiveData(conn, dataChan, fileSize)

	// writeData(file, dataChan)
}

func receiveData(conn net.Conn, dataChan chan<- []byte, totalSize int) {

	defer conn.Close()

	// 设置接收缓冲区大小
	conn.(*net.TCPConn).SetReadBuffer(receiveBuffSize)

	// 接收文件内容
	receiveBuffer := make([]byte, blockSize)
	bytesReceived := 0
	for {
		bytesRead, err := conn.Read(receiveBuffer)
		fmt.Println(bytesRead, err)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to receive data:", err)
			return
		}

		dataChan <- receiveBuffer[:bytesRead]
		bytesReceived += bytesRead
		if bytesReceived == totalSize {
			break
		}
	}

	// 关闭通道,通知写入goroutine结束
	close(dataChan)
}

func writeData(file *os.File, dataChan <-chan []byte) {
	for data := range dataChan {

		fmt.Println("data", len(data))

		_, err := file.Write(data)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ServerCmd() *cli.Command {
	return &cli.Command{
		Name: "s",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Value: "0.0.0.0:17733",
			},
		},
		Action: ServerMainAction,
	}
}

func ServerMainAction(cCtx *cli.Context) error {
	serverAddr := cCtx.String("addr")
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("Failed to listen:", err)
	}
	defer listener.Close()

	fmt.Println("Server started, waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func ServerHttp() {
	http.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a new file for storing the uploaded content
	uploadedFile, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	// Use sendfile to efficiently transfer data from the request file to the uploaded file
	if _, err := syscall.Sendfile(int(uploadedFile.Fd()), int(file.(interface {
		// Fd returns the file descriptor number.
		Fd() uintptr
	}).Fd()), nil, int(header.Size)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", header.Filename)
}
