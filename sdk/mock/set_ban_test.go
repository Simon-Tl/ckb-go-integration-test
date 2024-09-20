package mock

import (
	"context"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func TestSetBan(t *testing.T) {

	t.Run("set_ban/[address,command,ban_time,absolute,reason]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed")

		address := mockData.Request.Params[0].(string)
		command := mockData.Request.Params[1].(string)
		banTime, err := interfaceToUint(mockData.Request.Params[2])
		g.Expect(err).To(gomega.BeNil(), "banTime interfaceToUint failed")
		absolute := mockData.Request.Params[3].(bool)
		reason := mockData.Request.Params[4].(string)

		info := client.SetBan(context.Background(), address, command, uint64(banTime), absolute, reason)
		g.Expect(info).To(gomega.BeNil(), "SetBan failed")

		fmt.Println(info)
		// Description with marker

	})
}
