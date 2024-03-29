package domain

// CoordenadasGeograficas representa as coordenadas de latitude e longitude
type GeoCoordinates struct {
	Name      string  `json:"location_name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// NewGeoCoordinates cria e retorna novas coordenadas geográficas
func NewGeoCoordinates(name string, latitude, longitude float64) (*GeoCoordinates, error) {
	return &GeoCoordinates{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
