package rpc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/nervosnetwork/ckb-sdk-go/v2/mocking"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"io"
	"net/http"
	"testing"
)

type Request struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type Response struct {
	ID      int                    `json:"id"`
	JSONRPC string                 `json:"jsonrpc"`
	Result  map[string]interface{} `json:"result"`
}

type MockData struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func fetchMockData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func TestGetLiveCell(t *testing.T) {
	t.Run("should return msg, when with_data = false", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		Body, err := fetchMockData("https://ckb-rpc-mock-data.vercel.app/test/get_live_cell/[out_point,with_data=false]")
		g.Expect(err).To(gomega.BeNil(), "Expected no error from GetLiveCell, but got one")

		var bodyJson MockData
		err = json.Unmarshal(Body, &bodyJson)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while unmarshalling Body")

		responseByte, err := json.Marshal(bodyJson.Response)
		g.Expect(err).To(gomega.BeNil(), "Expected no error while marshalling response")

		mockClient, err := mocking.DialContext(context.Background(), "mock://ckb")
		g.Expect(err).To(gomega.BeNil(), "Expected no error from DialContext")

		test, err := mockClient.GetRPCJsonMessage("get_live_cell", bodyJson.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Expected no error from GetRPCJsonMessage")

		var result types.CellWithStatus
		mockClient.Expect(test, responseByte)

		err = mockClient.CallContext(context.Background(), &result, "get_live_cell", bodyJson.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "Expected no error from CallContext")

		g.Expect(result.Status).To(gomega.Equal(bodyJson.Response.Result["status"]))
		g.Expect("0x" + hex.EncodeToString(result.Cell.Output.Lock.Args)).To(gomega.Equal(bodyJson.Response.Result["cell"].(map[string]interface{})["output"].(map[string]interface{})["lock"].(map[string]interface{})["args"].(string)))

	})
}
