package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestSendTransaction(t *testing.T) {
	t.Run("send_transaction/TransactionFailedToResolve", func(t *testing.T) {
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

		_, errMesg := client.SendTransaction(context.Background(), &tx)
		errmesg, err := interfaceToMapString(mockData.Response.Error)
		//g.Expect(err).To(gomega.BeNil(), "GetBlockEconomicState failed")
		fmt.Println(errmesg["message"].(string))
		fmt.Println(errMesg.Error())
		//g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")
		g.Expect(errMesg.Error()).To(gomega.Equal(errmesg["message"].(string)))

	})

	t.Run("send_transaction/data2", func(t *testing.T) {
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

		info, err := client.SendTransaction(context.Background(), &tx)
		g.Expect(err).To(gomega.BeNil(), "TestTxPoolAccept failed")
		fmt.Println(info)
		fmt.Println(mockData.Response.Result)
		g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result))

	})
}
