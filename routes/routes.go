package routes

import (
	"net/http"

	"github.com/VictorCrespo/SISS/authentication"
	"github.com/VictorCrespo/SISS/handlers"
	"github.com/VictorCrespo/SISS/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {

	private := r.NewRoute().Subrouter()

	private.Use(middleware.AuthJwtToken)

	private.HandleFunc("/actividades", handlers.GetActividades(db)).Methods(http.MethodGet)
	private.HandleFunc("/actividades/{id}", handlers.GetActividad(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/actividades", handlers.CreateActividad(db)).Methods(http.MethodPost)
	private.HandleFunc("/actividades/{id}", handlers.DeleteActividad(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/actividades/{id}", handlers.UpdateActividad(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/alumnos/programas", handlers.GetAlumnos_Programas(db)).Methods(http.MethodGet)
	private.HandleFunc("/alumnos/{alumno_id}/programas/{programa_id}", handlers.GetAlumnos_programa(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/alumnos/programas", handlers.CreateAlumnos_programa(db)).Methods(http.MethodPost)
	private.HandleFunc("/alumnos/{alumno_id}/programas/{programa_id}", handlers.DeleteAlumnos_Programa(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/alumnos/{alumno_id}/programas/{programa_id}", handlers.UpdateAlumnos_Programa(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/alumnos", handlers.GetAlumnos(db)).Methods(http.MethodGet)
	private.HandleFunc("/alumnos/{id}", handlers.GetAlumno(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/alumnos", handlers.CreateAlumno(db)).Methods(http.MethodPost)
	private.HandleFunc("/alumnos/{id}", handlers.DeleteAlumno(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/alumnos/{id}", handlers.UpdateAlumno(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/carreras", handlers.GetCarreras(db)).Methods(http.MethodGet)
	private.HandleFunc("/carreras/{id}", handlers.GetCarrera(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/carreras", handlers.CreateCarrera(db)).Methods(http.MethodPost)
	private.HandleFunc("/carreras/{id}", handlers.DeleteCarrera(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/carreras/{id}", handlers.UpdateCarrera(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/control_expendientes", handlers.GetControl_Expedientes(db)).Methods(http.MethodGet)
	private.HandleFunc("/control_expendientes/{id}", handlers.GetControl_Expediente(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/control_expendientes", handlers.CreateControl_Expediente(db)).Methods(http.MethodPost)
	private.HandleFunc("/control_expendientes/{id}", handlers.DeleteControl_Expediente(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/control_expendientes/{id}", handlers.UpdateControl_Expediente(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/dependencias", handlers.GetDependencias(db)).Methods(http.MethodGet)
	private.HandleFunc("/dependencias/{id}", handlers.GetDependencia(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/dependencias", handlers.CreateDependecia(db)).Methods(http.MethodPost)
	private.HandleFunc("/dependencias/{id}", handlers.DeleteDependencia(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/dependencias/{id}", handlers.UpdateDependecia(r, db)).Methods(http.MethodPut)

	r.HandleFunc("/login", authentication.Login(r, db)).Methods(http.MethodPost)

	private.HandleFunc("/modalidades", handlers.GetModalidades(db)).Methods(http.MethodGet)
	private.HandleFunc("/modalidades/{id}", handlers.GetModalidad(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/modalidades", handlers.CreateModalidad(db)).Methods(http.MethodPost)
	private.HandleFunc("/modalidades/{id}", handlers.DeleteModalidad(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/modalidades/{id}", handlers.UpdateModalidad(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/permisos", handlers.GetPermisos(db)).Methods(http.MethodGet)
	private.HandleFunc("/permisos/{id}", handlers.GetPermiso(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/permisos", handlers.CreatePermiso(db)).Methods(http.MethodPost)
	private.HandleFunc("/permisos/{id}", handlers.DeletePermiso(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/permisos/{id}", handlers.UpdatePermiso(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/programas", handlers.GetProgramas(db)).Methods(http.MethodGet)
	private.HandleFunc("/programas/{id}", handlers.GetPrograma(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/programas", handlers.CreatePrograma(db)).Methods(http.MethodPost)
	private.HandleFunc("/programas/{id}", handlers.DeletePrograma(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/programas/{id}", handlers.UpdatePrograma(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/roles/permisos", handlers.GetRoles_Permisos(db)).Methods(http.MethodGet)
	private.HandleFunc("/roles/{rol_id}/permisos/{permiso_id}", handlers.GetRol_Permiso(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/roles/permisos", handlers.CreateRol_Permiso(db)).Methods(http.MethodPost)
	private.HandleFunc("/roles/{rol_id}/permisos/{permiso_id}", handlers.DeleteRol_Permiso(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/roles/{rol_id}/permisos/{permiso_id}", handlers.UpdateRol_Permiso(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/roles", handlers.GetRoles(db)).Methods(http.MethodGet)
	private.HandleFunc("/roles/{id}", handlers.GetRol(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/roles", handlers.CreateRol(db)).Methods(http.MethodPost)
	private.HandleFunc("/roles/{id}", handlers.DeleteRol(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/roles/{id}", handlers.UpdateRol(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/tipo_programas", handlers.GetTipo_Programas(db)).Methods(http.MethodGet)
	private.HandleFunc("/tipo_programas/{id}", handlers.GetTipo_Programa(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/tipo_programas", handlers.CreateTipo_Programa(db)).Methods(http.MethodPost)
	private.HandleFunc("/tipo_programas/{id}", handlers.DeleteTipo_Programa(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/tipo_programas/{id}", handlers.UpdateTipo_Programa(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/usuarios/roles", handlers.GetUsuarios_Roles(db)).Methods(http.MethodGet)
	private.HandleFunc("/usuarios/{usuario_id}/roles/{rol_id}", handlers.GetUsuario_Rol(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/usuarios/roles", handlers.CreateUsuario_Rol(db)).Methods(http.MethodPost)
	private.HandleFunc("/usuarios/{usuario_id}/roles/{rol_id}", handlers.DeleteUsuario_Rol(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/usuarios/{usuario_id}/roles/{rol_id}", handlers.UpdateUsuario_Rol(r, db)).Methods(http.MethodPut)

	private.HandleFunc("/usuarios", handlers.GetUsuarios(db)).Methods(http.MethodGet)
	private.HandleFunc("/usuarios/{id}", handlers.GetUsuario(r, db)).Methods(http.MethodGet)
	private.HandleFunc("/usuarios", handlers.CreateUsuario(db)).Methods(http.MethodPost)
	private.HandleFunc("/usuarios/{id}", handlers.DeleteUsuario(r, db)).Methods(http.MethodDelete)
	private.HandleFunc("/usuarios/{id}", handlers.UpdateUsuario(r, db)).Methods(http.MethodPut)
}
