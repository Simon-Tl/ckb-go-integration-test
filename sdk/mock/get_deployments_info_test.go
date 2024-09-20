package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

// 返回值和json unmarshal 字段不匹配             json: cannot unmarshal string into Go struct field DeploymentInfo.Deployments.Period of type uint64
func TestGetDeploymentsInfo(t *testing.T) {

	t.Run("get_deployments_info/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetDeploymentsInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetDeploymentsInfo failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		mockEpochResult, err := interfaceToUint(mockResult["epoch"])
		g.Expect(err).To(gomega.BeNil(), "mockEpochResult interfaceToUint failed")
		g.Expect(info.Epoch).To(gomega.Equal(uint64(mockEpochResult)))
		// Description with marker
	})

}
