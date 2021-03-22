package models

type Persona struct {
	ID        int64  `json:"id" grom:primary_key;auto_increment"`
	Nombre    string `json:"nombre`
	Apellido  string `json:"apellido`
	Direccion string `json:"direccion`
	Telefono  string `json:"telefono`
}
