PKG := "C0lliNN/card-game-service"
PKG_LIST := $(shell go list ${PKG}/...)

test:
	@go test -v ${PKG_LIST}

start:
	@go run C0lliNN/card-game-service/cmd/game-service-api