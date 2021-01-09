package a23interpreter

import (
	"strings"
)

type ContextStu struct {
	Tokens []string
	Cnt    int
	Idx    int
}

func (r *ContextStu) Init(str string) {
	filterFunc := func(c rune) (out bool) {
		if ' ' == byte(c) || ',' == byte(c) {
			return true
		}
		return false

	}
	r.Tokens = strings.FieldsFunc(str, filterFunc)
	r.Cnt = len(r.Tokens)
	r.Idx = 0
}
func (r *ContextStu) NextToken() (toekn string, ret int) {
	ret = -1
	if r.Idx < r.Cnt {
		toekn = r.Tokens[r.Idx]
		r.Idx++
		ret = 0
	}
	return
}

func (r *ContextStu) CurrentToken() (token string) {
	if r.Idx < r.Cnt {
		token = r.Tokens[r.Idx]
		return
	}
	token = ""
	return
}

func (r *ContextStu) SkipToken(token string) (ret int) {
	if r.CurrentToken() != token {
		ret = -1
	}
	_, ret = r.NextToken()
	return
}

func (r *ContextStu) GetData() (toekns []string) {
	toekns = r.Tokens
	return
}
