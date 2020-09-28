package models

/**
公用的用于返回结构体的定义，Data可以表示任意类型
 */

type Result struct {
	COde int `json:"code"`//接口返回状态的类型  1表示成功，0表示失败
	Message string `json:"message"`//接口返回状态对应的描述信息
	Data interface{}`json:"data"` //返回的数据

}
