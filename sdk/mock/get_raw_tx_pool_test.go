package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

// get_raw_tx_pool 缺少参数“verbose”
func TestGetRawTxPool(t *testing.T) {

	t.Run("get_raw_tx_pool/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetRawTxPool(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetRawTxPool failed")

		mockResult, err := interfaceToMapString(mockdata.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		mockPendingResult, err := interfaceSliceToStringSlice(mockResult["pending"].([]interface{}))
		g.Expect(err).To(gomega.BeNil(), "mockPendingResult interfaceSliceToStringSlice failed")

		g.Expect(info.Pending[0].Hex()).To(gomega.Equal(mockPendingResult[0]), "Result Unequal")
		fmt.Println(mockPendingResult[0])
		fmt.Println(info.Pending[0].Hex())
		// Description with marker
	})
	t.Run("get_raw_tx_pool/[verbose=true]]", func(t *testing.T) {
		println("Running test case:", t.Name()) // Identifiable marker

		// Description with marker
	})

	t.Run("get_raw_tx_pool/[verbose=false]]", func(t *testing.T) {
		println("Running test case:", t.Name()) // Identifiable marker

		// Description with marker
	})
}
