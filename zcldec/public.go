package zcldec

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-zcl/zcl"
)

// Decode interprets the given body using the given specification and returns
// the resulting value. If the given body is not valid per the spec, error
// diagnostics are returned and the returned value is likely to be incomplete.
//
// The ctx argument may be nil, in which case any references to variables or
// functions will produce error diagnostics.
func Decode(body zcl.Body, spec Spec, ctx *zcl.EvalContext) (cty.Value, zcl.Diagnostics) {
	val, _, diags := decode(body, nil, ctx, spec, false)
	return val, diags
}

// PartialDecode is like Decode except that it permits "leftover" items in
// the top-level body, which are returned as a new body to allow for
// further processing.
//
// Any descendent block bodies are _not_ decoded partially and thus must
// be fully described by the given specification.
func PartialDecode(body zcl.Body, spec Spec, ctx *zcl.EvalContext) (cty.Value, zcl.Body, zcl.Diagnostics) {
	return decode(body, nil, ctx, spec, true)
}