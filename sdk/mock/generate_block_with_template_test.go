package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestGenerateBlockWithTemplate(t *testing.T) {

	t.Run("generate_block_with_template/[block_template]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockdata, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		// CellBase,err := parseCellbase(MockBlock["cellbase"])

		BlockTemplate, err := parseBlockTemplate(mockdata.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "BlockTemplate Failed")
		info, err := client.GenerateBlockWithTemplate(context.Background(), BlockTemplate)

		g.Expect(err).To(gomega.BeNil(), "BlockTemplate failed")

		fmt.Println(mockdata.Response.Result)
		g.Expect(info.String()).To(gomega.Equal(mockdata.Response.Result))
		// Description with marker
	})
}
