package vorbereitung

//Registrierung ist das zentrale Domänenobjekt des Registrierungsservice
type Registrierung struct {
	Firstname             string
	Lastname              string
	Email                 string
	Firma                 string
	Schulungscode         string
	Datum                 string
	DatenschutzAkzeptiert bool
	//Adresse, etc nicht dargestellt
}
