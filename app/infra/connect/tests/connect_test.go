package connect_test

import (
	"cqrsclient/infra/connect"
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/connectパッケージのテスト")
}

var _ = Describe("CommandService QueryService接続", Ordered, Label("サービスへの接続テスト"), func() {
	var query connect.ServerConnector
	var command connect.ServerConnector
	// 前処理
	BeforeAll(func() {
		query = connect.NewqueryConnector()
		command = connect.NewcommandConnector()
	})
	It("QueryServiceへの接続", func() {
		connect, err := query.Connect()
		if err != nil {
			Expect(err).Error()
			return
		}
		defer connect.Close()
		log.Println(connect)
		Expect(connect).NotTo(BeNil())
	})
	It("CommandServiceへの接続", func() {
		connect, err := command.Connect()
		if err != nil {
			Expect(err).Error()
			return
		}
		defer connect.Close()
		log.Println(connect)
		Expect(connect).NotTo(BeNil())
	})
})
