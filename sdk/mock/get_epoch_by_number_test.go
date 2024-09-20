package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetEpochByNumber(t *testing.T) {

	t.Run("get_epoch_by_number/[epoch_number]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		mockParams, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "number parser faild") // Identifiable description for the expectation

		info, err := client.GetEpochByNumber(context.Background(), uint64(mockParams))
		g.Expect(err).To(gomega.BeNil(), "GetEpochByNumber failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		mockCompactTargetResult, err := interfaceToUint(mockResult["compact_target"])
		g.Expect(err).To(gomega.BeNil(), "mockCompactTargetResult  interfaceToUint failed")

		fmt.Println(info.CompactTarget)
		fmt.Println(uint64(mockCompactTargetResult))

		g.Expect(info.CompactTarget).To(gomega.Equal(uint64(mockCompactTargetResult)), "Result Unequal")
		// Description with marker
	})
}
