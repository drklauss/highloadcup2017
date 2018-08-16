package models

var (
	UCache   *Users
	VCache   *Visits
	LocCache *Locations
	UvlCache *UserVisitLinks
)

func Init() {
	UCache = new(Users)
	VCache = new(Visits)
	LocCache = new(Locations)
	UvlCache = new(UserVisitLinks)
}
