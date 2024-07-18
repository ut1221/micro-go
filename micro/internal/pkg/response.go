package pkg

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http"
	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
	stdhttp "net/http"
)

// httpResponse 响应结构体
type httpResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// EncoderError 错误响应封装
func EncoderError() http.EncodeErrorFunc {
	return func(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
		if err == nil {
			return
		}
		se := &httpResponse{}
		if gs, ok := status.FromError(err); ok {
			se.Code = httpstatus.FromGRPCCode(gs.Code())
			se.Message = gs.Message()
		} else {
			se.Code = stdhttp.StatusInternalServerError
			se.Message = gs.Message()
		}
		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(se)
		if err != nil {
			w.WriteHeader(stdhttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/"+codec.Name())
		w.WriteHeader(stdhttp.StatusOK)
		_, _ = w.Write(body)
	}
}

// EncoderResponse  请求响应封装
func EncoderResponse() http.EncodeResponseFunc {
	return func(w stdhttp.ResponseWriter, request *stdhttp.Request, i interface{}) error {
		if i == nil {
			return nil
		}
		// 假设传入的 `i` 是一个 map 包含所有必要的数据
		dataMap, ok := i.(map[string]interface{})
		if !ok {
			return stdhttp.ErrContentLength
		}
		resp := &httpResponse{
			Code:    stdhttp.StatusOK,
			Message: "",
			Data:    dataMap,
		}
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(resp)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}
