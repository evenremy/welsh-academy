package errorx

import "net/http"

func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case ApiError:
		return int(e.StatusHttp()), e.Data()
	default:
		return http.StatusInternalServerError, nil
	}
}
