package rotas

import "net/http"

var rotasUsuarios = []Rota{

	{
		URI:                "/usuarios",
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		Metodo:             http.MethodPost,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios",
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		Metodo:             http.MethodPut,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		Metodo:             http.MethodDelete,
		RequerAutenticacao: false,
	},
}
