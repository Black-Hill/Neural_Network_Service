package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type response struct {
	Indicator string
	Results   interface{}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	text := `
	<h2>How to use the API</h2>
	<table>
	<tr>
		<td><b>Method</b></td>
		<td><b>Parameters</b></td>
		<td><b>Example</b></td>
		<td><b>Data</b></td>
	</tr>
	<tr>
		<td>
			Run (GET) Return: Bool
		</td>
		<td>
			<p>/run</p>
		</td>
		<td>
			<p>/run</p>
		</td>
		<td >
			
		</td>
	</tr>
	</table>
	`

	w.Write([]byte("<h1>Welcome to the Neural Network Service.</h1> \n " + text))
}

func RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	/*
	vars := mux.Vars(r)
	if checkRequestVariables(vars) == false {
		log.Warn("[Request][Run] Request missing parameters.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	*/

	log.Info("[Request][Run] Running NN")


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("true"))
}

func checkRequestVariables(parameters map[string]string) bool {
	for _, variable := range parameters {
		if variable == "" {
			return false
		}
	}

	return true
}
