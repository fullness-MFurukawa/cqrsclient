package provider_test

import (
	"context"
	"cqrsclient/infra"
	"cqrsclient/infra/provider"
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/providerパッケージのテスト")
}

var _ = Describe("ClientProvider", Ordered, Label("プロバイダの生成"), func() {
	var command *provider.CommandClientProvider
	var query *provider.QueryClientProvider
	var ctx context.Context
	var container *fx.App
	// 前処理
	BeforeAll(func() {
		// Contextの生成
		ctx = context.Background()
		// サービスのインスタンス生成
		container = fx.New(
			infra.InfraDepend,
			fx.Populate(&command),
			fx.Populate(&query),
		)
		// fxを起動し、起動時にエラーがないことを確認する
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	// 後処理
	AfterAll(func() {
		err := container.Stop(context.Background())
		Expect(err).NotTo(HaveOccurred())
	})

	It("QueryClientProviderの生成", func() {
		Expect(query.Category).NotTo(BeNil())
		Expect(query.Product).NotTo(BeNil())
		log.Println(query.Category)
		log.Println(query.Product)
	})
	It("CommandClientProviderの生成", func() {
		Expect(command.Category).NotTo(BeNil())
		Expect(command.Product).NotTo(BeNil())
		log.Println(command.Category)
		log.Println(command.Product)
	})
})
