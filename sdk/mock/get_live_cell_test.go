package mock

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetLiveCell(t *testing.T) {
	t.Run("get_live_cell/[out_point,with_data=false]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), nil)

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		mockCellResult, err := interfaceToMapString(mockResult["cell"])
		g.Expect(err).To(gomega.BeNil(), "mockCellResult failed")

		mockCellOutPutResult, err := interfaceToMapString(mockCellResult["output"])

		mockCellResultUint, err := interfaceToUint(mockCellOutPutResult["capacity"])

		g.Expect(info.Cell.Output.Capacity).To(gomega.Equal(uint64(mockCellResultUint)), "Capacity Unequal")
		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

	t.Run("get_live_cell/[out_point,with_data=true]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), nil)
		g.Expect(err).To(gomega.BeNil(), "sdk GetLiveCell failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		mockCellResult, err := interfaceToMapString(mockResult["cell"])
		g.Expect(err).To(gomega.BeNil(), "mockCellResult failed")

		mockCellOutPutResult, err := interfaceToMapString(mockCellResult["output"])

		mockCellResultUint, err := interfaceToUint(mockCellOutPutResult["capacity"])

		g.Expect(info.Cell.Output.Capacity).To(gomega.Equal(uint64(mockCellResultUint)), "Capacity Unequal")
		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

	t.Run("get_live_cell/[out_point,with_data=false,include_tx_pool=false]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		includeTxPool := mockData.Request.Params[2].(bool)

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), &includeTxPool)
		g.Expect(err).To(gomega.BeNil(), "sdk GetLiveCell failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		mockCellResult, err := interfaceToMapString(mockResult["cell"])
		g.Expect(err).To(gomega.BeNil(), "mockCellResult failed")

		mockCellOutPutResult, err := interfaceToMapString(mockCellResult["output"])

		mockCellResultUint, err := interfaceToUint(mockCellOutPutResult["capacity"])

		g.Expect(info.Cell.Output.Capacity).To(gomega.Equal(uint64(mockCellResultUint)), "Capacity Unequal")
		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

	t.Run("get_live_cell/[out_point,with_data=false,include_tx_pool=true]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		includeTxPool := mockData.Request.Params[2].(bool)

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), &includeTxPool)
		g.Expect(err).To(gomega.BeNil(), "sdk GetLiveCell failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		mockCellResult, err := interfaceToMapString(mockResult["cell"])
		g.Expect(err).To(gomega.BeNil(), "mockCellResult failed")

		mockCellOutPutResult, err := interfaceToMapString(mockCellResult["output"])

		mockCellResultUint, err := interfaceToUint(mockCellOutPutResult["capacity"])

		g.Expect(info.Cell.Output.Capacity).To(gomega.Equal(uint64(mockCellResultUint)), "Capacity Unequal")
		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

	t.Run("get_live_cell/[out_point,with_data=true,include_tx_pool=true]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		includeTxPool := mockData.Request.Params[2].(bool)

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), &includeTxPool)
		g.Expect(err).To(gomega.BeNil(), "sdk GetLiveCell failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		mockCellResult, err := interfaceToMapString(mockResult["cell"])
		g.Expect(err).To(gomega.BeNil(), "mockCellResult failed")

		mockCellOutPutResult, err := interfaceToMapString(mockCellResult["output"])

		mockCellResultUint, err := interfaceToUint(mockCellOutPutResult["capacity"])

		g.Expect(info.Cell.Output.Capacity).To(gomega.Equal(uint64(mockCellResultUint)), "Capacity Unequal")
		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

	t.Run("get_live_cell/[out_point,with_data=true,include_tx_pool=false]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		includeTxPool := mockData.Request.Params[2].(bool)

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), &includeTxPool)
		g.Expect(err).To(gomega.BeNil(), "sdk GetLiveCell failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		mockCellResult, err := interfaceToMapString(mockResult["cell"])
		g.Expect(err).To(gomega.BeNil(), "mockCellResult failed")

		mockCellOutPutResult, err := interfaceToMapString(mockCellResult["output"])

		mockCellResultUint, err := interfaceToUint(mockCellOutPutResult["capacity"])

		g.Expect(info.Cell.Output.Capacity).To(gomega.Equal(uint64(mockCellResultUint)), "Capacity Unequal")
		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

	t.Run("get_live_cell/unknown", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		Parms, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Parms failed")
		index, err := interfaceToUint(Parms["index"])
		g.Expect(err).To(gomega.BeNil(), "index failed")
		outPoint := types.OutPoint{
			TxHash: types.HexToHash(Parms["tx_hash"].(string)),
			Index:  uint32(index),
		}

		includeTxPool := mockData.Request.Params[2].(bool)

		info, err := client.GetLiveCell(context.Background(), &outPoint, mockData.Request.Params[1].(bool), &includeTxPool)
		g.Expect(err).To(gomega.BeNil(), "sdk GetLiveCell failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult failed")

		g.Expect(info.Status).To(gomega.Equal(mockResult["status"].(string)), "Status Unequal")

	})

}
