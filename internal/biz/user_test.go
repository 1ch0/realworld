package biz

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/davecgh/go-spew/spew"
)

func Test_hashPassword(t *testing.T) {
	s := hashPassword("sav")
	spew.Dump(s)
}

func Test_verifyPassword(t *testing.T) {
	a := assert.New(t)

	a.True(verifyPassword("$2a$10$uPFbJ2C4S/1TD3Yj/un/ge2.I1YffQcSOp9DTU7Hj5.njgceLi7k.", "sav"))
	a.False(verifyPassword("$2a$10$uPFbJ2C4S/1TD3Yj/un/ge2.I1YffQcSOp9DTU7Hj5.njgceLi7k.", "sav1"))
	a.False(verifyPassword("2C4S/1TD3Yj/un/ge2.I1YffQcSOp9DTU7Hj5.njgceLi7k.", "sav1"))
}
