package mock

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestSyncState(t *testing.T) {
	t.Run("sync_state/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		println("Running test case:", t.Name())

		client, mockData, err := getMockRpcClientByName(t.Name())

		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info, err := client.SyncState(context.Background())
		g.Expect(err).To(gomega.BeNil(), "client.GetBlock failed")

		localresult, err := interfaceToMapString(mockData.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "localresult interfaceToMapString failed")

		g.Expect(info.AssumeValidTarget.String()).To(gomega.Equal(localresult["assume_valid_target"].(string)))
		g.Expect(err).To(gomega.BeNil(), "Unequal results")

		structKeys, err := getStructJSONKeys(types.SyncState{})
		err = compareKeys(localresult, structKeys)
		g.Expect(err).To(gomega.BeNil(), "")

	})
}
