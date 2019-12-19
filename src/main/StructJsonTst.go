package main

import (
	"encoding/json"
	"fmt"
)

// 定义手机屏幕
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

// 定义电池
type Battery struct {
	Capacity int // 容量
}

// 生成json数据
func genJsonData() []byte {
	// 完整数据结构
	raw := &struct { //直接定义结构 再赋值
	/*
	`json:" "` 标签的使用总结为以下几点：
	FieldName int `json:"-"`：表示该字段被本包忽略；
	FieldName int `json:"myName"`：表示该字段在 JSON 里使用“myName”作为键名；
	FieldName int `json:"myName,omitempty"`：表示该字段在 JSON 里使用“myName”作为键名，
						并且如果该字段为空时将其省略掉；
	FieldName int `json:",omitempty"`：该字段在json里的键名使用默认值，但如果该字段为空时会被省略掉，
						注意 omitempty 前面的逗号不能省略。
	 */
		Screen `json:"ScreenJson,omitempty"`//json标签,在json encode时替换key值，去除空值
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},

		// 电池参数
		Battery: Battery{
			2910,
		},

		// 是否有指纹识别
		HasTouchID: true,
	}

	// 将数据序列化为json
	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {

	// 生成一段json数据
	jsonData := genJsonData()

	fmt.Println(string(jsonData))

	// 只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	// 反序列化到screenAndTouch
	json.Unmarshal(jsonData, &screenAndTouch)

	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", screenAndTouch)

	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)

	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)
}
