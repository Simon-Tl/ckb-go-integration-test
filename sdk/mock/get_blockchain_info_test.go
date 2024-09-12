package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlockchainInfo(t *testing.T) {

	t.Run("get_blockchain_info/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.GetBlockchainInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching SetNetworkActive")

		localresult, _ := interfaceToMapString(mockData.Response.Result)
		localresultNum, _ := interfaceToUint(localresult["epoch"])
		g.Expect(info.Epoch).To(gomega.Equal(uint64(localresultNum)))
		// Description with marker
	})
}
