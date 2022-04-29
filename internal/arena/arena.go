package arena

type Battle struct {
	Red  Team
	Blue Team
}
type Arena struct {
	Battles map[int]Battle
}
