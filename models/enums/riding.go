package enums

const (
	Biginner SkillLevel = iota + 1
	Novice
	Intermidiate
	Advanced
	Expert
	Instructor
)

const (
	AdventureRiding RidingStyle = iota + 1
	BeginnerFriendly
	Biker
	Cruising
	DirtOffRoad
	RacingCompetition
	Social
	SportRiding
	SportTouring
	Touring
	TrackDay
	Urban
)

const (
	Adventure RideType = iota + 1
	CafeRacer
	Cruiser
	Custom
	Dirtbike
	DualSport
	Scooter
	SportTourer
	Sportbike
	StandardNaked
	StreetFighter
	Supermoto
	Tourer
	Trials
	Trike3Wheel
	VintageAntique
)

type (
	RidingStyle int
	SkillLevel  int
	RideType    int
)
