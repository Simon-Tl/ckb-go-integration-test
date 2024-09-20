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
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetIndexerTip(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetIndexerTip failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		g.Expect(info.BlockHash.Hex()).To(gomega.Equal(mockResult["block_hash"]),
			"Result Unequal")
		fmt.Println(info) // Description added for clarity
	})
}
