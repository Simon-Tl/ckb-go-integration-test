package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetCellsCapacity(t *testing.T) {

	t.Run("get_cells_capacity/[search_key]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		Parms, err := interfaceToMapString(mockData.Request.Params)
		Parm := Parms["script"].(map[string]interface{})

		arg, err := interfaceToBytes(Parm["args"])

		script := types.Script{
			CodeHash: types.HexToHash(Parm["code_hash"].(string)),
			HashType: types.HashTypeType,
			Args:     arg,
		}

		indexerSearch := indexer.SearchKey{
			Script:           &script,
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
		}

		info, err := client.GetCellsCapacity(context.Background(), &indexerSearch)
		fmt.Println(indexerSearch)
		fmt.Println(info)
		// Description with marker
	})

}
