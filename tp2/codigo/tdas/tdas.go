package tdas

type Usuario interface {
	LeerNombreDeUsuario() string
	MostrarAfinidad() int
	HayMasPosts() bool
	VerSiguientePost() error
	AgregarAlFeed(Post)
	PublicarPost(int, string) Post
}

type Post interface {
	LeerNombreDelPublicador() string
	MostrarAfinidadDelPublicador() int
	RecibirLike(Usuario)
	MostrarPost() string
	MostrarLikes() (string, error)
}
