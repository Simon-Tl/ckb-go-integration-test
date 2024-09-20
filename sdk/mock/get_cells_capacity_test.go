package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

// get_cells_capacity IndexerSearchKey结构体中的  Filter 字段缺少“output_data，output_data_filter_mode” 多出来的withdata,导致数据无法返回

func TestGetCellsCapacity(t *testing.T) {

	t.Run("get_cells_capacity/[search_key]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		mockParams, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "mockParams interfaceToMapString ")

		mockScriptParams, err := interfaceToMapString(mockParams["script"])
		g.Expect(err).To(gomega.BeNil(), "mockScriptParams interfaceToMapString failed")

		mockArg, err := interfaceToBytes(mockScriptParams["args"])
		g.Expect(err).To(gomega.BeNil(), "mockArg interfaceToBytes failed")

		script := types.Script{
			CodeHash: types.HexToHash(mockScriptParams["code_hash"].(string)),
			HashType: types.ScriptHashType(mockScriptParams["hash_type"].(string)),
			Args:     mockArg,
		}

		indexerSearch := indexer.SearchKey{
			Script:           &script,
			ScriptType:       types.ScriptType(mockParams["script_type"].(string)),
			ScriptSearchMode: types.ScriptSearchMode(mockParams["script_search_mode"].(string)),
		}

		info, err := client.GetCellsCapacity(context.Background(), &indexerSearch)
		fmt.Println(indexerSearch)
		fmt.Println(info)
		// Description with marker
	})

}
