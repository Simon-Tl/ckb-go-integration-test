package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestTxPoolReady(t *testing.T) {

	t.Run("tx_pool_ready/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info,err := client.TxPoolReady(context.Background())
		
		g.Expect(err).To(gomega.BeNil(), "ClearTxPool failed")
		
		
		fmt.Println(mockdata.Response.Result)
		g.Expect(info).To(gomega.Equal(mockdata.Response.Result))
		// Description with marker
	})
}
