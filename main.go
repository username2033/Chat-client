package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var mutex sync.Mutex

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	//mutex.Lock()
	//fmt.Println("Введите Ваше имя пользователя (username)")
	//conn.Write([]byte(scanner.Text() + "\n"))
	//fmt.Println("Введите Ваш пароль")
	//conn.Write([]byte(scanner.Text() + "\n"))
	//mutex.Unlock()

	go func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(message)
		}
	}()

	for scanner.Scan() {
		text := scanner.Text()
		conn.Write([]byte(text + "\n"))
	}
}
