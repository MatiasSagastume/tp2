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
	UBICACION_COMANDO       = 0
	INDICE_PARAMETROS       = 1
	LOGIN                   = "login"
	LOGOUT                  = "logout"
	SALUDO                  = "Hola"
	MENSAJE_DESPEDIDA       = "Adios"
	PUBLICAR                = "publicar"
	MENSAJE_PUBLICAR_POST   = "Post publicado"
	SIGUIENTE_FEED          = "ver_siguiente_feed"
	LIKEAR                  = "likear_post"
	MENSAJE_POST_LIKEADO    = "Post likeado"
	MOSTRAR_LIKES           = "mostrar_likes"
	OPCION_INVALIDA         = ": la opcion ingresada no es valida"
)

func main() {
	parametros := os.Args[1:]
	if len(parametros) < LARGO_PARAMETROS {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	rutaUsuarios := parametros[UBICACION_ARCH_USUARIOS]
	baseDeDatosUsuarios, err := crearBaseDeDatosUsuarios(rutaUsuarios)
	if err != nil {
		fmt.Println(errores.ErrorLeerArchivo{})
		return
	}
	arrayDePosts := make([]algogram_tdas.Post, 0, 1)
	lecturaInput(baseDeDatosUsuarios, arrayDePosts)

}

func crearBaseDeDatosUsuarios(ruta string) (algogram_tdas.BaseDeDatosUsuarios, error) {
	res := algogram_tdas.CrearBaseDeDatosUsuarios()
	errorGenerico := errores.ErrorLeerArchivo{}
	archUsuarios, err := os.Open(ruta)
	if err != nil {
		return res, errorGenerico
	}
	lector := bufio.NewScanner(archUsuarios)
	for lector.Scan() {
		nombreUsuario := lector.Text()
		res.AgregarUsuario(algogram_tdas.CrearUsuario(nombreUsuario, res.Cantidad()))
	}
	return res, nil
}

func lecturaInput(baseDeDatosUsuarios algogram_tdas.BaseDeDatosUsuarios, arrayDePosts []algogram_tdas.Post) {
	escaner := bufio.NewScanner(os.Stdin)
	var usuarioLogueado algogram_tdas.Usuario
	for escaner.Scan() {
		entrada := escaner.Text()
		palabras := strings.Split(entrada, ESPACIO)
		comando := palabras[UBICACION_COMANDO]
		parametroIngresado := strings.Join(palabras[INDICE_PARAMETROS:], ESPACIO)
		switch {
		case comando == LOGIN:
			if hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorUsuarioLogueado{}.Error())
				continue
			}
			usuario, err := login(parametroIngresado, baseDeDatosUsuarios)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			usuarioLogueado = usuario
			fmt.Println(SALUDO, usuario.LeerNombreDeUsuario())

		case comando == LOGOUT:
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorNadieLoggeado{}.Error())
				continue
			}
			usuarioLogueado = nil
			fmt.Println(MENSAJE_DESPEDIDA)

		case comando == PUBLICAR:
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorNadieLoggeado{}.Error())
				continue
			}
			publicarPost(usuarioLogueado, parametroIngresado, &arrayDePosts, baseDeDatosUsuarios)
			fmt.Println(MENSAJE_PUBLICAR_POST)

		case comando == SIGUIENTE_FEED:
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

		case comando == LIKEAR:
			if !hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorPostLikear{}.Error())
				continue
			}
			err := likearPost(usuarioLogueado, parametroIngresado, arrayDePosts)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(MENSAJE_POST_LIKEADO)

		case comando == MOSTRAR_LIKES:
			fmt.Println(mostrarLikes(parametroIngresado, arrayDePosts))

		default:
			fmt.Println(comando, OPCION_INVALIDA)
			continue
		}
	}
}

func login(nombre string, BaseUsuarios algogram_tdas.BaseDeDatosUsuarios) (algogram_tdas.Usuario, error) {
	usuario, err := BaseUsuarios.DevolverUsuario(nombre)
	return usuario, err
}

func hayAlguienLogueado(usuarioLogueado algogram_tdas.Usuario) bool {
	return usuarioLogueado != nil
}

func publicarPost(usuario algogram_tdas.Usuario, texto string, arrayDePosts *[]algogram_tdas.Post, BaseUsuarios algogram_tdas.BaseDeDatosUsuarios) {
	post := usuario.PublicarPost(len(*arrayDePosts), texto)
	*arrayDePosts = append(*arrayDePosts, post)
	BaseUsuarios.AgregarPost(post)
}

func likearPost(usuario algogram_tdas.Usuario, id string, arrayDePosts []algogram_tdas.Post) error {
	nroId, err := strconv.Atoi(id)
	if err != nil || nroId >= len(arrayDePosts) {
		return errores.ErrorPostLikear{}
	}
	arrayDePosts[nroId].RecibirLike(usuario)
	return nil
}

func mostrarLikes(id string, arrayDePosts []algogram_tdas.Post) string {
	nroId, err := strconv.Atoi(id)
	if err != nil || nroId >= len(arrayDePosts) {
		return errores.ErrorPostLikeados{}.Error()
	}
	return arrayDePosts[nroId].MostrarLikes()
}
