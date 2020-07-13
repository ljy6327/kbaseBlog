package response

/**
使用建造者设计模式，统一数据返回
*/
type Response struct {
	Success bool
	Count   int
	Msg     string
	Data    interface{}
}

func IsSuccess() Response {
	return Response{Success: true}
}

func IsFalse() Response {
	return Response{Success: false}
}

func (response Response) SetMsg(msg string) Response {
	response.Msg = msg
	return response
}

func (response Response) SetData(data interface{}) Response {
	response.Data = data
	return response
}

func (response Response) SetCount(count int) Response {
	response.Count = count
	return response
}
