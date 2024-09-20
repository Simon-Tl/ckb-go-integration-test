package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlockHash(t *testing.T) {

	t.Run("get_block_hash/[block_number]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		mockParams, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "mockParams interfaceToUint failed")

		info, err := client.GetBlockHash(context.Background(), uint64(mockParams))
		g.Expect(err).To(gomega.BeNil(), "GetBlockHash failed")

		g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result.(string)), "Result Unequal")

	})

	t.Run("get_block_hash/null", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		mockParams, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "mockParams interfaceToUint failed")

		info, err := client.GetBlockHash(context.Background(), uint64(mockParams))
		g.Expect(err).To(gomega.BeNil(), "GetBlockHash failed")

		if mockData.Response.Result == nil {
			mockData.Response.Result = "0x0000000000000000000000000000000000000000000000000000000000000000"
			g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result), "Result Unequal")
		}

	})
}
