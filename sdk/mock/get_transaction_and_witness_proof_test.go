package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetTransactionAndWitnessProof(t *testing.T) {
	// 当blockhash 为nil 时，请求中为None，但是mock中没有为空
	t.Run("get_transaction_and_witness_proof/[hashs]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		txHashs, err := interfaceSliceToStringSlice(mockData.Request.Params[0].([]interface{}))
		g.Expect(err).To(gomega.BeNil(), "txHashs interfaceSliceToStringSlice failed")

		info, err := client.GetTransactionAndWitnessProof(context.Background(), txHashs, nil)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching TxPoolInfo")
		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")

		g.Expect(info.BlockHash).To(gomega.Equal(types.HexToHash(localresult["block_hash"].(string))))
		fmt.Println(info.BlockHash, localresult["block_hash"].(string))
		// Description with marker

	})

	t.Run("get_transaction_and_witness_proof/[tx_hashs,block_hash]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		txHashs, err := interfaceSliceToStringSlice(mockData.Request.Params[0].([]interface{}))
		g.Expect(err).To(gomega.BeNil(), "txHashs interfaceSliceToStringSlice failed")

		empty32Bytes := types.HexToHash(mockData.Request.Params[1].(string))
		info, err := client.GetTransactionAndWitnessProof(context.Background(), txHashs, &empty32Bytes)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching TxPoolInfo")
		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")

		g.Expect(info.BlockHash).To(gomega.Equal(types.HexToHash(localresult["block_hash"].(string))))
		fmt.Println(info.BlockHash, localresult["block_hash"].(string))
		// Description with marker

	})
}
