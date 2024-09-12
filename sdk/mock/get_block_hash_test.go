package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlockHash(t *testing.T) {

	t.Run("get_block_hash/[block_number]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		Parms, err := interfaceToUint(mockData.Request.Params[0])
		ParmUint64 := uint64(Parms)
		info, err := client.GetBlockHash(context.Background(), ParmUint64)
		g.Expect(err).To(gomega.BeNil(), "GetBlockHash failed")
		fmt.Println(info)
		fmt.Println(mockData.Response.Result)
		g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result.(string)))

	})

	t.Run("get_block_hash/null", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")
		// Identifiable description for the expectation		fmt.Println(mockData.Request.Params)

		Parms, err := interfaceToUint(mockData.Request.Params[0])
		ParmUint64 := uint64(Parms)
		info, err := client.GetBlockHash(context.Background(), ParmUint64)
		g.Expect(err).To(gomega.BeNil(), "GetBlockHash failed")
		fmt.Println(info)
		if mockData.Response.Result == nil {
			mockData.Response.Result = "0x0000000000000000000000000000000000000000000000000000000000000000"
			g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result))
		}
		//g.Expect(info.Hex()).To(gomega.Equal(mockData.Response.Result.(string)))

	})
}
