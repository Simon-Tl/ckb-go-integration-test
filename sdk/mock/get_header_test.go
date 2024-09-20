package mock

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

// GetHeader缺少参数verbosity

func TestGetHeader(t *testing.T) {

	t.Run("get_header/[block_hash,verbosity=1]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetHeader(context.Background(), types.HexToHash(mockData.Request.Params[0].(string)))
		g.Expect(err).To(gomega.BeNil(), "GetHeader failed")

		mockResult, err := interfaceToUint(mockData.Response.Result)
		g.Expect(info).To(gomega.Equal(uint64(mockResult)), "Result Unequal")
		// Description with marker
	})
}
