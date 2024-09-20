package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

// mock 方法为get_pool_tx_detail_info，sdk提供方法为get_pool_tx_detail
func TestGetPoolTxDetailInfo(t *testing.T) {

	t.Run("get_pool_tx_detail_info/[tx_hash]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetPoolTxDetailInfo(context.Background(), types.HexToHash(mockData.Request.Params[0].(string)))
		g.Expect(err).To(gomega.BeNil(), "GetPoolTxDetailInfo failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")
		// Description with marker
		g.Expect(info.Timestamp).To(gomega.Equal(mockResult["timestamp"]),
			"Result Unequal")
		fmt.Println(info) // Description added for clarity
	})
}
