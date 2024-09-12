package mock

import (
	"context"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"strconv"
	"testing"
)

func TestCalculateDaoMaximumWithdraw(t *testing.T) {

	t.Run("calculate_dao_maximum_withdraw/[out_point,kind]", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		pointMap, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "pointMap interfaceSliceToMapString failed ")

		indexuint, err := interfaceToUint(pointMap["index"])
		index32 := uint32(indexuint)

		point := types.OutPoint{
			TxHash: types.HexToHash(pointMap["tx_hash"].(string)),
			Index:  index32,
		}

		info, err := client.CalculateDaoMaximumWithdraw(context.Background(), &point, types.HexToHash(mockData.Request.Params[1].(string)))
		hexString := "0x" + strconv.FormatUint(info, 16)
		fmt.Println(hexString)
		fmt.Println(mockData.Response.Result.(string))
		g.Expect(hexString).To(gomega.Equal(mockData.Response.Result.(string)), "Unequal results")
		// Description with marker
	})

	t.Run("calculate_dao_maximum_withdraw/DaoError", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		client, mockData, err := getMockRpcClientByName(t.Name())
		g.Expect(err).To(gomega.BeNil(), "Expected no error while getting mock RPC client") // Identifiable description for the expectation

		pointMap, err := interfaceSliceToMapString(mockData.Request.Params)
		g.Expect(err).To(gomega.BeNil(), "pointMap interfaceSliceToMapString failed ")

		indexuint, err := interfaceToUint(pointMap["index"])
		index32 := uint32(indexuint)

		point := types.OutPoint{
			TxHash: types.HexToHash(pointMap["tx_hash"].(string)),
			Index:  index32,
		}

		_, errmesg := client.CalculateDaoMaximumWithdraw(context.Background(), &point, types.HexToHash(mockData.Request.Params[1].(string)))
		errMesg, err := interfaceToMapString(mockData.Response.Error)
		g.Expect(err).To(gomega.BeNil(), "errMesg interfaceToMapString failed")
		fmt.Println(errMesg["message"].(string))
		fmt.Println(errmesg.Error())

		g.Expect(errmesg.Error()).To(gomega.Equal(errMesg["message"].(string)))

		// Description with marker
	})
}
