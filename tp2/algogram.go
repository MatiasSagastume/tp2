package main

import (
	"algogram/codigo/errores"
	algogram_tdas "algogram/codigo/tdas"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UBICACION_ARCH_USUARIOS = 0
	LARGO_PARAMETROS        = 1
	ESPACIO                 = " "
)

func main() {
	parametros := os.Args[1:]
	if len(parametros) < LARGO_PARAMETROS {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	rutaUsuarios := parametros[UBICACION_ARCH_USUARIOS]
	diccUsuarios, err := crearDiccionarioUsuarios(rutaUsuarios)
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
		return
	}
	listaDePosts := make([]algogram_tdas.Post, 0, 1)
	lectura(diccUsuarios, listaDePosts)

}

func crearDiccionarioUsuarios(ruta string) (algogram_tdas.DiccionarioUsuarios, error) {
	res := algogram_tdas.CrearDiccionarioDeUsuarios()
	errorGenerico := errores.ErrorLeerArchivo{}
	archLista, err := os.Open(ruta)
	if err != nil {
		return res, errorGenerico
	}
	lector := bufio.NewScanner(archLista)
	for lector.Scan() {
		nombreUsuario := lector.Text()
		res.AgregarUsuario(nombreUsuario, algogram_tdas.CrearUsuario(nombreUsuario, res.Cantidad()))
	}
	return res, nil
}

func lectura(diccUsuarios algogram_tdas.DiccionarioUsuarios, listaDePosts []algogram_tdas.Post) {
	escaner := bufio.NewScanner(os.Stdin)
	var usuarioLogueado algogram_tdas.Usuario
	for escaner.Scan() {
		entrada := escaner.Text()
		palabras := strings.Split(entrada, ESPACIO)
		comando := palabras[0]
		parametroIngresado := strings.Join(palabras[1:], ESPACIO)
		switch {
		case comando == "login":
			if hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorUsuarioLogueado{}.Error())
				continue
			}
			usuario, err := login(parametroIngresado, diccUsuarios)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			usuarioLogueado = usuario
			fmt.Println("Hola", usuario.LeerNombreDeUsuario())

		case comando == "logout":
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorNadieLoggeado{}.Error())
				continue
			}
			usuarioLogueado = nil
			fmt.Println("Adios")

		case comando == "publicar":
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorNadieLoggeado{}.Error())
				continue
			}
			publicarPost(usuarioLogueado, parametroIngresado, &listaDePosts, diccUsuarios)
			fmt.Println("Post publicado")

		case comando == "ver_siguiente_feed":
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorNoHayPostsOLogueado{}.Error())
				continue
			}
			post, err := usuarioLogueado.VerSiguientePost()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(post.MostrarPost())

		case comando == "likear_post":
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorPostLikear{}.Error())
				continue
			}
			err := likearPost(usuarioLogueado, parametroIngresado, listaDePosts)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("Post likeado")

		case comando == "mostrar_likes":
			fmt.Println(mostrarLikes(parametroIngresado, listaDePosts))

		default:
			fmt.Println(comando, ": la opcion ingresada no es valida")
			continue
		}
	}
}

func login(nombre string, diccUsuarios algogram_tdas.DiccionarioUsuarios) (algogram_tdas.Usuario, error) {
	usuario, err := diccUsuarios.DevolverUsuario(nombre)
	return usuario, err
}

func hayAlguienLogueado(usuarioLogueado algogram_tdas.Usuario) bool {
	return usuarioLogueado != nil
}

func publicarPost(usuario algogram_tdas.Usuario, texto string, listaDePosts *[]algogram_tdas.Post, diccUsuarios algogram_tdas.DiccionarioUsuarios) {
	post := usuario.PublicarPost(len(*listaDePosts), texto)
	*listaDePosts = append(*listaDePosts, post)
	diccUsuarios.AgregarPost(post)
}

func likearPost(usuario algogram_tdas.Usuario, id string, listaDePosts []algogram_tdas.Post) error {
	nroId, err := strconv.Atoi(id)
	if err != nil || nroId >= len(listaDePosts) {
		return errores.ErrorPostLikear{}
	}
	listaDePosts[nroId].RecibirLike(usuario)
	return nil
}

func mostrarLikes(id string, listaDePosts []algogram_tdas.Post) string {
	nroId, err := strconv.Atoi(id)
	if err != nil || nroId >= len(listaDePosts) {
		return errores.ErrorPostLikeados{}.Error()
	}
	return listaDePosts[nroId].MostrarLikes()
}
