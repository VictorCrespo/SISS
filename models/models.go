package models

type Usuario struct {
	Usuario_id uint   `gorm:"PRIMARY_KEY"`
	Usuario    string `gorm:"type:varchar(30); default ''"`
	Contrasena string `gorm:"type:varchar(16); default ''"`
	Activo     bool   `gorm:"type:tinyint; default 0"`
}

type Alumno struct {
	Alumno_id             uint   `gorm:"PRIMARY_KEY"`
	Nombrecompleto        string `gorm:"type:varchar(50); default ''"`
	Sexo                  string `gorm:"type:char(1); default:''"`
	Domicilio             string `gorm:"type:varchar(50); default ''"`
	Telefono              string `gorm:"type:varchar(25); default ''"`
	Correo_electronico    string `gorm:"type:varchar(50); default ''"`
	Url_foto              string `gorm:"type:varchar(50); default ''" `
	Carrera_id            uint   `gorm:"type:bigint UNSIGNED; NOT NULL"`
	No_control            string `gorm:"type:varchar(8); default 0"`
	Periodo               string `gorm:"type:varchar(50); default 0"`
	Semestre              int    `gorm:"type:tinyint; default 0"`
	Usuario_id            uint   `gorm:"type:bigint UNSIGNED; NOT NULL"`
	Porcentaje_creditos_A int    `gorm:"type:tinyint; default 0"`
}

type Usuarios []Usuario
type Alumnos []Alumno
