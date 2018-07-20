package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	return n
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func home(w http.ResponseWriter, req *http.Request) {

	datos, status := req.URL.Query()["number"]

	if !status || len(datos[0]) < 1 {
		log.Println("Sin parámetros recibidos")
	} else {
		dato := datos[0]
		println(dato)
		datoi, err := strconv.Atoi(dato)
		if err != nil {
			println("error al cambiar tipo de dato")
		}
		res2B, _ := json.Marshal(datoi)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(res2B)
	}

}
