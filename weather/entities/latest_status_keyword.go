package entities

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
