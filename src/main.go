package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID       string    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Telefone string    `json:"telefone,omitempty"`
	Endereco *Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Cidade   string `json:"cidade,omitempty"`
	Endereco string `json:"endereco,omitempty"`
}

var clientes []Cliente

func GetClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range clientes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Cliente{})
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(clientes)
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clientes {
		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(clientes)
}

func main() {

	router := mux.NewRouter()

	clientes = append(clientes,
		Cliente{
			ID:       "1",
			Nome:     "Quitanda do Seu Ze",
			Telefone: "(32)9999-9999",
			Endereco: &Endereco{
				Cidade:   "JF",
				Endereco: "Rua Francisco Faria, 20, Bairu"}})

	clientes = append(clientes,
		Cliente{
			ID:       "2",
			Nome:     "Padaria da Maria",
			Telefone: "(32)8888-8888",
			Endereco: &Endereco{
				Cidade:   "JF",
				Endereco: "Avenida Governador Valadares, 125, Manoel Honorio"}})

	clientes = append(clientes,
		Cliente{
			ID:       "3",
			Nome:     "Quitutes da Fulaninha",
			Telefone: "(32)1111-1111"})

	router.HandleFunc("/clientes", GetClientes).Methods("GET")
	router.HandleFunc("/cliente/{id}", GetCliente).Methods("GET")
	router.HandleFunc("/cliente", CreateCliente).Methods("POST")
	router.HandleFunc("/cliente/{id}", DeleteCliente).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
