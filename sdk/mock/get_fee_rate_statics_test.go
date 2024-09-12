package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetFeeRateStatics(t *testing.T) {

	t.Run("get_fee_rate_statics/[null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.GetFeeRateStatistics(context.Background(), nil)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetFeeRateStatistics")
		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")
		// Description with marker
		g.Expect(info.Mean).To(gomega.Equal(localresult["mean"]),
			"Expected TipHash to match mock data") // Description added for clarity
	})
}
