package main

type Strategy struct {
	ChargeTime  int64
	Speed       int64
	Distance    int64
	BeatsRecord bool
}

func GetStrategies(race Race, start int64, end int64) *[]Strategy {
	stragies := []Strategy{}

	for chargeTime := start; chargeTime < end; chargeTime++ {
		distance := getDistance(chargeTime, race)
		strategy := Strategy{
			ChargeTime:  chargeTime,
			Speed:       chargeTime,
			Distance:    distance,
			BeatsRecord: distance > race.RecordDistance,
		}
		stragies = append(stragies, strategy)
	}
	return &stragies
}

func getDistance(chargeTime int64, race Race) int64 {
	travelTime := race.Time - chargeTime
	speed := chargeTime
	return travelTime * speed
}

func getWinning(strategies *[]Strategy) int64 {
	var winning int64 = 0

	for _, strategy := range *strategies {
		if !strategy.BeatsRecord {
			continue
		}
		winning++
	}
	return winning
}
