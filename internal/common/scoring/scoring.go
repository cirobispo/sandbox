package scoring

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

func Score2GameText(values []string, scoreA, scoreB int) (string, string) {
	if len(values) != 6 {
		panic("Descrição dos pontos incorreta (0, 15, 30, 40, vantagem e jogo)")
	}

	getText := func(score, b int) string {
		if score == 5 || score == 4 && b < 3 {
			return values[len(values)-1]
		}
		return values[score]
	}

	return getText(scoreA, scoreB), getText(scoreB, scoreA)
}
