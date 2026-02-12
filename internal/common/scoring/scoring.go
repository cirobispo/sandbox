package scoring

func Score2GameText(scoreA, scoreB int) (string, string) {
	getText := func(score int) string {
		switch score {
		case 1:
			return "15"
		case 2:
			return "30"
		case 3:
			return "40"
		case 4, 5:
			if score == 5 {
				return "game"
			}
			return "ad"
		default:
			return "love"
		}
	}

	return getText(scoreA), getText(scoreB)
}
