package jogos

type AoAdicionarPonto func(placarA, placarB int, terminado bool)
type OnAfterAddingGame func(placarA, placarB int, terminado bool)
type AoAdicionarSet func(placarA, placarB int, terminado bool)
type OnPlayerChangeSide func()
