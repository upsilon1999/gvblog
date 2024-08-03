package es_ser

type Option struct {
	Page   int    `form:"page"`
	Key    string `form:"key"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort"`
	Fields []string
	Tag    string `form:"tag"`
}

func (op Option) GetFrom() int {
	if op.Limit == 0 {
		op.Limit = 10
	}
	if op.Page == 0 {
		op.Page = 1
	}
	return (op.Page - 1) * op.Limit
}
