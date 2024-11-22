package utils

func Fill(Aircraftno, flownhours, dailylandings int) *Aircraft{
	tempAircraft := &Aircraft{}

	tempAircraft.AircraftNo = Aircraftno
	tempAircraft.FlownHours = flownhours
	tempAircraft.DailyLandings = dailylandings

	return tempAircraft
}