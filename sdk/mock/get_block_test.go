package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlock(t *testing.T) {
	t.Run("get_block/extension2", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())

		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info, err := client.GetBlock(context.Background(), types.HexToHash(mockData.Request.Params[0].(string)))
		g.Expect(err).To(gomega.BeNil(), "client.GetBlock failed")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		localheader, err := interfaceToMapString(localresult["header"])
		g.Expect(err).To(gomega.BeNil(), "localheader interfaceToMapString failed")

		g.Expect(info.Header.Hash.String()).To(gomega.Equal(localheader["hash"].(string)))
		g.Expect(err).To(gomega.BeNil(), "Unequal results")
		fmt.Println(info.Header.Hash.String())
		fmt.Println(localheader["hash"])

	})

	//	// 缺少参数“verbosity，with_cycles”
	//	t.Run("get_block/extension", func(t *testing.T) {
	//		g := gomega.NewGomegaWithT(t)
	//		println("Running test case:", t.Name())
	//
	//		client, mockData, err := getMockRpcClientByName(t.Name())
	//		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")
	//
	//		info, err := client.GetBlock(context.Background(), types.HexToHash(mockData.Request.Params[0].(string)))
	//		g.Expect(err).To(gomega.BeNil(), "client.GetBlock failed")
	//
	//		localresult, err := interfaceToMapString(mockData.Response.Result)
	//		localheader, err := interfaceToMapString(localresult["header"])
	//		g.Expect(err).To(gomega.BeNil(), "localheader interfaceToMapString failed")
	//
	//		g.Expect(info.Header.Hash.String()).To(gomega.Equal(localheader["hash"].(string)))
	//		g.Expect(err).To(gomega.BeNil(), "Unequal results")
	//		fmt.Println(info.Header.Hash.String())
	//		fmt.Println(localheader["hash"])
	//
	//	})
	//
	//}
}
