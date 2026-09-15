package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioID}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioID}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{usuarioID}/seguir", // usuarioID é do que vai ser seguido
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.PararDeSeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/seguidores",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/seguindo",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/atualizar-senha",
		Metodo:     http.MethodPost, // Nao é put pq se nao iria mudar o hash se eu mandasse mesmo que fosse e mesma senha (aula 99)
		Funcao:     controllers.AtualizarSenha,
		RequerAuth: true,
	},
}
