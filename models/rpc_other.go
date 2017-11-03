package models

import (
	"ms/sun/models/x"
)

type rpcOther int

func (rpcOther) Echo(param *x.PB_OtherParam_Echo, userParam x.RPC_UserParam) (res x.PB_OtherResponse_Echo, err error) {
	//fmt.Println("Echo defulates: ", res, err)
	res.Text = param.Text
	return
}
