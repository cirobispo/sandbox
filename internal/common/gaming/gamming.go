package gaming

type OnAfterAddingPoint func(scoreA, scoreB int, done bool)
type OnAfterAddingGame func(scoreA, scoreB int, done bool)
type OnAfterAddingSet func(scoreA, scoreB int, done bool)
type OnPlayerChangeSide func()
