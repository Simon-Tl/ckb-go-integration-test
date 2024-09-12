package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetDeploymentsInfo(t *testing.T) {

	t.Run("get_deployments_info/empty", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.GetDeploymentsInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetDeploymentsInfo")

		localresult, _ := interfaceToMapString(mockData.Response.Result)
		localresultNum, _ := interfaceToUint(localresult["epoch"])
		g.Expect(info.Epoch).To(gomega.Equal(uint64(localresultNum)))
		// Description with marker
	})

	t.Run("get_deployments_info/null", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation
		info, err := client.GetDeploymentsInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetDeploymentsInfo")

		localresult, _ := interfaceToMapString(mockData.Response.Result)
		localresultNum, _ := interfaceToUint(localresult["epoch"])
		g.Expect(info.Epoch).To(gomega.Equal(uint64(localresultNum)))
		// Description with marker
	})

	t.Run("get_deployments_info/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())

		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		info, err := client.GetDeploymentsInfo(context.Background())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetDeploymentsInfo")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "localresult parser failed")

		localresultNum, err := interfaceToUint(localresult["epoch"])
		g.Expect(err).To(gomega.BeNil(), "localresultNum parser failed")
		g.Expect(info.Epoch).To(gomega.Equal(uint64(localresultNum)))
		// Description with marker
	})

}
