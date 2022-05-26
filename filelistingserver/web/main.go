package main

import (
	"fmt"
	"github.com/jonah-lab/goLearn/filelistingserver/filter"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()

}
func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func fileHandler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		err := userError("path must start with /list/")
		fmt.Println(err)
		return err
	}
	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {

		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

func main() {
	http.HandleFunc("/list/", filter.ErrWrapper(fileHandler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
