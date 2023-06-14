package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read commands from the client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read from client:", err)
		return
	}

	command := string(buffer[:n])

	// Process the command and generate a response
	response := processCommand(command)

	// Send the response to the client
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Failed to send response to client:", err)
	}
}

func processCommand(command string) string {
	// Process the command and generate the response
	// You can implement your logic here based on the command received
	// and generate the appropriate response

	return "Response to the command: " + command
}

func main() {
	// Create a Unix domain socket

	dir := "."
	absDir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	socketPath := filepath.Join(absDir, "socket") // "/path/to/unix/socket"
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Failed to create socket:", err)
		return
	}
	defer os.Remove(socketPath)
	defer l.Close()

	// Handle termination signals to gracefully stop the daemon
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start accepting and handling client connections
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Failed to accept client connection:", err)
				break
			}

			go handleConnection(conn)
		}
	}()

	fmt.Println("Daemon is running")

	// Wait for termination signal
	<-stop

	fmt.Println("Stopping daemon")
}

func f(){

// Connect to the daemon
	conn, err := net.Dial("unix", "/path/to/unix/socket")
	if err != nil {
		fmt.Println("Failed to connect to the daemon:", err)
		return
	}
	defer conn.Close()

	// Read commands from stdin and send them to the daemon
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		// Send the command to the daemon
		_, err := conn.Write([]byte(command))
		if err != nil {
			fmt.Println("Failed to send command to the daemon:", err)
			break
		}

		// Read the response from the daemon
		response := make([]byte, 1024)
		_, err = conn.Read(response)
		if err != nil {
			fmt.Println("Failed to read response from the daemon:", err)
			break
		}

		fmt.Println("Response:", string(response))
	}
}


func HanMsg(){
	
}

type Msger interface {
	MID() string
	MType() string	
}

type Request struct{
		
}

type Reseonse struct{
	
}


func dispatch(msg Msger){
	switch msg.MType(){
	case "":
	default:
		
	}
}

func action(req Request) (res Reseonse){
	
}
