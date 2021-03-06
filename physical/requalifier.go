package physical

import (
	"context"

	"github.com/cube2222/octosql/execution"
	"github.com/pkg/errors"
)

type Requalifier struct {
	Qualifier string
	Source    Node
}

func NewRequalifier(qualifier string, child Node) *Requalifier {
	return &Requalifier{Qualifier: qualifier, Source: child}
}

func (node *Requalifier) Transform(ctx context.Context, transformers *Transformers) Node {
	var transformed Node = &Requalifier{
		Qualifier: node.Qualifier,
		Source:    node.Source.Transform(ctx, transformers),
	}
	if transformers.NodeT != nil {
		transformed = transformers.NodeT(transformed)
	}
	return transformed
}

func (node *Requalifier) Materialize(ctx context.Context) (execution.Node, error) {
	materialized, err := node.Source.Materialize(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't materialize Source node")
	}
	return execution.NewRequalifier(node.Qualifier, materialized), nil
}
