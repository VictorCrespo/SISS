package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Actividad struct {
	Actividad_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Descripcion  string `gorm:"type:varchar(50); default:''" json:"descripcion"`
	Activo       bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Alumno_Programa struct {
	Alumno_id   uint   `gorm:"Primary_Key; autoIncrement:false" json:"alumno_id"`
	Programa_id uint   `gorm:"Primary_Key; autoIncrement:false" json:"programa_id"`
	Objetivo    string `gorm:"type:varchar(50); default:''" json:"objetivo"`
	Alumno      Alumno
	Programa    Programa
}

type Alumno struct {
	Alumno_id             uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Nombrecompleto        string `gorm:"type:varchar(50); default:''" json:"nombrecompleto"`
	Sexo                  string `gorm:"type:char(1); default:''" json:"sexo"`
	Domicilio             string `gorm:"type:varchar(50); default:''" json:"domicilio"`
	Telefono              string `gorm:"type:varchar(25); default:''" json:"telefono"`
	Correo_electronico    string `gorm:"type:varchar(50); default:''" json:"correo_electronico"`
	Url_foto              string `gorm:"type:varchar(50); default:''" json:"url_foto"`
	Carrera_id            uint   `gorm:"type:bigint UNSIGNED; NOT NULL" json:"carrera_id"`
	No_control            string `gorm:"type:varchar(8); default:0" json:"no_control"`
	Periodo               string `gorm:"type:varchar(50); default:''" json:"periodo"`
	Semestre              int    `gorm:"type:tinyint; default:0" json:"semestre"`
	Usuario_id            uint   `gorm:"type:bigint UNSIGNED; NOT NULL" json:"usuario_id"`
	Carrera               Carrera
	Usuario               Usuario
	Porcentaje_creditos_a int `gorm:"type:tinyint; default:0" json:"porcentaje_creditos_a"`
}

type Carrera struct {
	Carrera_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Nombre     string `gorm:"type:varchar(50); default:''" json:"nombre"`
	Activo     bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Claim struct {
	Usuario `json:"usuario"`
	jwt.RegisteredClaims
}

type Control_Expendiente struct {
	Control_expedientes_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Solicitud              bool   `gorm:"type:boolean; default:0" json:"solicitud"`
	Curso_induccion        bool   `gorm:"type:boolean; default:0" json:"curso_induccion"`
	Carta_aprobacion       bool   `gorm:"type:boolean; default:0" json:"carta_aprobacion"`
	Plan_trabajo           bool   `gorm:"type:boolean; default:0" json:"plan_trabajo"`
	Reporte_bimestral_p    bool   `gorm:"type:boolean; default:0" json:"reporte_bimestral_p"`
	Reporte_bimestral_s    bool   `gorm:"type:boolean; default:0" json:"reporte_bimestral_s"`
	Reporte_bimestral_t    bool   `gorm:"type:boolean; default:0" json:"reporte_bimestral_t"`
	Reporte_final          bool   `gorm:"type:boolean; default:0" json:"reporte_final"`
	Carta_terminacion      bool   `gorm:"type:boolean; default:0" json:"carta_terminacion"`
	Constancia_oficial     bool   `gorm:"type:boolean; default:0" json:"constancia_oficial"`
	Observaciones          string `gorm:"type:varchar(100); default:''" json:"observaciones"`
	Alumno_id              uint   `gorm:"type:bigint UNSIGNED; NOT NULL" json:"alumno_id"`
	Alumno                 Alumno
}

type Dependencia struct {
	Dependencia_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Nombre         string `gorm:"type:varchar(50); default:''" json:"nombre"`
	Titular        string `gorm:"type:varchar(50); default:''" json:"titular"`
	Puesto         string `gorm:"type:varchar(50); default:''" json:"puesto"`
	Domicilio      string `gorm:"type:varchar(75); default:''" json:"domicilio"`
	Activo         bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Modalidad struct {
	Modalidad_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Nombre       string `gorm:"type:varchar(50); default:''" json:"nombre"`
	Activo       bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Permiso struct {
	Permiso_id  uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Descripcion string `gorm:"type:varchar(50); default:''" json:"descripcion"`
	Activo      bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Rol struct {
	Rol_id      uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Descripcion string `gorm:"type:varchar(50); default:''" json:"descripcion"`
	Activo      bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Programa struct {
	Programa_id      uint      `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Nombre           string    `gorm:"type:varchar(50); default:''" json:"nombre"`
	Capacidad        int       `gorm:"type:int; default:0" json:"capacidad"`
	Activo           bool      `gorm:"type:boolean; default:0" json:"activo"`
	Dependencia_id   uint      `gorm:"type:bigint UNSIGNED; NOT NULL" json:"dependencia_id"`
	Tipo_programa_id uint      `gorm:"type:bigint UNSIGNED; NOT NULL" json:"tipo_programa_id"`
	Modalidad_id     uint      `gorm:"type:bigint UNSIGNED; NOT NULL" json:"modalidad_id"`
	Actividad_id     uint      `gorm:"type:bigint UNSIGNED; NOT NULL" json:"actividad_id"`
	Fecha_inicio     time.Time `gorm:"type:date;" json:"fecha_inicio"`
	Fecha_fin        time.Time `gorm:"type:date;" json:"fecha_fin"`
	Dependencia      Dependencia
	Tipo_Programa    Tipo_Programa
	Modalidad        Modalidad
	Actividad        Actividad
}

type ResponseToken struct {
	Token string `json:"token"`
}

type Rol_Permiso struct {
	Rol_id     uint `gorm:"Primary_Key; autoIncrement:false" json:"rol_id"`
	Permiso_id uint `gorm:"Primary_Key; autoIncrement:false" json:"permiso_id"`
	Rol        Rol
	Permiso    Permiso
}

type Tipo_Programa struct {
	Tipo_programa_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Nombre           string `gorm:"type:varchar(50); default:''" json:"nombre"`
	Activo           bool   `gorm:"type:boolean; default:0" json:"activo"`
}

type Usuario_Rol struct {
	Usuario_id uint `gorm:"Primary_Key; autoIncrement:false" json:"usuario_id"`
	Rol_id     uint `gorm:"Primary_Key; autoIncrement:false" json:"rol_id"`
	Usuario    Usuario
	Rol        Rol
}

type Usuario struct {
	Usuario_id uint   `gorm:"Primary_Key; type:bigint UNSIGNED; not null; auto_increment" json:"id"`
	Usuario    string `gorm:"type:varchar(200); default:''" json:"usuario"`
	Contrasena string `gorm:"type:varchar(200); default:''" json:"contrasena,omitempty"`
	Activo     bool   `gorm:"type:tinyint; default:0" json:"activo"`
}

type Actividades []Actividad
type Alumnos_Programas []Alumno_Programa
type Alumnos []Alumno
type Carreras []Carrera
type Controles_Expedientes []Control_Expendiente
type Depedencias []Dependencia
type Modalidades []Modalidad
type Permisos []Permiso
type Programas []Programa
type Roles_Permisos []Rol_Permiso
type Roles []Rol
type Tipos_Programas []Tipo_Programa
type Usuarios_Roles []Usuario_Rol
type Usuarios []Usuario

func (Actividad) TableName() string {
	return "activadades"
}

func (Alumno_Programa) TableName() string {
	return "alumnos_programas"
}

func (Alumno) TableName() string {
	return "alumnos"
}

func (Carrera) TableName() string {
	return "carreras"
}

func (Control_Expendiente) TableName() string {
	return "control_expendientes"
}

func (Dependencia) TableName() string {
	return "dependecias"
}

func (Modalidad) TableName() string {
	return "modalidades"
}

func (Permiso) TableName() string {
	return "permisos"
}

func (Programa) TableName() string {
	return "programas"
}

func (Rol_Permiso) TableName() string {
	return "roles_permisos"
}

func (Rol) TableName() string {
	return "roles"
}

func (Tipo_Programa) TableName() string {
	return "tipo_programas"
}

func (Usuario_Rol) TableName() string {
	return "usarios_roles"
}

func (Usuario) TableName() string {
	return "usuarios"
}
