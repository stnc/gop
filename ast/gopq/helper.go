/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gopq

import (
	"strconv"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/token"
)

// -----------------------------------------------------------------------------

func (p NodeSet) UnquotedString__1(exactly bool) (ret string, err error) {
	item, err := p.CollectOne__1(exactly)
	if err != nil {
		return
	}
	if lit, ok := item.Obj().(*ast.BasicLit); ok && lit.Kind == token.STRING {
		return strconv.Unquote(lit.Value)
	}
	return "", ErrUnexpectedNode
}

func (p NodeSet) UnquotedString__0() (ret string, err error) {
	return p.UnquotedString__1(false)
}

// -----------------------------------------------------------------------------

func (p NodeSet) UnquotedStringElts__1(exactly bool) (ret []string, err error) {
	item, err := p.CollectOne__1(exactly)
	if err != nil {
		return
	}
	if lit, ok := item.Obj().(*ast.CompositeLit); ok {
		ret = make([]string, len(lit.Elts))
		for i, elt := range lit.Elts {
			if lit, ok := elt.(*ast.BasicLit); ok && lit.Kind == token.STRING {
				if ret[i], err = strconv.Unquote(lit.Value); err != nil {
					return
				}
			} else {
				return nil, ErrUnexpectedNode
			}
		}
		return
	}
	return nil, ErrUnexpectedNode
}

func (p NodeSet) UnquotedStringElts__0() (ret []string, err error) {
	return p.UnquotedStringElts__1(false)
}

// -----------------------------------------------------------------------------

func (p NodeSet) Ident__1(exactly bool) (ret string, err error) {
	item, err := p.CollectOne__1(exactly)
	if err != nil {
		return
	}
	if ident, ok := item.Obj().(*ast.Ident); ok {
		return ident.Name, nil
	}
	return "", ErrUnexpectedNode
}

func (p NodeSet) Ident__0() (ret string, err error) {
	return p.Ident__1(false)
}

// -----------------------------------------------------------------------------

/*
// VarSpec returns variables *ast.ValueSpec node set.
func (p NodeSet) VarSpec() NodeSet {
	return p.GenDecl(token.VAR).Child()
}

// ConstSpec returns constants *ast.ValueSpec node set.
func (p NodeSet) ConstSpec() NodeSet {
	return p.GenDecl(token.CONST).Child()
}
*/

// -----------------------------------------------------------------------------
