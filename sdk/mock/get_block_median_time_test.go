package mock

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlockMedianTime(t *testing.T) {

	t.Run("get_block_median_time/[block_hash]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetBlockMedianTime(context.Background(), types.HexToHash(mockData.Request.Params[0].(string)))
		g.Expect(err).To(gomega.BeNil(), "GetBlockMedianTime failed")

		mockResult, err := interfaceToUint(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToUint failed")

		g.Expect(info).To(gomega.Equal(uint64(mockResult)), "Result Unequal")
	})
}
