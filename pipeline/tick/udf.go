package tick

import (
	"time"

	"github.com/influxdata/kapacitor/pipeline"
	"github.com/influxdata/kapacitor/tick/ast"
	"github.com/influxdata/kapacitor/udf/agent"
)

// UDF converts the UDF pipeline node into the TICKScript AST
type UDF struct {
	Function
}

// Build creates a UDF ast.Node
func (n *UDF) Build(u *pipeline.UDFNode) (ast.Node, error) {
	n.At(u.UDFName)
	for _, o := range u.Options {
		args := []interface{}{}
		for _, v := range o.Values {
			switch v.Type {
			case agent.ValueType_BOOL:
				args = append(args, v.GetBoolValue())
			case agent.ValueType_INT:
				args = append(args, v.GetIntValue())
			case agent.ValueType_DOUBLE:
				args = append(args, v.GetDoubleValue())
			case agent.ValueType_STRING:
				args = append(args, v.GetStringValue())
			case agent.ValueType_DURATION:
				dur := time.Duration(v.GetDurationValue())
				args = append(args, dur)
			}
		}
		n.Dot(o.Name, args)
	}
	return n.prev, n.err
}
