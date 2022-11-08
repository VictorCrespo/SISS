package models

type Usuario struct {
	Usuario_id uint   `gorm:"usuario_id"`
	Usuario    string `gorm:"type:varchar(30)"`
	Contrasena string `gorm:"type:varchar(16)"`
	Activo     bool   `gorm:"activo"`
}

type Usuarios []Usuario
