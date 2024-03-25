package domain

// CoordenadasGeograficas representa as coordenadas de latitude e longitude
type GeoCoordinates struct {
	Latitude  float64
	Longitude float64
}

// NewGeoCoordinates cria e retorna novas coordenadas geográficas
func NewGeoCoordinates(latitude, longitude float64) (*GeoCoordinates, error) {
	return &GeoCoordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
