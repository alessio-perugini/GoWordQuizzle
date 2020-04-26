package main

import (
	"bufio"
	"fmt"
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"log"
	"net"
	"os"
	"strings"
	"text/scanner"
	"time"
)

var (
	quit    bool
	s       scanner.Scanner
	profile common.User
)

var conn *net.TCPConn

func main() {
	serverAddr := "localhost:" + fmt.Sprint(common.TCP_PORT)
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && !quit {
		command := scanner.Text()

		//TODO scanner on command
		s.Init(strings.NewReader(command))
		token := common.ReadNextToken(&s)

		switch token {
		case "quit":
			os.Exit(1)
			break
		case "register":
			break
		case "login":
			login()
			break
		case "logout":
			break
		case "add_friend":
			break
		case "friendlist":
			break
		case "challenge":
			break
		case "show_score":
			break
		case "show_leaderboard":
			break
		case "help":
			break
		}
	}
	//TODO remove
	conn.Write([]byte{0})

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func login() {
	username := common.ReadNextToken(&s)
	pw := common.ReadNextToken(&s)

	if username == "" {
		log.Println("username not valid")
	}
	if pw == "" {
		log.Println("pw not valid")
	}
	if profile.GetNickname() == username {
		log.Println("already loggedin as: " + username)
	}
	if profile.GetNickname() != "" {
		logout()
	}

	//TODO change udp port assignment
	cmdToSend := fmt.Sprintf("LOGIN %s %s %d\n", username, pw, common.UDP_PORT)
	write(cmdToSend)
	result := read()
	fmt.Printf("%s <- %s", time.Now().String(), result)
}

func logout() {

}

func read() string {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Close()
		log.Println("Error on Write" + err.Error())
		return ""
	}

	return msg
}

func write(msg string) {
	writer := bufio.NewWriter(conn)
	byteWritten, err := writer.WriteString(msg)

	if err != nil {
		conn.Close()
		log.Println("Error on Write" + err.Error())
		return
	} else {
		err = writer.Flush() //TODO handle err
	}

	if byteWritten < len(msg) {
		fmt.Printf("Bytes left to wrote: %d \n", len(msg)-byteWritten)
	}
}
