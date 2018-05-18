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
	Adventure      RideType = iota + 1
	CafeRacer               // 2
	Cruiser                 // 3
	Custom                  // 4
	Dirtbike                // 5
	DualSport               // 6
	Scooter                 // 7
	SportTourer             // 8
	Sportbike               // 9
	StandardNaked           // 10
	StreetFighter           // 11
	Supermoto               // 12
	Tourer                  // 13
	Trials                  // 14
	Trike3Wheel             // 15
	VintageAntique          // 16
)

type (
	RidingStyle int
	SkillLevel  int
	RideType    int
)
