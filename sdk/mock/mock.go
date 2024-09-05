package mock

import (
	"context"
	"encoding/json"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"io"
	"net/http"
	"strings"
)

const BASE_MOCK_URL = "http://127.0.0.1:5000/test/"

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

func getMockRpcClient(url string) (rpc.Client, error) {

	client, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getMockRpcClientByName(testName string) (rpc.Client, MockData, error) {
	// testName -> url
	urlNames := strings.Split(testName, "/")[1:]
	urlName := strings.Join(urlNames, "/")
	mockUrl := BASE_MOCK_URL + urlName
	// client
	client, err := getMockRpcClient(mockUrl)
	if err != nil {
		return nil, MockData{}, err
	}
	// mock data
	mockData, err := fetchMockData(mockUrl)
	if err != nil {
		return client, MockData{}, err
	}
	var bodyJson MockData
	err = json.Unmarshal(mockData, &bodyJson)
	return client, bodyJson, nil
}
