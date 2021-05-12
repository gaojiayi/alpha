package services

import (
	"alpha-core/src/common"
	"alpha-core/src/config"
	"alpha-core/src/views"
	"encoding/json"
	"strings"
)

/**
处理图像的业务逻辑层
 */

func ImageDetect(fileName string, fileSaveName string, ch chan views.ImageDetectItem) {
	// 1 获取保存文件名
	// 2 执行底层python 输出文件
	var execResult string
	output := common.ExecPythonCommand([]string{config.Detect_script_path, "--imageName", fileSaveName})
	index := strings.Index(output, "data:")
	if index >= 0 {
		execResult = string(output[index+5:])
	}
	// 3 解析结果返回
	var mapArray []map[string]string
	json.Unmarshal([]byte(string(execResult)), &mapArray)

	var resp views.ImageDetectItem
	var detects []views.ImageDetect
	for _, detectItem := range mapArray {
		var imageDetect = views.ImageDetect{}
		imageDetect.Name = detectItem["object"]
		imageDetect.Url = "http://localhost/alpha-static/output/" + detectItem["path"]
		imageDetect.Percet = detectItem["percet"]
		detects = append(detects, imageDetect)
	}
	// 设置默认值
	if detects == nil{
		var imageDetect = views.ImageDetect{}
		imageDetect.Name = "未知图片"
		imageDetect.Url = "http://localhost/alpha-static/unknown.jpg"
		imageDetect.Percet = ""
		detects = append(detects, imageDetect)
	}
	resp.FileName = fileName
	resp.ImageDetectObjects = detects
	ch <- resp
}

func ImageIdentify(fileName string, fileSaveName string, ch chan views.ImageIdentifyResp) {
	args := []string{config.Identify_script_path, "--imageName", fileSaveName, "--modelType", config.Identify_model_type}
	output := common.ExecPythonCommand(args)
	index := strings.Index(output, "data:")
	var execResult string
	if index >= 0 {
		execResult = string(output[index+5:])
	}
	var ret map[string]string
	json.Unmarshal([]byte(string(execResult)), &ret)
	resp :=views.ImageIdentifyResp{}
	var identifyRetArray []views.IdentifyRet
	for key := range ret {
		identifyRet := views.IdentifyRet{}
		identifyRet.Name = key
		identifyRet.Percet = ret[key]
		identifyRetArray = append(identifyRetArray,identifyRet)
	}
	resp.FileName = fileName
	resp.IdentifyRets = identifyRetArray
	ch <- resp

}
