package routes

import (
	"github.com/VictorCrespo/SISS/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/actividades", handlers.GetActividades(db)).Methods("GET")
	r.HandleFunc("/actividades/{id}", handlers.GetActividad(r, db)).Methods("GET")
	r.HandleFunc("/actividades", handlers.CreateActividad(db)).Methods("POST")
	r.HandleFunc("/actividades/{id}", handlers.DeleteActividad(r, db)).Methods("DELETE")
	r.HandleFunc("/actividades/{id}", handlers.UpdateActividad(r, db)).Methods("PUT")

	r.HandleFunc("/alumnos/programas", handlers.GetAlumnos_Programas(db)).Methods("GET")
	r.HandleFunc("/alumnos/programas/{id}", handlers.GetAlumnos_programa(r, db)).Methods("GET")
	r.HandleFunc("/alumnos/programas", handlers.CreateAlumnos_programa(db)).Methods("POST")
	r.HandleFunc("/alumnos/programas/{id}", handlers.DeleteAlumnos_Programa(r, db)).Methods("DELETE")
	r.HandleFunc("/alumnos/programas/{id}", handlers.UpdateAlumnos_Programa(r, db)).Methods("PUT")

	r.HandleFunc("/alumnos", handlers.GetAlumnos(db)).Methods("GET")
	r.HandleFunc("/alumnos/{id}", handlers.GetAlumno(r, db)).Methods("GET")
	r.HandleFunc("/alumnos", handlers.CreateAlumno(db)).Methods("POST")
	r.HandleFunc("/alumnos/{id}", handlers.DeleteAlumno(r, db)).Methods("DELETE")
	r.HandleFunc("/alumnos/{id}", handlers.UpdateAlumno(r, db)).Methods("PUT")

	r.HandleFunc("/carreras", handlers.GetCarreras(db)).Methods("GET")
	r.HandleFunc("/carreras/{id}", handlers.GetCarrera(r, db)).Methods("GET")
	r.HandleFunc("/carreras", handlers.CreateCarrera(db)).Methods("POST")
	r.HandleFunc("/carreras/{id}", handlers.DeleteCarrera(r, db)).Methods("DELETE")
	r.HandleFunc("/carreras/{id}", handlers.UpdateCarrera(r, db)).Methods("PUT")

	r.HandleFunc("/control_expendientes", handlers.GetControl_Expedientes(db)).Methods("GET")
	r.HandleFunc("/control_expendientes/{id}", handlers.GetControl_Expediente(r, db)).Methods("GET")
	r.HandleFunc("/control_expendientes", handlers.CreateControl_Expediente(db)).Methods("POST")
	r.HandleFunc("/control_expendientes/{id}", handlers.DeleteControl_Expediente(r, db)).Methods("DELETE")
	r.HandleFunc("/control_expendientes/{id}", handlers.UpdateControl_Expediente(r, db)).Methods("PUT")

	r.HandleFunc("/dependencias", handlers.GetDependencias(db)).Methods("GET")
	r.HandleFunc("/dependencias/{id}", handlers.GetDependencia(r, db)).Methods("GET")
	r.HandleFunc("/dependencias", handlers.CreateDependecia(db)).Methods("POST")
	r.HandleFunc("/dependencias/{id}", handlers.DeleteDependencia(r, db)).Methods("DELETE")
	r.HandleFunc("/dependencias/{id}", handlers.UpdateDependecia(r, db)).Methods("PUT")

	r.HandleFunc("/modalidades", handlers.GetModalidades(db)).Methods("GET")
	r.HandleFunc("/modalidades/{id}", handlers.GetModalidad(r, db)).Methods("GET")
	r.HandleFunc("/modalidades", handlers.CreateModalidad(db)).Methods("POST")
	r.HandleFunc("/modalidades/{id}", handlers.DeleteModalidad(r, db)).Methods("DELETE")
	r.HandleFunc("/modalidades/{id}", handlers.UpdateModalidad(r, db)).Methods("PUT")

	r.HandleFunc("/permisos", handlers.GetPermisos(db)).Methods("GET")
	r.HandleFunc("/permisos/{id}", handlers.GetPermiso(r, db)).Methods("GET")
	r.HandleFunc("/permisos", handlers.CreatePermiso(db)).Methods("POST")
	r.HandleFunc("/permisos/{id}", handlers.DeletePermiso(r, db)).Methods("DELETE")
	r.HandleFunc("/permisos/{id}", handlers.UpdatePermiso(r, db)).Methods("PUT")

	r.HandleFunc("/programas", handlers.GetProgramas(db)).Methods("GET")
	r.HandleFunc("/programas/{id}", handlers.GetPrograma(r, db)).Methods("GET")
	r.HandleFunc("/programas", handlers.CreatePrograma(db)).Methods("POST")
	r.HandleFunc("/programas/{id}", handlers.DeletePrograma(r, db)).Methods("DELETE")
	r.HandleFunc("/programas/{id}", handlers.UpdatePrograma(r, db)).Methods("PUT")

	r.HandleFunc("/roles/permisos", handlers.GetRoles_Permisos(db)).Methods("GET")
	r.HandleFunc("/roles/permisos/{id}", handlers.GetRol_Permiso(r, db)).Methods("GET")
	r.HandleFunc("/roles/permisos", handlers.CreateRol_Permiso(db)).Methods("POST")
	r.HandleFunc("/roles/permisos/{id}", handlers.DeleteRol_Permiso(r, db)).Methods("DELETE")
	r.HandleFunc("/roles/permisos/{id}", handlers.UpdateRol_Permiso(r, db)).Methods("PUT")

	r.HandleFunc("/roles", handlers.GetRoles(db)).Methods("GET")
	r.HandleFunc("/roles/{id}", handlers.GetRol(r, db)).Methods("GET")
	r.HandleFunc("/roles", handlers.CreateRol(db)).Methods("POST")
	r.HandleFunc("/roles/{id}", handlers.DeleteRol(r, db)).Methods("DELETE")
	r.HandleFunc("/roles/{id}", handlers.UpdateRol(r, db)).Methods("PUT")

	r.HandleFunc("/tipo_programas", handlers.GetTipo_Programas(db)).Methods("GET")
	r.HandleFunc("/tipo_programas/{id}", handlers.GetTipo_Programa(r, db)).Methods("GET")
	r.HandleFunc("/tipo_programas", handlers.CreateTipo_Programa(db)).Methods("POST")
	r.HandleFunc("/tipo_programas/{id}", handlers.DeleteTipo_Programa(r, db)).Methods("DELETE")
	r.HandleFunc("/tipo_programas/{id}", handlers.UpdateTipo_Programa(r, db)).Methods("PUT")

	r.HandleFunc("/usuarios/roles", handlers.GetUsuarios_Roles(db)).Methods("GET")
	r.HandleFunc("/usuarios/roles/{id}", handlers.GetUsuario_Rol(r, db)).Methods("GET")
	r.HandleFunc("/usuarios/roles", handlers.CreateUsuario_Rol(db)).Methods("POST")
	r.HandleFunc("/usuarios/roles/{id}", handlers.DeleteUsuario_Rol(r, db)).Methods("DELETE")
	r.HandleFunc("/usuarios/roles/{id}", handlers.UpdateUsuario_Rol(r, db)).Methods("PUT")

	r.HandleFunc("/usuarios", handlers.GetUsuarios(db)).Methods("GET")
	r.HandleFunc("/usuarios/{id}", handlers.GetUsuario(r, db)).Methods("GET")
	r.HandleFunc("/usuarios", handlers.CreateUsuario(db)).Methods("POST")
	r.HandleFunc("/usuarios/{id}", handlers.DeleteUsuario(r, db)).Methods("DELETE")
	r.HandleFunc("/usuarios/{id}", handlers.UpdateUsuario(r, db)).Methods("PUT")
}
