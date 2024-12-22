package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
)

func TestRemoveTransaction(t *testing.T) {

	t.Run("remove_transaction/[tx_hash]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		Parms := mockdata.Request.Params[0].(string)
		info, err := client.RemoveTransaction(context.Background(),types.HexToHash(Parms))

		g.Expect(err).To(gomega.BeNil(), "RemoveTransaction failed")

		fmt.Println(mockdata.Response.Result)
		g.Expect(info).To(gomega.Equal(mockdata.Response.Result))
		// Description with marker
	})
}
