package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetPeers(t *testing.T) {
	t.Run("get_peers/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")

		info, err := client.GetPeers(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetPeers")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceSliceToMapString failed")
		// Description with marker
		g.Expect(info[0].NodeID).To(gomega.Equal(localresult["node_id"]))
		fmt.Println(info[0].NodeID)
		fmt.Println(localresult["node_id"])

	})

}
