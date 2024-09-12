package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlockByNumber(t *testing.T) {

	t.Run("get_block_by_number/[block_number]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		Params, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "Params parser failed")
		info, err := client.GetBlockByNumber(context.Background(), uint64(Params))
		fmt.Println(info)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetBlockByNumber")

		// Description with marker
	})
}
