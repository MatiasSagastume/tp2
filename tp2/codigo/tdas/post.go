package algogram_tdas

type Post interface {

	// LeerNombreDelPublicador devuelve el nombre del usuario que publico el post
	LeerNombreDelPublicador() string

	// ObtenerAfinidadDelPublicador devuelve la afinidad del publicador segun como esté definida
	ObtenerAfinidadDelPublicador() int

	// MostrarID devuelve la afinidad del publicador
	MostrarID() int

	// RecibirLike recibe el usuario el cual ha dado like al post y agrega el like al post
	RecibirLike(Usuario)

	// MostrarPost devuelve un string acorde al formato con el cual se debe mostrar el post
	MostrarPost() string

	// MostrarLikes devuelve un string acorde al formato con el cual se debe mostrar los likes de un post,
	// en caso de que no tenga likes devuelve el error correspondiente
	MostrarLikes() string
}
