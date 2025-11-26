package growel

import (

)

type growel struct { }

func (growel) New() (*growel) {
	return &growel{}
}

func (*growel) GET() {}
func (*growel) POST() {}
func (*growel) PUT() {}
func (*growel) DELETE() {}

func (*growel) Start() {}
