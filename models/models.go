package models

type Usuario struct {
	UsuarioID  uint   `gorm:"usuarioID"`
	Usuario    string `gorm:"usuario"`
	Contrasena string `gorm:"contrasena"`
	Activo     bool   `gorm:"activo"`
}

type Usuarios []Usuario
