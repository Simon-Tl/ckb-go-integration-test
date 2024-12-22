package mock

import (
	"context"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
)

func TestTruncate(t *testing.T) {

	t.Run("truncate/[target_tip_hash]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		errMsg := client.Truncate(context.Background(), types.HexToHash(mockdata.Request.Params[0].(string)))

		g.Expect(errMsg).To(gomega.BeNil(), "ClearTxPool failed")

		// Description with marker
	})
}
