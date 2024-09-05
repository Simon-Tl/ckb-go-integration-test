package mock

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTxPoolInfo(t *testing.T) {

	t.Run("tx_pool_info/[]", func(t *testing.T) {
		println(t.Name())
		client, mockData, err := getMockRpcClientByName(t.Name())
		assert.Nil(t, err)
		info, err := client.TxPoolInfo(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, mockData.Response.Result["tip_hash"], info.TipHash.Hex())
	})
}
