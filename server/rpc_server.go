package server

import (
	"errors"
	"fmt"
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	Username, Password string
}

type UserRegistration int

func (r *UserRegistration) RegisterUser(args *Args, reply *bool) error {
	if args.Username == "" || args.Password == "" {
		return errors.New("username or password not set")
	}

	ul := common.GetInstanceUsersList()
	u := common.User{}
	u.SetNickname(args.Username)
	u.SetPssword(args.Password)

	if err := ul.AddUser(u); err != nil {
		return errors.New("username already exists")
	}

	*reply = true
	return nil
}

func StartRPC() {
	uReg := new(UserRegistration)
	err := rpc.Register(uReg)
	if err != nil {
		log.Fatal("format of service Task isn't correct. ", err)
	}

	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":"+fmt.Sprint(common.RPC_PORT))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("serving RPC server on port %d\n", common.RPC_PORT)

	go func() {
		if err := http.Serve(listener, nil); err != nil {
			log.Fatal("error serving: ", err)
		}
	}()

}
