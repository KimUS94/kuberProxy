package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var request string
	var response string

	request = string(req.URL.Path)
	request = strings.TrimLeft(request, "/")

	if strings.Compare(req.Method, "POST") == 0 {
		response = postMethodHandler(req)
	} else {
		response = getMethodHandler(request)
	}

	w.Write([]byte(response))
}

func postMethodHandler(req *http.Request) string {
	return "temp"
}

func getMethodHandler(request string) string {
	var response string
	parms := strings.Split(request, "/")

	switch parms[0] {
	case "get_nodes":
		e, err := json.Marshal(kuber.Nodes)
		if err == nil {
			response = string(e)
		}
	default:
		var temp []string
		if len(parms) < 2 {
			response = executeCommad(parms[0], temp...)
		} else if len(parms) >= 2 {
			for i := 1; i < len(parms); i++ {
				temp = append(temp, parms[i])
			}
			response = executeCommad(parms[0], temp...)
		}
	}

	return response
}

// func getHttpRequestBody(r *http.Request) string {
// 	len := r.ContentLength
// 	body := make([]byte, len)
// 	r.Body.Read(body)

// 	return string(body)
// }

func executeCommad(cmdType string, args ...string) string {
	var res string

	cmd := exec.Command(cmdType, args...)
	output, err := cmd.Output()
	if err != nil {
		res = string("err")
	} else {
		res = string(output)
	}

	return res
}
