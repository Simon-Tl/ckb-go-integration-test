package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetIndexerTip(t *testing.T) {

	t.Run("get_indexer_tip/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.GetIndexerTip(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetIndexerTip")
		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")
		// Description with marker
		g.Expect(info.BlockHash.Hex()).To(gomega.Equal(localresult["block_hash"]),
			"Expected BlockHash to match mock data")
		fmt.Println(info) // Description added for clarity
	})
}
