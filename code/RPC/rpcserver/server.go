package rpcserver

import (
	"log"
	"rpc/data"
)

type Calc struct {
}

func (c *Calc) Add(request *data.CalcRequest, response *data.CalcResponse) error {
	log.Println("call add method")
	response.Result = request.Left + request.Right
	return nil
}
