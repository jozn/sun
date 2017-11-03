package models

import (
	"fmt"
	"github.com/jozn/protobuf/proto"
	"io/ioutil"
	"ms/sun/models/x"
	"net/http"
)

//TODO: add session handlerer/checker
func HttpRpcHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1e7)
	r.ParseForm()

	var err error = nil
	//time.Sleep(2 * time.Second)
	fmt.Printf("=====================================\ncalled MsgUploadV1: %v ", r.Form)

	for key, value := range r.Form {
		fmt.Sprintf("++++++ %s - %v \n", key, value)
	}

	pbFile, _, err := r.FormFile("blob")
	if err != nil {
		print(err.Error())
		return
	}
	defer pbFile.Close()

	/*pbStr := r.Form.Get("blob")
	if len(pbStr) == 0 {
		print("len == 0 \n")
		return
	}*/

	//pbBytes := []byte(pbFile)
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

	x.HandleRpcs(*cmdToServer, userParam, RpcAll,rpcResHand)

}
