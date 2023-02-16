package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type DataRequest struct {
	DatosSplit string `json:"manejo"`
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT,DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func calculates(w http.ResponseWriter, r *http.Request) {
	var datas DataRequest
	var Envio string
	indexHandler(w, r)
	fmt.Println("Entramos ptaaaaas")
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	nuevo := datas.DatosSplit
	if strings.Contains(nuevo, "+") {
		fmt.Println("Suma")
		delimitador := "+"
		ComoArreglo := strings.Split(nuevo, delimitador)
		for _, nombre := range ComoArreglo {
			fmt.Println(nombre)
		}
		Derecha, _ := strconv.ParseFloat(ComoArreglo[0], 32)
		Izquierda, _ := strconv.ParseFloat(ComoArreglo[1], 32)
		Previa := Derecha + Izquierda
		Envio = strconv.FormatFloat(Previa, 'g', 5, 64)
	} else if strings.Contains(nuevo, "-") {
		fmt.Println("Resta")
		delimitador := "-"
		ComoArreglo := strings.Split(nuevo, delimitador)
		fmt.Printf("Datatype of i : %T\n", ComoArreglo[0])
		for _, nombre := range ComoArreglo {
			fmt.Println(nombre)
		}
		Derecha, _ := strconv.ParseFloat(ComoArreglo[0], 32)
		Izquierda, _ := strconv.ParseFloat(ComoArreglo[1], 32)
		Previa := Derecha - Izquierda
		Envio = strconv.FormatFloat(Previa, 'g', 5, 64)
	} else if strings.Contains(nuevo, "*") {
		fmt.Println("Multiplicacion")
		delimitador := "*"
		ComoArreglo := strings.Split(nuevo, delimitador)
		for _, nombre := range ComoArreglo {
			fmt.Println(nombre)
		}
		Derecha, _ := strconv.ParseFloat(ComoArreglo[0], 32)
		Izquierda, _ := strconv.ParseFloat(ComoArreglo[1], 32)
		Previa := Derecha * Izquierda
		Envio = strconv.FormatFloat(Previa, 'g', 5, 64)
	} else if strings.Contains(nuevo, "/") {
		fmt.Println("Division")
		delimitador := "/"
		ComoArreglo := strings.Split(nuevo, delimitador)
		Derecha, _ := strconv.ParseFloat(ComoArreglo[0], 32)
		Izquierda, _ := strconv.ParseFloat(ComoArreglo[1], 32)
		fmt.Println(Izquierda)
		if Izquierda == 0 && Derecha == 0 {
			Envio = strconv.Quote("Syntax Error")
		} else if Izquierda == 0 {
			Envio = strconv.Quote("Syntax Error")
		} else {
			Previa := Derecha / Izquierda
			Envio = strconv.FormatFloat(Previa, 'g', 5, 64)
		}
	} else {
		Envio = nuevo
		fmt.Println("Si viene solo un numero :3")
	}
	fmt.Println("Que estoy mandando")
	fmt.Println(Envio)
	fmt.Println(datas)
	json.NewEncoder(w).Encode("{\"Resultado\":" + Envio + "}")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	indexHandler(w, r)
	fmt.Println("Entramos ptaaaaa")

	json.NewEncoder(w).Encode("{\"Message\" : \"Hola Mundo \"}")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleIndex).Methods("GET")
	r.HandleFunc("/calcular", calculates).Methods("POST")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	fmt.Println("Listening...")
	//log.Println("Listening...")
	srv.ListenAndServe()
}
