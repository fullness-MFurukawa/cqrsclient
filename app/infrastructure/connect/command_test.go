package connect_test

import (
	"client/infrastructure/connect"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandConnect(t *testing.T) {
	connector := connect.CommandConnector{}
	connect, err := connector.Connect()
	if err != nil {
		log.Println(err.Error())
		assert.Fail(t, "テスト失敗")
	}
	defer connect.Close()
	log.Println(connect)
	assert.True(t, true)
}
