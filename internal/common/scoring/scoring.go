package scoring

import "fmt"

type ScoringSide int

const (
	SSNone     ScoringSide = 0
	SSServing  ScoringSide = 1
	SSOpposite ScoringSide = 2
)

type ScoringType int

const (
	STPoint ScoringType = 0
	STGame  ScoringType = 1
	STSet   ScoringType = 2
	STMatch ScoringType = 3
)

type Scoring interface {
	Done() bool
	Side() ScoringSide
	Type() ScoringType
}

type Resulting interface {
	Result() (int, int)
}

type ScoringResulting interface {
	Scoring
	Resulting
}

func Side(ss ScoringResulting) ScoringSide {
	if !ss.Done() {
		return SSNone
	}

	if a, b := ss.Result(); b > a {
		return SSOpposite
	}

	return SSServing
}

func Score2GameText(scoreA, scoreB int) (string, string) {
	getText := func(score, b int) string {
		switch score {
		case 1, 2, 3:
			return fmt.Sprintf("%d", 15*score-((score/3)*5)) // works on int type
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
