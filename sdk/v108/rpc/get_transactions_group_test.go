package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetTransactionsGroup(t *testing.T) {

	t.Run("should return msg ,when script_search_mode == extra", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		txsResponse, err := sdk.Client.GetTransactionsGrouped(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
			WithData:         false,
		}, indexer.SearchOrderAsc, 1000, "")
		gomega.Expect(err, nil)
		checkGetTransactionsGroupedResponse(txsResponse)
	})

	t.Run("should return msg ,when script_search_mode == pre", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		txsResponse, err := sdk.Client.GetTransactionsGrouped(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
			WithData:         false,
		}, indexer.SearchOrderAsc, 1000, "")
		gomega.Expect(err, nil)
		checkGetTransactionsGroupedResponse(txsResponse)

	})

	t.Run("should return msg ,when script_search_mode == nil", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		txsResponse, err := sdk.Client.GetTransactionsGrouped(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType: types.ScriptTypeLock,
			WithData:   false,
		}, indexer.SearchOrderAsc, 1000, "")
		gomega.Expect(err, nil)
		checkGetTransactionsGroupedResponse(txsResponse)
	})

}

type CellMsg struct {
	blockNum uint64
	blockIdx uint64
	ioType   indexer.IoType
	ioTdx    indexer.IoType
	ioIdx    uint64
}

func checkGetTransactionsGroupedResponse(txsResponse *indexer.TxsWithCells) {
	// query one data check tx idx is right
	blockResponse, err := sdk.Client.GetBlockByNumber(sdk.Ctx, txsResponse.Objects[0].BlockNumber)
	gomega.Expect(blockResponse.Transactions[txsResponse.Objects[0].TxIndex], txsResponse.Objects[0].TxHash)

	// query cell CodeHash && HashType is right ,skip arg
	txResponse, err := sdk.Client.GetTransaction(sdk.Ctx, txsResponse.Objects[0].TxHash, &sdk.Enable)
	if txsResponse.Objects[0].Cells[0].IoType == indexer.IOTypeIn {
		cellInput := txResponse.Transaction.Inputs[txsResponse.Objects[0].Cells[0].IoIndex]
		cellOutPut, err := getCellOutputByHashAndIdx(cellInput.PreviousOutput.TxHash, cellInput.PreviousOutput.Index)
		gomega.Expect(err, nil)
		gomega.Expect(cellOutPut.Lock.CodeHash, types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"))
		gomega.Expect(cellOutPut.Lock.HashType, types.HashTypeData)
		return
	}
	cellOutput := txResponse.Transaction.Outputs[txsResponse.Objects[0].Cells[0].IoIndex]
	gomega.Expect(err, nil)
	gomega.Expect(cellOutput.Lock.CodeHash, types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"))
	gomega.Expect(cellOutput.Lock.HashType, types.HashTypeData)

}
func getCellOutputByHashAndIdx(txHash types.Hash, idx uint32) (*types.CellOutput, error) {
	tx, err := sdk.Client.GetTransaction(sdk.Ctx, txHash, &sdk.Enable)
	if err != nil {
		return nil, err
	}
	return tx.Transaction.Outputs[idx], nil
}
