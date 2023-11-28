package algogram_tdas

import (
	"algogram/codigo/errores"
	"tdas/diccionario"
)

type BaseDeDatosUsuariosImplementacion struct {
	dicc diccionario.Diccionario[string, Usuario]
}

func (dic *BaseDeDatosUsuariosImplementacion) DevolverUsuario(nombre string) (Usuario, error) {
	if !dic.dicc.Pertenece(nombre) {
		return nil, errores.ErrorUsuarioNoExiste{}
	}
	return dic.dicc.Obtener(nombre), nil
}

func (dic *BaseDeDatosUsuariosImplementacion) AgregarUsuario(usuario Usuario) {
	dic.dicc.Guardar(usuario.LeerNombreDeUsuario(), usuario)
}

func (dic *BaseDeDatosUsuariosImplementacion) Cantidad() int {
	return dic.dicc.Cantidad()
}

func (dic *BaseDeDatosUsuariosImplementacion) AgregarPost(post Post) {
	dic.dicc.Iterar(func(clave string, dato Usuario) bool {
		if clave != post.LeerNombreDelPublicador() {
			dato.AgregarAlFeed(post)
		}
		return true
	})
}

func CrearBaseDeDatosUsuarios() BaseDeDatosUsuarios {
	dic := new(BaseDeDatosUsuariosImplementacion)
	dic.dicc = diccionario.CrearHash[string, Usuario]()
	return dic
}
