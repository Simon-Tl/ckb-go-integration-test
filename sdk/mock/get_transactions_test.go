package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

// get_transactions   SearchKey 多出来的WithData 导致无法返回
func TestGetTransactions(t *testing.T) {

	t.Run("get_transactions/[search_key,order,limit,null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		Params, err := interfaceToMapString(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "Params interfaceToMapString failed")

		ParamsScript, err := interfaceToMapString(Params["script"])
		g.Expect(err).To(gomega.BeNil(), "ParamsScript interfaceToMapString failed")

		ParamsScriptArgs, err := interfaceToBytes(ParamsScript["args"])
		g.Expect(err).To(gomega.BeNil(), "ParamsScriptArgs interfaceToBytes failed")
		script := types.Script{
			CodeHash: types.HexToHash(ParamsScript["code_hash"].(string)),
			HashType: types.HashTypeType,
			Args:     ParamsScriptArgs,
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

}
