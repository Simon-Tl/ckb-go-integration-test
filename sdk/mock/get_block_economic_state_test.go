package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetBlockEconomicState(t *testing.T) {
	t.Run("get_block_economic_state/[block_hash]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		hash := types.HexToHash(mockData.Request.Params[0].(string))
		info, err := client.GetBlockEconomicState(context.Background(), hash)
		g.Expect(err).To(gomega.BeNil(), "GetBlockEconomicState failed")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "interfaceToMapString failed")

		g.Expect(info.FinalizedAt.String()).To(gomega.Equal(localresult["finalized_at"]))
		fmt.Println(info.FinalizedAt)
		fmt.Println(localresult["finalized_at"])

	})

}
