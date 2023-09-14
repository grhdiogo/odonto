package appointment

import "time"

type Item struct {
	ProcedurePID   string
	ProcedureName  string
	ProcedureValue float64
	Tooth          int
}
type Entity struct {
	Status      string
	DoctorDid   string
	PatientPid  string
	Observation string
	ID          Identity
	Items       []Item
}

type EntityProxy struct {
	Entity
	DoctorName  string
	PatientName string
	CreatedAt   time.Time
}
