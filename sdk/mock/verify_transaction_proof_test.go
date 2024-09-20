package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestVerifyTransactionProof(t *testing.T) {
	t.Run("verify_transaction_proof/[tx_proof]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		data, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "data interfaceSliceToMapString failed")

		proofType, err := interfaceToMapString(data["proof"])
		g.Expect(err).To(gomega.BeNil(), "proofType interfaceToMapString failed")

		indices, err := interfaceToUintSlice(proofType["indices"])
		g.Expect(err).To(gomega.BeNil(), "indices interfaceToUintSlice failed")

		lemmas, err := interfaceToHashSlice(proofType["lemmas"])
		g.Expect(err).To(gomega.BeNil(), "lemmas interfaceToHashSlice failed")

		prooftype := types.Proof{
			Indices: indices,
			Lemmas:  lemmas,
		}
		transactionData := types.TransactionProof{
			Proof:         &prooftype,
			BlockHash:     types.HexToHash(data["block_hash"].(string)),
			WitnessesRoot: types.HexToHash(data["witnesses_root"].(string)),
		}

		fmt.Println(data)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation		fmt.Println(mockData.Request.Params)
		info, err := client.VerifyTransactionProof(context.Background(), &transactionData)
		g.Expect(err).To(gomega.BeNil(), "VerifyTransactionProof failed")

		fmt.Println(info)
		fmt.Println(mockData.Response.Result)
		g.Expect(hashSliceToInterfaceSlice(info)).To(gomega.Equal(mockData.Response.Result), "Result Unequal")

	})

}
