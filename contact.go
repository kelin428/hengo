package hengo

type StatusResponseCode int

const (
	StatusOK                  StatusResponseCode = 200
	StatusBadRequest          StatusResponseCode = 400
	StatusNotFound            StatusResponseCode = 404
	StatusInternalServerError StatusResponseCode = 500
)

type H map[string]interface{}

type Response struct {
	Code StatusResponseCode `json:"code"`
	Msg  string             `json:"msg"`
	Data *H                 `json:"data"`
}
