package algogram_tdas

import (
	"algogram/codigo/errores"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
)

type postImplementacion struct {
	id         int
	publicador Usuario
	contenido  string
	likes      diccionario.DiccionarioOrdenado[string, int]
}

type usuarioImplementacion struct {
	nombre     string
	posicion   int
	postActual Post
	feed       cola_prioridad.ColaPrioridad[Post] // Heap de minimos
}

type diccionarioUsuariosImplementacion struct {
	dicc diccionario.Diccionario[string, Usuario]
}

func (usuario *usuarioImplementacion) LeerNombreDeUsuario() string {
	return usuario.nombre
}

func (usuario *usuarioImplementacion) MostrarAfinidad() int {
	return usuario.posicion
}

func (usuario *usuarioImplementacion) HayMasPosts() bool {
	return !usuario.feed.EstaVacia()
}

func (usuario *usuarioImplementacion) VerSiguientePost() error {
	if !usuario.HayMasPosts() {
		return errores.ErrorNoHayPosts{}
	}
	usuario.postActual = usuario.feed.Desencolar()
	return nil
}

func (usuario *usuarioImplementacion) AgregarAlFeed(post Post) {
	if usuario.LeerNombreDeUsuario() == post.LeerNombreDelPublicador() {
		return
	}
	usuario.feed.Encolar(post)
}

func (usuario *usuarioImplementacion) PublicarPost(id int, texto string) Post {
	return CrearPost(id, usuario, texto)
}

func (post *postImplementacion) LeerNombreDelPublicador() string {
	return post.publicador.LeerNombreDeUsuario()
}

func (post *postImplementacion) RecibirLike(usuario Usuario) {
	post.likes.Guardar(usuario.LeerNombreDeUsuario(), 1)
}

func (post *postImplementacion) MostrarPost() string {
	var res string
	res += "Post ID: " + strconv.Itoa(post.id) + "\n"
	res += post.LeerNombreDelPublicador() + " dijo: " + post.contenido + "\n"
	res += "Likes: " + strconv.Itoa(post.likes.Cantidad())
	return res
}

func (post *postImplementacion) MostrarLikes() (string, error) {
	var res string
	if post.likes.Cantidad() == 0 {
		return res, errores.ErrorPostLikeados{}
	}
	res = "El post tiene " + strconv.Itoa(post.likes.Cantidad()) + " likes:\n"
	for iter := post.likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre, _ := iter.VerActual()
		res += nombre + "\n"
	}
	return res, nil
}

func (post *postImplementacion) MostrarAfinidadDelPublicador() int {
	return post.publicador.MostrarAfinidad()
}

func CrearPost(id int, usuario Usuario, texto string) Post {
	post := new(postImplementacion)
	post.id = id
	post.publicador = usuario
	post.contenido = texto
	post.likes = diccionario.CrearABB[string, int](strings.Compare)
	return post
}

func (dic *diccionarioUsuariosImplementacion) DevolverUsuario(nombre string) (Usuario, error) {
	if !dic.dicc.Pertenece(nombre) {
		return nil, errores.ErrorUsuarioNoExiste{}
	}
	return dic.dicc.Obtener(nombre), nil
}

func (dic *diccionarioUsuariosImplementacion) AgregarUsuario(nombre string, usuario Usuario) {
	dic.dicc.Guardar(nombre, usuario)
}

func (dic *diccionarioUsuariosImplementacion) Cantidad() int {
	return dic.dicc.Cantidad()
}

func CrearUsuario(nombre string, posicion int) Usuario {
	usuario := new(usuarioImplementacion)
	usuario.nombre = nombre
	usuario.posicion = posicion
	usuario.feed = cola_prioridad.CrearHeap[Post](func(a, b Post) int {
		distanciaA := distancia(usuario.posicion, a.MostrarAfinidadDelPublicador())
		distanciaB := distancia(usuario.posicion, b.MostrarAfinidadDelPublicador())
		if distanciaA > distanciaB {
			return -1
		}
		if distanciaA < distanciaB {
			return 1
		}
		return 0
	})
	return usuario
}

func CrearDiccionarioDeUsuarios() DiccionarioUsuarios {
	dic := new(diccionarioUsuariosImplementacion)
	dic.dicc = diccionario.CrearHash[string, Usuario]()
	return dic
}

func distancia(x, y int) int {
	return modulo(x - y)
}

func modulo(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
