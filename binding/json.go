package binding

import (
	"encoding/json"
	"errors"
	"net/http"
)

type jsonBinding struct{}

func (j jsonBinding) Name() string {
	return "json"
}

func (j jsonBinding) Bind(r *http.Request, obj interface{}) error {
	if r == nil || obj == nil {
		return errors.New("invalid request or object")
	}
	return json.NewDecoder(r.Body).Decode(obj)
}
