package ctrl

type Pager struct {
	Limit int
	Page  int
}

func (p *Pager) GetOffset() int {
	return p.Limit * p.Page
}
