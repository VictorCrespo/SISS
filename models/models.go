package models

type Usuario struct {
	Usuario_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment"`
	Usuario    string `gorm:"type:varchar(30); default:''"`
	Contrasena string `gorm:"type:varchar(16); default:''"`
	Activo     bool   `gorm:"type:tinyint; default:0"`
}

type Carrera struct {
	Carrera_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment"`
	Nombre     string `gorm:"type:varchar(50); default:''"`
	Activa     bool   `gorm:"type:boolean; default:0"`
}

type Alumno struct {
	Alumno_id             uint    `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment"`
	Nombrecompleto        string  `gorm:"type:varchar(50); default:''"`
	Sexo                  string  `gorm:"type:char(1); default:''"`
	Domicilio             string  `gorm:"type:varchar(50); default:''"`
	Telefono              string  `gorm:"type:varchar(25); default:''"`
	Correo_electronico    string  `gorm:"type:varchar(50); default:''"`
	Url_foto              string  `gorm:"type:varchar(50); default:''" `
	Carrera_id            Carrera `gorm:"foreignKey:Carrera_id; type:bigint UNSIGNED; NOT NULL"`
	No_control            string  `gorm:"type:varchar(8); default:0"`
	Periodo               string  `gorm:"type:varchar(50); default:0"`
	Semestre              int     `gorm:"type:tinyint; default:0"`
	Usuario_id            Usuario `gorm:"foreignKey:Usuario_id; type:bigint UNSIGNED; NOT NULL"`
	Porcentaje_creditos_A int     `gorm:"type:tinyint; default:0"`
}

type Usuarios []Usuario
type Carreras []Carrera
type Alumnos []Alumno
