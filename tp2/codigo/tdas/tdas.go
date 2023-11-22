package algogram_tdas

type Usuario interface {

	// LeerNombreDeUsuario devuelve el nombre del usuario
	LeerNombreDeUsuario() string

	// MostrarAfinidad devuelve la afinidad del usuario
	MostrarAfinidad() int

	// HayMasPosts devuelve verdadero en caso de que haya al menos un post en el feed para ver, en caso contrario false
	HayMasPosts() bool

	// VerSiguientePost devuelve el proximo post en el feed, en caso de que no hayan más, devuelve el error correspondiente
	VerSiguientePost() (Post, error)

	// AgregarAlFeed agrega el post recibido al feed del usuario
	AgregarAlFeed(Post)

	// PublicarPost recibe el id y el texto del post a crear y devuelve el post creado
	PublicarPost(int, string) Post
}

type Post interface {

	// LeerNombreDelPublicador devuelve el nombre del usuario que publico el post
	LeerNombreDelPublicador() string

	// MostrarAfinidadDelPublicador devuelve la afinidad del publicador segun como esté definida
	MostrarAfinidadDelPublicador() int

	// MostrarID devuelve la afinidad del publicador
	MostrarID() int

	// RecibirLike recibe el usuario el cual ha dado like al post y agrega el like al post
	RecibirLike(Usuario)

	// MostrarPost devuelve un string acorde al formato con el cual se debe mostrar el post
	MostrarPost() string

	// MostrarLikes devuelve un string acorde al formato con el cual se debe mostrar los likes de un post
	MostrarLikes() string
}

type DiccionarioUsuarios interface {

	// DevolverUsuario recibe un nombre y devuelve el Usuario con ese nombre, en caso de que no exista devuelve el error correspondiente
	DevolverUsuario(string) (Usuario, error)

	// AgregarUsuario recibe un usuario y lo guarda
	AgregarUsuario(Usuario)

	// AgregarPost recibe un post y lo agrega al feed de los usuarios correspondientes
	AgregarPost(Post)

	// Cantidad devuelve la cantidad de usuarios guardados actualmente
	Cantidad() int
}
