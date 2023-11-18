package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type ErrorUsuarioLogueado struct{}

func (e ErrorUsuarioLogueado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorNadieLoggeado struct{}

func (e ErrorNadieLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorNoHayPostsOLogueado struct{}

func (e ErrorNoHayPostsOLogueado) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorPostLikear struct{}

func (e ErrorPostLikear) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorPostLikeados struct{}

func (e ErrorPostLikeados) Error() string {
	return "Error: Post inexistente o sin likes"
}
