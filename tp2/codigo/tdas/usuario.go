package algogram_tdas

type Usuario interface {

	// LeerNombreDeUsuario devuelve el nombre del usuario
	LeerNombreDeUsuario() string

	// ObtenerAfinidad devuelve la afinidad del usuario
	ObtenerAfinidad() int

	// HayMasPosts devuelve verdadero en caso de que haya al menos un post en el feed para ver, en caso contrario false
	HayMasPosts() bool

	// VerSiguientePost devuelve el proximo post en el feed, en caso de que no hayan más, devuelve el error correspondiente
	VerSiguientePost() (Post, error)

	// AgregarAlFeed agrega el post recibido al feed del usuario
	AgregarAlFeed(Post)

	// PublicarPost recibe el id y el texto del post a crear y devuelve el post creado
	PublicarPost(int, string) Post
}
