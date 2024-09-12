package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestClearTxPool(t *testing.T) {

	t.Run("clear_tx_pool/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, _, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info := client.ClearTxPool(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetHeader")

		g.Expect(info).To(gomega.BeNil())
		// Description with marker
	})
}
