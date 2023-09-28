package modelos

type Mujer struct {
	Hombre // Herencia, mujer hereda las propiedades del hombre
}

func (m *Mujer) Respirar() {
	m.respirando=true
}

func (m *Mujer) Comer() {
	m.comiendo=true
}

func (m *Mujer) Pensar() {
	m.pensando=true
}

func (h *Mujer) Sexo() string {
	return "Mujer"
}

func (m *Mujer) EstaVivo() bool {
	return true
}