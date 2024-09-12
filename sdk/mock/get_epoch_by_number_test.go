package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetEpochByNumber(t *testing.T) {

	t.Run("get_epoch_by_number/[epoch_number]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		number, err := interfaceToUint(mockData.Request.Params[0])
		g.Expect(err).To(gomega.BeNil(), "number parser faild") // Identifiable description for the expectation

		umber64 := uint64(number)

		info, err := client.GetEpochByNumber(context.Background(), umber64)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while fetching GetEpochByNumber")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "localresult parser failed")
		compact_target, err := interfaceToUint(localresult["compact_target"])
		g.Expect(err).To(gomega.BeNil(), "compact_target parser failed")

		fmt.Println(info.CompactTarget)
		fmt.Println(uint64(compact_target))

		g.Expect(info.CompactTarget).To(gomega.Equal(compact_target), "result unequal")
		// Description with marker
	})
}
