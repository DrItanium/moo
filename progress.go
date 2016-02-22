package moo

const (
	strPROGRESS_MESSAGES   = 143
	_distribute_map_single = iota
	_distribute_map_multiple
	_receiving_map
	_awaiting_map
	_distribute_physics_single
	_distribute_physics_multiple
	_receiving_physics
)

func openProgressDialog(messageId int16) {

}

func closeProgressDialog() {

}

func setProgressDialogMessage(messageId int16) {

}

func drawProgressBar(sent, total int32) {

}

func resetProgressBar() {

}
