package mock

import (
	"testing"
)

func TestSendAlert(t *testing.T) {

	t.Run("send_alert/AlertFailedToVerifySignatures", func(t *testing.T) {
		// g := gomega.NewGomegaWithT(t) // Initialize Gomega

		println("Running test case:", t.Name()) // Identifiable marker
		// client, mockdata, err := getMockRpcClientByName(t.Name())
		// g.Expect(err).To(gomega.BeNil(), "getMockRpcClientByName failed") // Identifiable description for the expectation

		// Parms, err := interfaceSliceToMapString(mockdata.Request.Params)
		// cancel, err := interfaceToUint(Parms["cancel"])
		// g.Expect(err).To(gomega.BeNil(), "interfaceToUint cancel failed")
		// id, err := interfaceToUint(Parms["id"])
		// g.Expect(err).To(gomega.BeNil(), "interfaceToUint id failed")
		// notice, err := interfaceToUint(Parms["notice_until"])
		// g.Expect(err).To(gomega.BeNil(), "interfaceToUint notice_until failed")
		// priority, err := interfaceToUint(Parms["priority"])
		// g.Expect(err).To(gomega.BeNil(), "interfaceToUint priority failed")

		// sin := Parms["signatures"].([]string)

		// signatures := make([]json.RawMessage, 0)
		// for _,s := range sin {
		// 	jsonString := fmt.Sprintf("%q", s) // %q 自动为字符串添加引号
		// 	signatures = append(signatures, json.RawMessage(jsonString))
		// }

		// types.Alert
		// alert := types.AlertMessage{
		// 	Cancel:      uint32(cancel),
		// 	Id:          uint32(id),
		// 	Message:     Parms["message"].(string),
		// 	NoticeUntil: uint64(notice),
		// 	Priority:    uint32(priority),
		// 	Signatures:  signatures,
		// }
		// errmsg := client.SendAlert(context.Background(), alert)

		// g.Expect(err).To(gomega.BeNil(), "RemoveTransaction failed")

		// fmt.Println(mockdata.Response.Result)
		// g.Expect(info).To(gomega.Equal(mockdata.Response.Result))
		// Description with marker
	})
}
