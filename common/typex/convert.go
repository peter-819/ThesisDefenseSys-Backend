package typex

import (
	"TDS-backend/common/errorx"
	"encoding/json"
	"fmt"
)

func Convert(in interface{}, out interface{}) error {
	jsonStr, err := json.Marshal(in)
	if err != nil {
		return errorx.NewDefaultError("class转换成json失败: " + err.Error())
	}
	fmt.Println("parsing json: " + string(jsonStr))
	err = json.Unmarshal(jsonStr, out)
	if err != nil {
		return errorx.NewDefaultError("json转换成class失败: " + err.Error())
	}
	fmt.Println("parsing done: ", out)
	return nil
}
