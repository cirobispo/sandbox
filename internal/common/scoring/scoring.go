package scoring

type ScoringSide int

const (
	SSNone ScoringSide = 0
	SSA    ScoringSide = 1
	SSB    ScoringSide = 2
)

type ScoringType int

const (
	STGame  ScoringType = 0
	STSet   ScoringType = 1
	STMatch ScoringType = 2
)

type Scoring interface {
	Done() bool
	Result() (int, int)
	Side() ScoringSide
	Type() ScoringType
}

func Side(ss Scoring) ScoringSide {
	if !ss.Done() {
		return SSNone
	}

	if a, b := ss.Result(); b > a {
		return SSB
	}

	return SSA
}

func Score2GameText(scoreA, scoreB int) (string, string) {
	getText := func(score, b int) string {
		switch score {
		case 1:
			return "15"
		case 2:
			return "30"
		case 3:
			return "40"
		case 4, 5:
			if (score == 4 && b < 3) || (score == 5) {
				return "game"
			}
			return "ad"
		default:
			return "love"
		}
	}

	return getText(scoreA, scoreB), getText(scoreB, scoreA)
}
