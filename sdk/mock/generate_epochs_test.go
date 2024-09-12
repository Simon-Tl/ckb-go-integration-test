package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGenerateEpochs(t *testing.T) {
	t.Run("generate_epochs/[num_epochs]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client")

		numEpoch, err := interfaceToUint(mockData.Request.Params[0])
		var numEpoch64 uint64 = uint64(numEpoch)
		g.Expect(err).To(gomega.BeNil(), "numEpoch interfaceToUint64 failed")

		info, err := client.GenerateEpochs(context.Background(), numEpoch64)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetPeers")

		// Description with marker
		localresult, err := interfaceToUint(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "localresult interfaceToUint failed")
		localresult64 := uint64(localresult)
		g.Expect(info).To(gomega.Equal(localresult64))
		g.Expect(err).To(gomega.BeNil(), "result unEqual")

		fmt.Println(info)
		fmt.Println(localresult)

	})

}
