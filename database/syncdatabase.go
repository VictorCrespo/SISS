package database

import "github.com/VictorCrespo/SISS/models"

func SyncDatabase() {
	DB.AutoMigrate(models.Usuario{}, models.Rol{}, models.Usuario_Rol{}, models.Permiso{},
		models.Rol_Permiso{}, models.Carrera{}, models.Alumno{}, models.Control_Expendiente{},
		models.Dependencia{}, models.Tipo_Programa{}, models.Modalidad{}, models.Actividad{},
		models.Alumno_Programa{})
}
