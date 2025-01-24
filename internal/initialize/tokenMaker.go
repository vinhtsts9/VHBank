package initialize

import (
	"Golang-Masterclass/simplebank/global"
	"Golang-Masterclass/simplebank/util/token"
	"log"
)

func InitTokenMaker() {
	tokenMaker, err := token.NewPasetoMaker(global.Config.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %w", err)
	}
	global.TokenMaker = tokenMaker

}
