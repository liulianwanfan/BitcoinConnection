package models

type RPCrequest struct {
	Id      int64 `json:"id"`
	Method  string `json:"method"`
	Jsonrpc string `json:"jsonrpc"`
	Param   interface{} `json:"param"`
}
