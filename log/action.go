package log

type ActionLog struct {
	Ip      string
	Content string
}

func (al ActionLog) WriteAction() (err error) {
	return nil
}
