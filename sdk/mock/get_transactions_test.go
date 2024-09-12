package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetTransactions(t *testing.T) {

	t.Run("get_transactions/[search_key,order,limit,null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		Parmsscript, err := interfaceToMapString(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "Parmsscript parser failed")

		scripts, err := interfaceToMapString(Parmsscript["script"])
		args, err := interfaceToBytes(scripts["args"])
		g.Expect(err).To(gomega.BeNil(), "args parser failed")
		script := types.Script{
			CodeHash: types.HexToHash(scripts["code_hash"].(string)),
			HashType: types.HashTypeType,
			Args:     args,
		}

		indexerSearchKey := indexer.SearchKey{
			Script:           &script,
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
		}

		limit, err := interfaceToUint(mockData.Request.Params[2])
		fmt.Println(limit)
		limit64 := uint64(limit)

		info, err := client.GetTransactions(context.Background(), &indexerSearchKey, indexer.SearchOrderAsc, limit64, "None")
		g.Expect(err).To(gomega.BeNil(), "GetBlockHash failed")
		fmt.Println(info.LastCursor)
		fmt.Println(mockData.Response.Result)
		//g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result.(string)))

	})

	t.Run("get_block_hash/null", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		Parms, err := interfaceToUint(mockData.Request.Params[0])
		ParmUint64 := uint64(Parms)
		info, err := client.GetBlockHash(context.Background(), ParmUint64)
		g.Expect(err).To(gomega.BeNil(), "GetBlockHash failed")
		fmt.Println(info)
		if mockData.Response.Result == nil {
			mockData.Response.Result = "0x0000000000000000000000000000000000000000000000000000000000000000"
			g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result))
		}
		//g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result.(string)))

	})
}
