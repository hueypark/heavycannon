package bullet

import "github.com/heavycannon/heavycannon/shape"

type Bullet struct {
}

func New() *Bullet {
	return &Bullet{}
}

func (b *Bullet) Type() int64 {
	return shape.BULLET
}
