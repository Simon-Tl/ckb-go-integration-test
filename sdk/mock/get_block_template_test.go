package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestGetBlockTemplate(t *testing.T) {

	t.Run("get_block_template/[null,null,null]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetBlockTemplate(context.Background())

		g.Expect(err).To(gomega.BeNil(), "get_block_template failed")

		fmt.Println(info.BytesLimit)
		fmt.Println(mockdata.Response.Result)
		// Description with marker
	})

	t.Run("get_block_template/[]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		info, err := client.GetBlockTemplate(context.Background())

		g.Expect(err).To(gomega.BeNil(), "get_block_template failed")

		fmt.Println(info.BytesLimit)
		fmt.Println(mockdata.Response.Result)
		// Description with marker
	})
}
