// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package override

import ()

type Bar struct {
	Other      string
	AlsoTagged string `also:"tagged"`
}

type Baz struct {
	Other      string
	AlsoTagged string `also:"tagged"`
}

type Foo struct {
	Other  string
	Tagged string `a:"b" x:"y,z"`
}
