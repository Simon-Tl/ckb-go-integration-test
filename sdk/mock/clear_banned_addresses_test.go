package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestClearBannedAddresses(t *testing.T) {

	t.Run("clear_banned_addresses/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info := client.ClearBannedAddresses(context.Background())
		g.Expect(info).To(gomega.BeNil(), "info is nil")
		g.Expect(mockData.Response.Result).To(gomega.BeNil(), "mockData.Response.Result")
		// Description with marker
	})
}
