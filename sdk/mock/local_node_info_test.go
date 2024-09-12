package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestLocalNodeInfo(t *testing.T) {
	t.Run("local_node_info/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceSliceToMapString failed")

		info, err := client.LocalNodeInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")
		g.Expect(info.NodeId).To(gomega.Equal(localresult["node_id"]), "Expected Value mismatch")

	})

}
