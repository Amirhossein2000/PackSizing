package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func apiServe() {
	http.HandleFunc("/packs", UpdatePacks)
	http.HandleFunc("/calc", CalculatePacks)
	http.ListenAndServe(":8080", nil)
}

func UpdatePacks(w http.ResponseWriter, r *http.Request) {
	packsRequest := []int{}
	defer r.Body.Close()

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf, &packsRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(packsRequest) < 1 {
		w.WriteHeader(http.StatusNotModified)
	}

	packStorage.update(packsRequest)
	w.WriteHeader(http.StatusAccepted)
}

func CalculatePacks(w http.ResponseWriter, r *http.Request) {
	req := make(map[string]int)
	defer r.Body.Close()

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	count, ok := req["count"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	com, num := calculate(count)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(
		map[string]string{
			"Answer": fmt.Sprintf("%s = %d", com, num),
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
