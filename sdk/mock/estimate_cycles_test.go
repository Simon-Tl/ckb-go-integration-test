package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"strconv"
	"testing"
)

func TestEstimateCycles(t *testing.T) {
	t.Run("estimate_cycles/[tx]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())

		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")
		Params, err := interfaceSliceToMapString(mockData.Request.Params)

		g.Expect(err).To(gomega.BeNil(), "Params interfaceSliceToMapString failed")

		version, err := interfaceToUint(Params["version"])
		g.Expect(err).To(gomega.BeNil(), "version interfaceToUint failed")
		version32 := uint32(version)

		celldepmap, err := interfaceToMapString(Params["cell_deps"])
		fmt.Println(celldepmap)
		g.Expect(err).To(gomega.BeNil(), "celldepmap interfaceSliceToMapString failed")

		CellDepBytes := make([]*types.CellDep, 0) // 初始化为空切片

		celldep, err := interfaceToMapString(celldepmap)

		g.Expect(err).To(gomega.BeNil(), "celldep interfaceToMapString failed")

		outpointmap, err := interfaceToMapString(celldep["out_point"])
		g.Expect(err).To(gomega.BeNil(), "outpointmap interfaceToMapString failed")

		indexer, err := interfaceToUint(outpointmap["index"])
		g.Expect(err).To(gomega.BeNil(), "indexer interfaceToUint failed")
		indexer32 := uint32(indexer)

		outpoint := types.OutPoint{
			TxHash: types.HexToHash(outpointmap["tx_hash"].(string)),
			Index:  indexer32,
		}

		cellDep := types.CellDep{
			DepType:  types.DepTypeDepGroup,
			OutPoint: &outpoint,
			// 使用正确的 DepType
		}

		CellDepBytes = append(CellDepBytes, &cellDep)

		headerdeps, err := interfaceToHashSlice(Params["header_deps"])

		inputmap, err := interfaceToMapString(Params["inputs"])
		g.Expect(err).To(gomega.BeNil(), "input interfaceToUint failed")

		InputsBytes := make([]*types.CellInput, 0) // 初始化为空切片

		inputMap, err := interfaceToMapString(inputmap)
		g.Expect(err).To(gomega.BeNil(), "inputMap interfaceToMapString failed")

		previousoutputmap, err := interfaceToMapString(inputMap["previous_output"])
		g.Expect(err).To(gomega.BeNil(), "previous_output interfaceToMapString failed")

		since, err := interfaceToUint(inputMap["since"])
		g.Expect(err).To(gomega.BeNil(), "since interfaceToUint failed")
		since64 := uint64(since)

		previousoutindex, err := interfaceToUint(previousoutputmap["index"])
		g.Expect(err).To(gomega.BeNil(), "previousoutindex interfaceToUint failed")
		previousoutindex32 := uint32(previousoutindex)

		previousout := types.OutPoint{
			TxHash: types.HexToHash(previousoutputmap["tx_hash"].(string)),
			Index:  previousoutindex32,
		}
		input := types.CellInput{
			Since:          since64,
			PreviousOutput: &previousout,
		}

		InputsBytes = append(InputsBytes, &input)

		outputList, ok := Params["outputs"].([]interface{})
		g.Expect(ok).To(gomega.BeTrue(), "outputs should be a list")

		OutputsBytes := make([]*types.CellOutput, 0) // 初始化为空切片

		for _, output := range outputList {
			outputMap, err := interfaceToMapString(output)
			g.Expect(err).To(gomega.BeNil(), "outputMap interfaceToMapString failed")

			lockMap, err := interfaceToMapString(outputMap["lock"])
			g.Expect(err).To(gomega.BeNil(), "lockMap interfaceToMapString failed")

			args, err := interfaceToBytes(lockMap["args"])
			g.Expect(err).To(gomega.BeNil(), "args interfaceToBytes failed")

			lock := types.Script{
				CodeHash: types.HexToHash(lockMap["code_hash"].(string)),
				HashType: types.ScriptHashType(lockMap["hash_type"].(string)), // 根据实际情况设置
				Args:     args,
			}

			capacityStr := outputMap["capacity"].(string)
			capacity, err := strconv.ParseUint(capacityStr[2:], 16, 64)
			g.Expect(err).To(gomega.BeNil(), "capacity conversion failed")

			outpus := types.CellOutput{
				Capacity: capacity,
				Lock:     &lock,
				Type:     nil, // 根据实际情况设置
			}

			OutputsBytes = append(OutputsBytes, &outpus)
		}

		outputDataList, ok := Params["outputs_data"].([]interface{})
		g.Expect(ok).To(gomega.BeTrue(), "outputs_data should be a list")

		OutputDataBytes := make([][]byte, 0)
		for _, d := range outputDataList {
			data, err := interfaceToBytes(d)
			g.Expect(err).To(gomega.BeNil(), "outputs_data interfaceToBytes failed")
			OutputDataBytes = append(OutputDataBytes, data)
		}

		// Process witnesses
		witnessList, ok := Params["witnesses"].([]interface{})
		g.Expect(ok).To(gomega.BeTrue(), "witnesses should be a list")

		WitnessesBytes := make([][]byte, 0)
		for _, w := range witnessList {
			data, err := interfaceToBytes(w)
			g.Expect(err).To(gomega.BeNil(), "witnesses interfaceToBytes failed")
			WitnessesBytes = append(WitnessesBytes, data)
		}

		// 初始化为空切片

		TransactionFailedToResolve := types.Transaction{
			Version:     version32,
			CellDeps:    CellDepBytes,
			HeaderDeps:  headerdeps,
			Inputs:      InputsBytes,
			Outputs:     OutputsBytes,
			OutputsData: OutputDataBytes,
			Witnesses:   WitnessesBytes,
		}

		info, err := client.EstimateCycles(context.Background(), &TransactionFailedToResolve)
		g.Expect(err).To(gomega.BeNil(), "EstimateCycles failed")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		localresultNum, err := interfaceToUint(localresult["cycles"])

		fmt.Println(info.Cycles)
		fmt.Println(uint64(localresultNum))
		g.Expect(info.Cycles).To(gomega.Equal(uint64(localresultNum)))
	})
}
