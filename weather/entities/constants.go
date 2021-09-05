package entities

type UnitType int

const (
	Feet                         UnitType = 0
	Inches                       UnitType = 1
	Miles                        UnitType = 2
	Millimeter                   UnitType = 3
	Centimeter                   UnitType = 4
	Meter                        UnitType = 5
	Kilometer                    UnitType = 6
	KilometersPerHour            UnitType = 7
	Knots                        UnitType = 8
	MilesPerHour                 UnitType = 9
	MetersPerSecond              UnitType = 10
	HectoPascals                 UnitType = 11
	InchesOfMercury              UnitType = 12
	KiloPascals                  UnitType = 13
	Millibars                    UnitType = 14
	MillimetersOfMercury         UnitType = 15
	PoundsPerSquareInch          UnitType = 16
	Celsius                      UnitType = 17
	Fahrenheit                   UnitType = 18
	Kelvin                       UnitType = 19
	Percent                      UnitType = 20
	Float                        UnitType = 21
	Integer                      UnitType = 22
	MicrogramsPerCubicMeterOfAir UnitType = 31
)

type LatestStatusKeyword string

const (
	Cancel   LatestStatusKeyword = "Cancel"
	Continue LatestStatusKeyword = "Continue"
	Correct  LatestStatusKeyword = "Correct"
	Expire   LatestStatusKeyword = "Expire"
	Extend   LatestStatusKeyword = "Extend"
	New      LatestStatusKeyword = "New"
	Update   LatestStatusKeyword = "Update"
	Upgrade  LatestStatusKeyword = "Upgrade"
)

type Quarter int

const (
	Morning   Quarter = 0
	Afternoon Quarter = 1
	Evening   Quarter = 2
	Overnight Quarter = 3
)
