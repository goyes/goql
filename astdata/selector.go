package astdata

import (
	"go/ast"
)

// SelectorType is the type in another package
type SelectorType struct {
	pkg *Package
	fl  *File

	selector string
	ident    string
	imp      *Import
}

func (s *SelectorType) String() string {
	return s.selector + "." + s.ident
}

// Package is the package of selector
func (s *SelectorType) Package() *Package {
	return s.pkg
}

// Selector is the selector type
func (s *SelectorType) Selector() string {
	return s.selector
}

// Ident is the ident after dot
func (s *SelectorType) Ident() string {
	return s.ident
}

// Import is the import of this selector
func (s *SelectorType) Import() *Import {
	return s.imp
}

func getSelector(p *Package, f *File, t *ast.SelectorExpr) Definition {
	it := t.X.(*ast.Ident)
	res := &SelectorType{
		pkg:      p,
		fl:       f,
		ident:    nameFromIdent(t.Sel),
		selector: nameFromIdent(it),
	}

	for i := range f.imports {
		if f.imports[i].Canonical() == res.selector || f.imports[i].TargetPackage() == res.selector {
			res.imp = f.imports[i]
		}
	}
	return res
}
