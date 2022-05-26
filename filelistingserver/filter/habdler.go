package filter

import (
	"fmt"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func ErrWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic:%s", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		if userErr, ok := err.(userError); ok {
			fmt.Println("userErr", userErr)
			http.Error(writer, userErr.Message(), http.StatusBadRequest)
			return
		}
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}

}
