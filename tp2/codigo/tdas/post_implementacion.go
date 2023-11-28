package algogram_tdas

import (
	"algogram/codigo/errores"
	"strconv"
	"strings"
	"tdas/diccionario"
)

const (
	TABULACION             = "	"
	SALTO_DE_LINEA         = "\n"
	CONSTANTE_GUARDAR_POST = 1
)

type postImplementacion struct {
	id         int
	publicador Usuario
	contenido  string
	likes      diccionario.DiccionarioOrdenado[string, int]
}

func (post *postImplementacion) LeerNombreDelPublicador() string {
	return post.publicador.LeerNombreDeUsuario()
}

func (post *postImplementacion) RecibirLike(usuario Usuario) {
	post.likes.Guardar(usuario.LeerNombreDeUsuario(), CONSTANTE_GUARDAR_POST)
}

func (post *postImplementacion) MostrarPost() string {
	var res string
	res += "Post ID " + strconv.Itoa(post.MostrarID()) + SALTO_DE_LINEA
	res += post.LeerNombreDelPublicador() + " dijo: " + post.contenido + SALTO_DE_LINEA
	res += "Likes: " + strconv.Itoa(post.likes.Cantidad())
	return res
}

func (post *postImplementacion) MostrarLikes() string {
	var res string
	if post.likes.Cantidad() == 0 {
		return errores.ErrorPostLikeados{}.Error()
	}
	res = "El post tiene " + strconv.Itoa(post.likes.Cantidad()) + " likes:"
	for iter := post.likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre, _ := iter.VerActual()
		res += SALTO_DE_LINEA + TABULACION + nombre
	}
	return res
}

func (post *postImplementacion) ObtenerAfinidadDelPublicador() int {
	return post.publicador.ObtenerAfinidad()
}
func (post *postImplementacion) MostrarID() int {
	return post.id
}

func CrearPost(id int, usuario Usuario, texto string) Post {
	post := new(postImplementacion)
	post.id = id
	post.publicador = usuario
	post.contenido = texto
	post.likes = diccionario.CrearABB[string, int](strings.Compare)
	return post
}
