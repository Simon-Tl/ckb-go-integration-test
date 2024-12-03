package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

// 传入方法名不一致mock方法get_fee_rate_statics sdk传入get_fee_rate_statistics

func TestGetFeeRateStatics(t *testing.T) {

	t.Run("get_fee_rate_statistics/[null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetFeeRateStatistics(context.Background(), nil)
		g.Expect(err).To(gomega.BeNil(), "GetFeeRateStatistics failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")
		// Description with marker
		LocaResult, err := interfaceToUint(mockResult["mean"])
		g.Expect(err).To(gomega.BeNil(), "LocaResult interfaceToUint failed")

		g.Expect(info.Mean).To(gomega.Equal(uint64(LocaResult)),
			"Result Unequal") // Description added for clarity
	})

	t.Run("get_fee_rate_statistics/[target]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		Params, err := interfaceSliceToStringSlice(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Params interfaceSliceToStringSlice failed")

		ParamsTarget, err := interfaceToUint(Params[0])
		g.Expect(err).To(gomega.BeNil(), "ParamsTarget interfaceToUint failed ")

		info, err := client.GetFeeRateStatistics(context.Background(), uint64(ParamsTarget))
		fmt.Println(info)
		g.Expect(err).To(gomega.BeNil(), "GetFeeRateStatistics failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")
		// Description with marker
		LocaResult, err := interfaceToUint(mockResult["mean"])
		g.Expect(err).To(gomega.BeNil(), "LocaResult interfaceToUint failed")
		fmt.Println(LocaResult)
		g.Expect(info.Mean).To(gomega.Equal(uint64(LocaResult)),
			"Result Unequal") // Description added for clarity
	})

}
