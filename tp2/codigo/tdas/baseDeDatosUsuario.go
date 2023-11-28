package algogram_tdas

type BaseDeDatosUsuarios interface {

	// DevolverUsuario recibe un nombre y devuelve el Usuario con ese nombre, en caso de que no exista devuelve el error correspondiente
	DevolverUsuario(string) (Usuario, error)

	// AgregarUsuario recibe un usuario y lo guarda
	AgregarUsuario(Usuario)

	// AgregarPost recibe un post y lo agrega al feed de los usuarios correspondientes
	AgregarPost(Post)

	// Cantidad devuelve la cantidad de usuarios guardados actualmente
	Cantidad() int
}
