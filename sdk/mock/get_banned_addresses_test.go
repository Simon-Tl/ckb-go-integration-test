package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBannedAddresses(t *testing.T) {

	t.Run("get_banned_addresses/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		info, err := client.GetBannedAddresses(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetBannedAddresses failed")

		g.Expect(info[0].Address).To(gomega.Equal(mockResult["address"]), "Result Unequal")

	})
}
