package pokemon

type Stats struct {
	HP  int
	Atk int
	Def int
	Spd int
	Spc int
}

var pokemonStats = map[string]Stats{
	"bulbasaur":  {HP: 45, Atk: 49, Def: 49, Spd: 45, Spc: 65},
	"ivysaur":    {HP: 60, Atk: 62, Def: 63, Spd: 60, Spc: 80},
	"venusaur":   {HP: 80, Atk: 82, Def: 83, Spd: 80, Spc: 100},
	"charmander": {HP: 39, Atk: 52, Def: 43, Spd: 65, Spc: 50},
	"charmeleon": {HP: 58, Atk: 64, Def: 58, Spd: 80, Spc: 65},
	"charizard":  {HP: 78, Atk: 84, Def: 78, Spd: 100, Spc: 85},
	"squirtle":   {HP: 44, Atk: 48, Def: 65, Spd: 43, Spc: 50},
	"wartortle":  {HP: 59, Atk: 63, Def: 80, Spd: 58, Spc: 65},
	"blastoise":  {HP: 79, Atk: 83, Def: 100, Spd: 78, Spc: 85},
	"caterpie":   {HP: 45, Atk: 30, Def: 35, Spd: 45, Spc: 20},
	"metapod":    {HP: 50, Atk: 20, Def: 55, Spd: 30, Spc: 25},
	"butterfree": {HP: 60, Atk: 45, Def: 50, Spd: 70, Spc: 80},
}

func GetStats(pokemon string) Stats {
	if stats, ok := pokemonStats[pokemon]; ok {
		return stats
	} else {
		return Stats{}
	}
}
