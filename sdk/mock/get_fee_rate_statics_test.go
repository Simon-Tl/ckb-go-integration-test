package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

// 传入方法名不一致mock方法get_fee_rate_statics sdk传入get_fee_rate_statistics

func TestGetFeeRateStatics(t *testing.T) {

	t.Run("get_fee_rate_statics/[null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetFeeRateStatistics(context.Background(), nil)
		g.Expect(err).To(gomega.BeNil(), "GetFeeRateStatistics failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")
		// Description with marker
		g.Expect(info.Mean).To(gomega.Equal(mockResult["mean"]),
			"Result Unequal") // Description added for clarity
	})
}
