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
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info := client.ClearBannedAddresses(context.Background())
		g.Expect(info).To(gomega.BeNil(), "Returns the result error")
		g.Expect(mockData.Response.Result).To(gomega.BeNil(), "Returns the result error")
	})
}
