package main

import (
	"log"
)

const (
	SERVER_HOST        = "localhost"
	SERVER_PORT        = 8332
	USER               = "rpczh01"
	PASSWD             = "123rpczh01456"
	USESSL             = false
	WALLET_PASSPHRASE  = "WalletPassphrase"
)

func main() {
	test02()

}

func test02(){
	rpcClient, err := newClient(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}
	//生成一个新地址
	reqJson := "{\"method\":\"getnewaddress\",\"params\":[\"labelName002\"],\"id\":1}";
	returnJson, err2 := rpcClient.send(reqJson)
	if err2 != nil {
		log.Fatalln(err2)
	}
	log.Println("returnJson:", returnJson)
}
