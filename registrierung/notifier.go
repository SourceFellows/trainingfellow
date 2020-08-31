package registrierung

//RegistrierungsNotifier gibt Informationen zu Registrierungen weiter
type RegistrierungsNotifier interface {
	//InformAboutNewRegistrierung informiert über eine neue Registrierung
	InformAboutNewRegistrierung(reg *Registrierung) error
}
