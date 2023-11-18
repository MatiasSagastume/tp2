package algogram_tdas

type Usuario interface {
	LeerNombreDeUsuario() string
	MostrarAfinidad() int
	HayMasPosts() bool
	VerSiguientePost() (Post, error)
	AgregarAlFeed(Post)
	PublicarPost(int, string) Post
}

type Post interface {
	LeerNombreDelPublicador() string
	MostrarAfinidadDelPublicador() int
	MostrarID() int
	RecibirLike(Usuario)
	MostrarPost() string
	MostrarLikes() (string, error)
}

type DiccionarioUsuarios interface {
	DevolverUsuario(string) (Usuario, error)
	AgregarUsuario(string, Usuario)
	AgregarPost(Post)
	Cantidad() int
}
