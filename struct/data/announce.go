package data

type Announce struct {
	Id string `bson:"_id"`
	Title string `bson:"title"`
}

func NewAnnounce() *Announce {
	return &Announce{}
}

func (a *Announce) CollectionName() string {
	return "Announce"
}
