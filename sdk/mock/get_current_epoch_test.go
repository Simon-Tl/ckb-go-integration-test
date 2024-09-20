package mock

import (
	"context"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetCurrentEpoch(t *testing.T) {
	t.Run("get_current_epoch/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		info, err := client.GetCurrentEpoch(context.Background())
		g.Expect(err).To(gomega.BeNil(), "GetCurrentEpoch failed")

		mockResult, err := interfaceToMapString(mockdata.Response.Result)
		g.Expect(err).To(gomega.BeNil(), "mockResult interfaceToMapString failed")

		mockCompactTargetResult, err := interfaceToUint(mockResult["compact_target"])

		g.Expect(err).To(gomega.BeNil(), "mockCompactTargetResult interfaceToMapString failed")
		g.Expect(info.CompactTarget).To(gomega.Equal(uint64(mockCompactTargetResult)), "Result Unequal")
	})

}
