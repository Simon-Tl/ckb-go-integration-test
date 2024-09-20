package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

// 缺少extension 返回值

func TestGetBlockByNumber(t *testing.T) {

	t.Run("get_block_by_number/[block_number]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		mockParams, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "Params parser failed")

		info, err := client.GetBlockByNumber(context.Background(), uint64(mockParams))
		g.Expect(err).To(gomega.BeNil(), "GetBlockByNumber failed")

		fmt.Println(info)

	})
}
