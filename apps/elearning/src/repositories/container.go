package repositories

type Container struct {
	AlbumRepository      Repository
	TrackRepository      Repository
	ShelfRepository      Repository
	PreferenceRepository Repository
	SubjectRepository    Repository
}
