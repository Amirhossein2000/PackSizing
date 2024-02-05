package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func apiServe() {
	http.HandleFunc("/", RootRoute)
	http.HandleFunc("/packs", UpdatePacks)
	http.HandleFunc("/calc", CalculatePacks)
	http.ListenAndServe(":8080", nil)
}

func RootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Postman Collection</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				text-align: center;
				margin: 50px;
			}
	
			a {
				display: inline-block;
				padding: 10px 20px;
				background-color: #ff6c37; /* Postman orange color */
				color: #fff;
				text-decoration: none;
				border-radius: 5px;
			}
	
			a:hover {
				background-color: #e85c31; /* Slightly darker shade on hover */
			}
		</style>
	</head>
	<body>
		<a href="https://github.com/Amirhossein2000/PackSizing/blob/main/gymshark.postman_collection.json" target="_blank">Visit Postman Collection</a>
	</body>
	</html>	
	`)
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
	w.Header().Set("Content-Type", "application/json")
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
