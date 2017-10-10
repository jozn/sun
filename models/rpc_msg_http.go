package models

import (
	"github.com/jozn/protobuf/proto"
	"io/ioutil"
	"ms/sun/models/x"
	"net/http"
)

//TODO: add session handlerer/checker
func HttpRpcHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1e6)
	//time.Sleep(2 * time.Second)
	print("called MsgUploadV1: ", r)

	pbFile, _, err := r.FormFile("pb")
	if err != nil {
		print(err.Error())
		return
	}
	defer pbFile.Close()

	pbBytes, err := ioutil.ReadAll(pbFile)
	if err != nil {
		print(err.Error())
		return
	}

	cmdToServer := &x.PB_CommandToServer{}
	err = proto.Unmarshal(pbBytes, cmdToServer)
	if err != nil {
		print(err.Error())
		return
	}
	userParam := &RPC_UserParam_Imple{
		userId: 6,
	}

	x.HandleRpcs(*cmdToServer, userParam, RpcAll)

}
