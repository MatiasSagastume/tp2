package main

import (
	"algogram/codigo/errores"
	algogram_tdas "algogram/codigo/tdas"
	"bufio"
	"fmt"
	"os"
)

const (
	UBICACION_ARCH_USUARIOS = 0
	LARGO_PARAMETROS        = 1
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
	}
	listaDePosts := make([]algogram_tdas.Post, 0, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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
	for escaner.Scan() {
	}
}
