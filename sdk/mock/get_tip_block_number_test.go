package mock

import (
	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	"testing"
)

func TestGetTipBlockNumber(t *testing.T) {

	t.Run("get_tip_block_number/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info, err := client.GetTipBlockNumber(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetTipBlockNumber failed")

		mockResult, err := interfaceToUint(mockdata.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToUint failed")

		g.Expect(info).To(gomega.Equal(uint64(mockResult)), "Result Unequal")
	})

}
