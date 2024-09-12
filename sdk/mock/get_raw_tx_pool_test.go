package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetRawTxPool(t *testing.T) {

	t.Run("get_raw_tx_pool/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, _, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.GetRawTxPool(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetHeader")

		fmt.Println(info)
		// Description with marker
	})
}
