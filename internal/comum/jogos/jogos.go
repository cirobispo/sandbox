package jogos

type AoAdicionarPonto func(placarA, placarB int, terminado bool)
type OnAfterAddingGame func(scoreA, scoreB int, done bool)
type AoAdicionarSet func(placarA, placarB int, terminado bool)
type OnPlayerChangeSide func()
