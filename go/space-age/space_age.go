package space

// Planet represents a planet
type Planet string

const oneEarthYear float64 = 31_557_600

var years = map[Planet]float64{
	"Earth":   oneEarthYear,
	"Mercury": oneEarthYear * 0.2408467,
	"Venus":   oneEarthYear * 0.61519726,
	"Mars":    oneEarthYear * 1.8808158,
	"Jupiter": oneEarthYear * 11.862615,
	"Saturn":  oneEarthYear * 29.447498,
	"Uranus":  oneEarthYear * 84.016846,
	"Neptune": oneEarthYear * 164.79132,
}

// Age returns how many ages represents some seconds in a planet
func Age(seconds float64, planet Planet) float64 {
	planetYears, ok := years[planet]
	if !ok {
		return -1
	}

	return seconds / planetYears
}
