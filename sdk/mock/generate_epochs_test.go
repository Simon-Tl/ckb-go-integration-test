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
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		numEpoch, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "numEpoch interfaceToUint64 failed")

		info, err := client.GenerateEpochs(context.Background(), uint64(numEpoch))
		g.Expect(err).To(gomega.BeNil(), "GenerateEpochs failed")

		// Description with marker
		mockResult, err := interfaceToUint(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToUint failed")

		g.Expect(info).To(gomega.Equal(uint64(mockResult)), "Result Unequal")

		fmt.Println(info)
		fmt.Println(mockResult)

	})

}
