package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:                "/usuarios",
		Funcao:             controllers.CriarUsuario,
		Metodo:             http.MethodPost,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios",
		Funcao:             controllers.BuscarUsuarios,
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Funcao:             controllers.BuscarUsuario,
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Funcao:             controllers.AtualizarUsuario,
		Metodo:             http.MethodPut,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Funcao:             controllers.DeletarUsuario,
		Metodo:             http.MethodDelete,
		RequerAutenticacao: false,
	},
}
