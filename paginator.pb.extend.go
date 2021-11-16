package paginator

func (x *Paginator) Offset() uint64 {
	if x.Page == 1 {
		return 0
	}

	return (x.Page - 1) * x.Limit
}
