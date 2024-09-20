package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestAddNode(t *testing.T) {

	t.Run("add_node/[peer_id,address]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)
		peer_id := fmt.Sprintf("%v", mockData.Request.Params[0])
		address := fmt.Sprintf("%v", mockData.Request.Params[1])

		fmt.Println(peer_id, address)
		info := client.AddNode(context.Background(), peer_id, address)
		g.Expect(info).To(gomega.BeNil(), "Returns the result error")

	})
}
