package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestPingPerrs(t *testing.T) {

	t.Run("ping_peers/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info := client.PingPeers(context.Background())
		g.Expect(info).To(gomega.BeNil(), "result errors")
		g.Expect(mockdata.Response.Result).To(gomega.BeNil(), "mockResult errors")

		// Description with marker
	})
}
