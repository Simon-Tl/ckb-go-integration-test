package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestTxPoolInfo(t *testing.T) {

	t.Run("tx_pool_info/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.TxPoolInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching TxPoolInfo")
		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")
		// Description with marker
		g.Expect(info.TipHash.Hex()).To(gomega.Equal(localresult["tip_hash"]),
			"Expected TipHash to match mock data") // Description added for clarity
	})
}
