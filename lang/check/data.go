// Code generated by running "go generate". DO NOT EDIT.

// Copyright 2017 The Wuffs Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package check

import (
	a "github.com/google/wuffs/lang/ast"
	t "github.com/google/wuffs/lang/token"
)

var reasons = [...]struct {
	s string
	r reason
}{

	{`"a < b: b > a"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		// b > a
		if err := proveReasonRequirement(q, t.IDXBinaryGreaterThan, xb, xa); err != nil {
			return err
		}
		return nil
	}},

	{`"a < b: a < c; c < b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		// a < c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xa, xc); err != nil {
			return err
		}
		// c < b
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a < b: a < c; c == b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		// a < c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xa, xc); err != nil {
			return err
		}
		// c == b
		if err := proveReasonRequirement(q, t.IDXBinaryEqEq, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a < b: a == c; c < b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		// a == c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryEqEq, xa, xc); err != nil {
			return err
		}
		// c < b
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a < b: a < c; c <= b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		// a < c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xa, xc); err != nil {
			return err
		}
		// c <= b
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a < b: a <= c; c < b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		// a <= c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xa, xc); err != nil {
			return err
		}
		// c < b
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a <= b: b >= a"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessEq {
			return errFailed
		}
		// b >= a
		if err := proveReasonRequirement(q, t.IDXBinaryGreaterEq, xb, xa); err != nil {
			return err
		}
		return nil
	}},

	{`"a <= b: a <= c; c <= b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessEq {
			return errFailed
		}
		// a <= c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xa, xc); err != nil {
			return err
		}
		// c <= b
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a <= b: a <= c; c == b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessEq {
			return errFailed
		}
		// a <= c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xa, xc); err != nil {
			return err
		}
		// c == b
		if err := proveReasonRequirement(q, t.IDXBinaryEqEq, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a <= b: a == c; c <= b"`, func(q *checker, n *a.Assert) error {
		op, xa, xb := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessEq {
			return errFailed
		}
		// a == c
		xc := argValue(q.tm, n.Args(), "c")
		if xc == nil {
			return errFailed
		}
		if err := proveReasonRequirement(q, t.IDXBinaryEqEq, xa, xc); err != nil {
			return err
		}
		// c <= b
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xc, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a < (b + c): a < c; 0 <= b"`, func(q *checker, n *a.Assert) error {
		op, xa, t0 := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		op, xb, xc := parseBinaryOp(t0)
		if op.Key() != t.KeyXBinaryPlus {
			return errFailed
		}
		// a < c
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xa, xc); err != nil {
			return err
		}
		// 0 <= b
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, zeroExpr, xb); err != nil {
			return err
		}
		return nil
	}},

	{`"a < (b + c): a < (b0 + c0); b0 <= b; c0 <= c"`, func(q *checker, n *a.Assert) error {
		op, xa, t0 := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessThan {
			return errFailed
		}
		op, xb, xc := parseBinaryOp(t0)
		if op.Key() != t.KeyXBinaryPlus {
			return errFailed
		}
		// a < (b0 + c0)
		xb0 := argValue(q.tm, n.Args(), "b0")
		if xb0 == nil {
			return errFailed
		}
		xc0 := argValue(q.tm, n.Args(), "c0")
		if xc0 == nil {
			return errFailed
		}
		t1 := a.NewExpr(a.FlagsTypeChecked, t.IDXBinaryPlus, 0, 0, xb0.Node(), nil, xc0.Node(), nil)
		if err := proveReasonRequirement(q, t.IDXBinaryLessThan, xa, t1); err != nil {
			return err
		}
		// b0 <= b
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xb0, xb); err != nil {
			return err
		}
		// c0 <= c
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xc0, xc); err != nil {
			return err
		}
		return nil
	}},

	{`"(a + b) <= c: a <= (c - b)"`, func(q *checker, n *a.Assert) error {
		op, t0, xc := parseBinaryOp(n.Condition())
		if op.Key() != t.KeyXBinaryLessEq {
			return errFailed
		}
		op, xa, xb := parseBinaryOp(t0)
		if op.Key() != t.KeyXBinaryPlus {
			return errFailed
		}
		// a <= (c - b)
		t1 := a.NewExpr(a.FlagsTypeChecked, t.IDXBinaryMinus, 0, 0, xc.Node(), nil, xb.Node(), nil)
		if err := proveReasonRequirement(q, t.IDXBinaryLessEq, xa, t1); err != nil {
			return err
		}
		return nil
	}},
}
