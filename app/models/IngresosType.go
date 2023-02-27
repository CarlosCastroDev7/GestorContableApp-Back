package models

type ListIngresosType struct {
	Fecha    string `json:"fecha,omitempty"`
	Concepto string `json:"concepto,omitempty"`
	Valor    int    `json:"valor,omitempty"`
}
