// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package override

import (
	"github.com/kyleconroy/sqlc-testdata/pkg"
)

type Foo struct {
	Other   string
	Total   int64
	Retyped pkg.CustomType
}
