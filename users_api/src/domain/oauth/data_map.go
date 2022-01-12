package oauth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
)

type OauthDataMap map[string]string

func (oauthmap *OauthDataMap) SetUuid(uuid string) {
	(*oauthmap)["uuid"] = uuid
}

func (oauthmap *OauthDataMap) SetStatus(code string) {
	(*oauthmap)["status"] = code
}

/*
Status
0: no error 200
1: unathenticated 401
2: permission denied 403
3: internal 500
*/

func GetError(status string) server_message.Svr_message {
	switch status {
	case "0":
		return nil
	case "1":
		return server_message.NewCustomMessage(http.StatusUnauthorized, "authentication invalid")
	case "2":
		return server_message.NewCustomMessage(http.StatusForbidden, "request rejected")
	case "3":
		return server_message.NewInternalError()
	default:
		logger.Error("unkown status code", errors.New(fmt.Sprint("status code is ", status)))
		return server_message.NewInternalError()
	}
}
