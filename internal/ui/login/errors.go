package login

import "github.com/ayn2op/tview"

func ErrorMsg(err error) tview.Msg {
	return errMsg{err: err}
}
