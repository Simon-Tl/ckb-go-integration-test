package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetConsensus(t *testing.T) {

	t.Run("get_consensus/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info, err := client.GetConsensus(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetTipBlockNumber failed")

		mockResult, err := interfaceToMapString(mockdata.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		mockCellbase_MaturityResult, err := interfaceToUint(mockResult["cellbase_maturity"])

		g.Expect(info.CellbaseMaturity).To(gomega.Equal(uint64(mockCellbase_MaturityResult)), "Result Unequal")
		g.Expect(info.Id).To(gomega.Equal(mockResult["id"].(string)), "Result Unequal")

	})
	//t.Run("get_consensus/empty", func(t *testing.T) {
	//	g := gomega.NewGomegaWithT(t)
	//	println("Running test case:", t.Name())
	//	client, mockdata, err := getMockRpcClientByName(t.Name())
	//	g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")
	//
	//	info, err := client.GetConsensus(context.Background())
	//	g.Expect(err).To(gomega.BeNil(), "GetTipBlockNumber failed")
	//
	//	mockResult, err := interfaceToMapString(mockdata.Response.Result)
	//	g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")
	//
	//	mockCellbase_MaturityResult, err := interfaceToUint(mockResult["cellbase_maturity"])
	//
	//	g.Expect(info.CellbaseMaturity).To(gomega.Equal(uint64(mockCellbase_MaturityResult)), "Result Unequal")
	//	g.Expect(info.Id).To(gomega.Equal(mockResult["id"].(string)), "Result Unequal")
	//
	//})
	//
	//t.Run("get_consensus/null", func(t *testing.T) {
	//	g := gomega.NewGomegaWithT(t)
	//	println("Running test case:", t.Name())
	//	client, mockdata, err := getMockRpcClientByName(t.Name())
	//	g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")
	//
	//	info, err := client.GetConsensus(context.Background())
	//	g.Expect(err).To(gomega.BeNil(), "GetTipBlockNumber failed")
	//
	//	mockResult, err := interfaceToMapString(mockdata.Response.Result)
	//	g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")
	//
	//	mockCellbase_MaturityResult, err := interfaceToUint(mockResult["cellbase_maturity"])
	//
	//	g.Expect(info.CellbaseMaturity).To(gomega.Equal(uint64(mockCellbase_MaturityResult)), "Result Unequal")
	//	g.Expect(info.Id).To(gomega.Equal(mockResult["id"].(string)), "Result Unequal")
	//
	//})

}
