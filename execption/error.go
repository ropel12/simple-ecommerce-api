package execption

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ropel12/simple-ecommerce-api/helper"
	"github.com/ropel12/simple-ecommerce-api/model"
)

func NotAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"Message":"Not Allowed"}`)
}

func UnAuthorized(w http.ResponseWriter, message string) {
	messages := &model.WebResponseWithMessage{Status: 401,
		Message: message}
	helper.WriteToResponseBody(w, messages, errors.New("t"), http.StatusUnauthorized)
}
