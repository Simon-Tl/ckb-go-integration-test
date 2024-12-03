package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
)

//sync_state 缺少“tip_number，unverified_tip_number，tip_hash，unverified_tip_hash”

func TestSyncState(t *testing.T) {
	t.Run("sync_state/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())

		client, mockData, err := getMockRpcClientByName(t.Name())

		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info, err := client.SyncState(context.Background())
		fmt.Println(info.TipNumber, info.UnverifiedTipNumber, info.TipHash, info.UnverifiedTipHash)
		g.Expect(err).To(gomega.BeNil(), "client.GetBlock failed")

		mockResult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		g.Expect(info.AssumeValidTarget.String()).To(gomega.Equal(mockResult["assume_valid_target"].(string)))
		g.Expect(err).To(gomega.BeNil(), "Unequal results")

		structKeys, err := getStructJSONKeys(types.SyncState{})
		err = compareKeys(mockResult, structKeys)
		g.Expect(err).To(gomega.BeNil(), "")

	})
}
