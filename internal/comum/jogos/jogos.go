package jogos

type AoAdicionarPonto func(scoreA, scoreB int, done bool)
type OnAfterAddingGame func(scoreA, scoreB int, done bool)
type AoAdicionarSet func(scoreA, scoreB int, done bool)
type OnPlayerChangeSide func()
