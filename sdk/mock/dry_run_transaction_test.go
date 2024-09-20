package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestDryRunTransaction(t *testing.T) {
	t.Run("dry_run_transaction/[tx]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Params, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Params interfaceSliceToMapString failed")

		version32, err := parseVersion(Params["version"])
		g.Expect(err).To(gomega.BeNil(), "parseVersion failed")

		CellDepBytes, err := parseCellDeps(Params["cell_deps"])
		g.Expect(err).To(gomega.BeNil(), "parseCellDeps failed")

		headerdeps, err := parseHeaderDeps(Params["header_deps"])
		g.Expect(err).To(gomega.BeNil(), "parseHeaderDeps failed")

		InputsBytes, err := parseInputs(Params["inputs"])
		g.Expect(err).To(gomega.BeNil(), "parseInputs failed")

		OutputsBytes, err := parseOutputs(Params["outputs"])
		g.Expect(err).To(gomega.BeNil(), "parseOutputs failed")

		OutputDataBytes, err := parseOutputData(Params["outputs_data"])
		g.Expect(err).To(gomega.BeNil(), "parseOutputsData failed")

		WitnessesBytes, err := parseWitnesses(Params["witnesses"])
		g.Expect(err).To(gomega.BeNil(), "parseWitnesses failed")

		tx := types.Transaction{
			Version:     version32,
			CellDeps:    CellDepBytes,
			HeaderDeps:  headerdeps,
			Inputs:      InputsBytes,
			Outputs:     OutputsBytes,
			OutputsData: OutputDataBytes,
			Witnesses:   WitnessesBytes,
		}

		info, err := client.DryRunTransaction(context.Background(), &tx)
		g.Expect(err).To(gomega.BeNil(), "DryRunTransaction failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString")

		mockCyclesResult, err := interfaceToUint(mockResult["cycles"])
		g.Expect(err).To(gomega.BeNil(), "mockCyclesResult interfaceToMapString")

		g.Expect(info.Cycles).To(gomega.Equal(uint64(mockCyclesResult)), "Cycles Unequal")

		fmt.Println(info.Cycles, uint64(mockCyclesResult))

	})
}
