package controllers

import (
	"net/http"
	"api/src/autenticacao"
	"api/src/respostas"
	"io/ioutil"
	"encoding/json"
	"api/src/repositorios"
	"api/src/modelos"
	"api/src/banco"
)

// CriarPublicacao adciona uma nova publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacao tras as publicaçoes que apareceriam no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request){

}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request){

}

// AtualizarPublicacao altera uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request){

}

// DeletarPublicacao exclui uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request){

}