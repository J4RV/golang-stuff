package main

func ChooseColor(s State) State {
	res := clone(s)
	res.Phase = PlayerChoosingColor
	return res
}

func End(s State) State {
	res := clone(s)
	res.Phase = Finished
	return res
}
