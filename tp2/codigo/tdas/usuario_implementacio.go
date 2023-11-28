package algogram_tdas

import (
	"algogram/codigo/errores"
	"tdas/cola_prioridad"
)

type usuarioImplementacion struct {
	nombre   string
	posicion int
	feed     cola_prioridad.ColaPrioridad[Post] // Heap de minimos
}

func (usuario *usuarioImplementacion) LeerNombreDeUsuario() string {
	return usuario.nombre
}

func (usuario *usuarioImplementacion) ObtenerAfinidad() int {
	return usuario.posicion
}

func (usuario *usuarioImplementacion) HayMasPosts() bool {
	return !usuario.feed.EstaVacia()
}

func (usuario *usuarioImplementacion) VerSiguientePost() (Post, error) {
	if !usuario.HayMasPosts() {
		return nil, errores.ErrorNoHayPostsOLogueado{}
	}
	return usuario.feed.Desencolar(), nil
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

func CrearUsuario(nombre string, posicion int) Usuario {
	usuario := new(usuarioImplementacion)
	usuario.nombre = nombre
	usuario.posicion = posicion
	usuario.feed = cola_prioridad.CrearHeap[Post](func(a, b Post) int {
		distanciaA := distancia(usuario.posicion, a.ObtenerAfinidadDelPublicador())
		distanciaB := distancia(usuario.posicion, b.ObtenerAfinidadDelPublicador())
		if distanciaA != distanciaB {
			return distanciaB - distanciaA
		}
		return b.MostrarID() - a.MostrarID()
	})
	return usuario
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
