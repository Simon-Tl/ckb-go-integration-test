package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	t.Run("get_transaction/[tx_hash,verbosity,only_committed=null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation\d

		Hash := mockData.Request.Params[0]

		info, err := client.GetTransaction(context.Background(), types.HexToHash(Hash.(string)), nil)
		fmt.Println(*info.Cycles)

		// Description with marker
	})

	t.Run("get_transaction/min_replace_fee,fee", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation\d

		Hash := mockData.Request.Params[0]

		info, err := client.GetTransaction(context.Background(), types.HexToHash(Hash.(string)), nil)
		fmt.Println(*info.Cycles)

		// Description with marker
	})

}
