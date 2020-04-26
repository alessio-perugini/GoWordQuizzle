package main

//TODO utilizzare la waitgroup nella gestione del lancio delal partita da parte dei 2 giocatori

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:"+fmt.Sprintf("%d", common.TCP_PORT))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err := listen.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println()
		}

		go handlerConnection(conn)
	}
}

func handlerConnection(conn net.Conn) {
	fmt.Printf("%s [%s] connected \n", time.Now().String(), conn.RemoteAddr().String())

	for {
		cMsg, err := read(conn)
		if err != nil {
			log.Printf("%s [%s] closed connection.\n", time.Now().String(), conn.RemoteAddr().String())
			return
		}

		messageParser(conn, cMsg)
	}
}

func read(conn net.Conn) (string, error) {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		conn.Close()
		return "", err
	}

	return strings.TrimSpace(msg), nil
}

func write(conn net.Conn, msg string) {
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

func messageParser(conn net.Conn, msg string) {
	if msg == "" {
		log.Println("message to parse is empty")
		return
	}
	//TODO use reader instead of split
	token := strings.SplitN(msg, " ", 2)
	otherTokens := strings.Split(token[1], " ")

	if len(token) <= 0 { //TODO controllare meglio
		fmt.Println("token 0")
		return
	}

	command := token[0]
	switch command {
	case "LOGIN":
		username := otherTokens[0]
		pw := otherTokens[1]
		log.Println("LOGIN!")

		login(username, pw, conn)
		break
	case "LOGOUT":
		username := otherTokens[0]
		logout(username)
		break
	case "ADD_FRIEND":
		myNick := otherTokens[0]
		friendNick := otherTokens[1]

		addFriend(myNick, friendNick)
		break
	case "LISTA_AMICI":
		myNick := otherTokens[0]

		friendList(myNick)
		break
	case "SFIDA":
		myNick := otherTokens[0]
		friendNick := otherTokens[1]

		challenge(myNick, friendNick)
		break
	case "MOSTRA_SCORE":
		myNick := otherTokens[0]

		showScore(myNick)
		break
	case "MOSTRA_CLASSIFICA":
		myNick := otherTokens[0]

		showLeaderboard(myNick)
		break
	default:
		break
	}
}

func login(username, pw string, conn net.Conn) error {
	if username == "" {
		return errors.New("username not valid")
	}
	if pw == "" {
		return errors.New("password not valid")
	}

	u := common.NewUser()
	u.SetNickname(username)
	u.SetPssword(pw)

	ul := common.GetInstanceUsersList()
	if ul.AddUser(*u) != nil {
		return errors.New("User already exists")
	}

	fmt.Println(ul.GetUsers())

	write(conn, "Login success!\n")
	return nil
}

func logout(username string) error {
	if username == "" {
		return errors.New("username not valid")
	}

	return nil
}

func addFriend(myNick, friendNick string) error {
	if myNick == "" {
		return errors.New("myNick not valid")
	}
	if friendNick == "" {
		return errors.New("friendNick not valid")
	}

	return nil
}

func friendList(myNick string) error {
	if myNick == "" {
		return errors.New("myNick not valid")
	}

	return nil
}

func challenge(myNick, friendNick string) error {
	if myNick == "" {
		return errors.New("myNick not valid")
	}
	if friendNick == "" {
		return errors.New("friendNick not valid")
	}

	return nil
}

func showScore(myNick string) error {
	if myNick == "" {
		return errors.New("myNick not valid")
	}

	return nil
}

func showLeaderboard(myNick string) error {
	if myNick == "" {
		return errors.New("myNick not valid")
	}

	return nil
}
