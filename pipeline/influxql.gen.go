// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: influxql.gen.go.tmpl

package pipeline

import "github.com/influxdata/influxdb/query"

//tick:ignore
type ReduceCreater struct {
	CreateFloatReducer func() (query.FloatPointAggregator, query.FloatPointEmitter)

	CreateFloatIntegerReducer func() (query.FloatPointAggregator, query.IntegerPointEmitter)

	CreateFloatStringReducer func() (query.FloatPointAggregator, query.StringPointEmitter)

	CreateFloatBooleanReducer func() (query.FloatPointAggregator, query.BooleanPointEmitter)

	CreateIntegerFloatReducer func() (query.IntegerPointAggregator, query.FloatPointEmitter)

	CreateIntegerReducer func() (query.IntegerPointAggregator, query.IntegerPointEmitter)

	CreateIntegerStringReducer func() (query.IntegerPointAggregator, query.StringPointEmitter)

	CreateIntegerBooleanReducer func() (query.IntegerPointAggregator, query.BooleanPointEmitter)

	CreateStringFloatReducer func() (query.StringPointAggregator, query.FloatPointEmitter)

	CreateStringIntegerReducer func() (query.StringPointAggregator, query.IntegerPointEmitter)

	CreateStringReducer func() (query.StringPointAggregator, query.StringPointEmitter)

	CreateStringBooleanReducer func() (query.StringPointAggregator, query.BooleanPointEmitter)

	CreateBooleanFloatReducer func() (query.BooleanPointAggregator, query.FloatPointEmitter)

	CreateBooleanIntegerReducer func() (query.BooleanPointAggregator, query.IntegerPointEmitter)

	CreateBooleanStringReducer func() (query.BooleanPointAggregator, query.StringPointEmitter)

	CreateBooleanReducer func() (query.BooleanPointAggregator, query.BooleanPointEmitter)

	TopBottomCallInfo      *TopBottomCallInfo
	IsSimpleSelector       bool
	IsStreamTransformation bool
	IsEmptyOK              bool
}
