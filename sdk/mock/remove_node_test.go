package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestRemoveNode(t *testing.T) {

	t.Run("remove_node/[peer_id]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info := client.RemoveNode(context.Background(), mockData.Request.Params[0].(string))
		g.Expect(info).To(gomega.BeNil(), "Expected no error while fetching SetNetworkActive")
		// Description with marker
	})
}
