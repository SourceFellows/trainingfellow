package registrierung

import "time"

//Registrierung ist das zentrale Domänenobjekt des Registrierungsservice
type Registrierung struct {
	ID                    string    `bson:"_id, omitempty"`
	Firstname             string    `form:"Firstname" binding:"required"`
	Lastname              string    `form:"Lastname" binding:"required"`
	Email                 string    `form:"Email" binding:"required"`
	Firma                 string    `form:"Firma"`
	Schulungscode         string    `form:"Schulungscode" binding:"required"`
	Datum                 time.Time `form:"Datum" binding:"required" time_format:"2006-01-02"`
	DatenschutzAkzeptiert bool      `form:"DatenschutzAkzeptiert" binding:"required"`
	Confirmed             bool
	//Adresse, etc nicht dargestellt
}
