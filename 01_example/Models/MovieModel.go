package Models

import "time"

type Movie struct {
	MovieId  int
	Title    string
	Released time.Time
	Price    float32
}

//func (m *Movie) SeedTestData() []Movie {

//}
