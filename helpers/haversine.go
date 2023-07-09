package helpers

import "math"

func HaversineDistance(lat2, lon2 float64) float64 {
	lat1Rad := degToRad(-6.713954)
	lon1Rad := degToRad(108.489270)
	lat2Rad := degToRad(lat2)
	lon2Rad := degToRad(lon2)

	latDiff := lat2Rad - lat1Rad
	lonDiff := lon2Rad - lon1Rad

	a := math.Pow(math.Sin(latDiff/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(lonDiff/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	earthRadius := 6371.0
	distance := earthRadius * c

	return distance
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
