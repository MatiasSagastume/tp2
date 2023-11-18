package main

import (
	"algogram/codigo/errores"
	algogram_tdas "algogram/codigo/tdas"
	"bufio"
	"fmt"
	"os"
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
		parametrosIngresados := palabras[1:]
		switch {
		case comando == "login":
			if hayAlguienLogueado(usuarioLogueado) {
				fmt.Println(errores.ErrorUsuarioLogueado{}.Error())
				continue
			}
			usuario, err := login(parametrosIngresados, diccUsuarios)
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
			publicarPost(usuarioLogueado, parametrosIngresados, &listaDePosts, diccUsuarios)
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
		}
	}
}

func login(parametros []string, diccUsuarios algogram_tdas.DiccionarioUsuarios) (algogram_tdas.Usuario, error) {
	var usuario algogram_tdas.Usuario
	if len(parametros) < 1 || len(parametros) > 2 {
		return usuario, errores.ErrorUsuarioNoExiste{}
	}
	usuario, err := diccUsuarios.DevolverUsuario(parametros[0])
	return usuario, err
}

func hayAlguienLogueado(usuarioLogueado algogram_tdas.Usuario) bool {
	return usuarioLogueado != nil
}

func publicarPost(usuario algogram_tdas.Usuario, texto []string, listaDePosts *[]algogram_tdas.Post, diccUsuarios algogram_tdas.DiccionarioUsuarios) {
	post := usuario.PublicarPost(len(*listaDePosts), strings.Join(texto, ESPACIO))
	*listaDePosts = append(*listaDePosts, post)
	diccUsuarios.AgregarPost(post)
}
