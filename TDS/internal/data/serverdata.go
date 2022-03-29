package serverdata

import (
	"TDS-backend/TDS/internal/types"    
	"encoding/json"
    "os"
	"fmt"
	"errors"
)

type DataMgr struct {
	Roles []string
	Routes types.RouteMenuReply
	Headers map[string]types.HeaderReply
}

var dataMgr *DataMgr

func GetServerData() *DataMgr {
	if dataMgr == nil {
		dataMgr = newServerData()
	}

	return dataMgr
}

func loadJsonObj(file string, obj interface{}) {
	filePtr, err := os.Open(file)
	if err != nil {
        fmt.Println("文件打开失败 [Err:%s]", err.Error())
    }
    defer filePtr.Close()
    decoder := json.NewDecoder(filePtr)

	switch obj.(type){
	case *types.RouteMenuReply:
		ptr:=obj.(*types.RouteMenuReply)
		err = decoder.Decode(&ptr)
	case *types.HeaderReply:
		ptr:=obj.(*types.HeaderReply)
		err = decoder.Decode(&ptr)
	default:
		err = errors.New("unknown type")
	}
	if err != nil {
        fmt.Println("解码失败", err.Error())
    } 
}

func newServerData() *DataMgr {
	d := DataMgr{}
	d.Roles = []string{"teacher","student"}
	d.Headers = make(map[string]types.HeaderReply)
	
    loadJsonObj("./internal/data/router/routes.json", &d.Routes)
	for _,role := range d.Roles {
		header := types.HeaderReply{}
		loadJsonObj("./internal/data/router/headers/" + role + ".json", &header)
		d.Headers[role] = header
	}
	fmt.Println(d.Headers)
	return &d
}