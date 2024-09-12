package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBannedAddresses(t *testing.T) {

	t.Run("get_banned_addresses/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		localresult, err := interfaceToMapString(mockData.Response.Result)
		info, err := client.GetBannedAddresses(context.Background())
		fmt.Println(info[0].Address)
		fmt.Println(localresult)
		g.Expect(info[0].Address).To(gomega.Equal(localresult["address"]))

	})
}
