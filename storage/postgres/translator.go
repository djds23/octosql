package postgres

import (
	"context"
	"fmt"

	"github.com/cube2222/octosql/physical"
	"github.com/pkg/errors"
)

type aliases struct {
	PlaceholderToExpression map[string]physical.Expression
	Counter                 int
	Alias                   string
}

func newAliases(alias string) *aliases {
	return &aliases{
		PlaceholderToExpression: make(map[string]physical.Expression),
		Counter:                 1,
		Alias:                   alias,
	}
}

func relationToSQL(rel physical.Relation) string {
	switch rel {
	case physical.Equal:
		return "="
	case physical.NotEqual:
		return "<>"
	case physical.MoreThan:
		return ">"
	case physical.LessThan:
		return "<"
	case physical.In:
		return "IN"
	case physical.Like:
		return "LIKE"
	default:
		panic("Invalid physical relation")
	}
}

func (aliases *aliases) newPlaceholder() string {
	str := fmt.Sprintf("$%d", aliases.Counter)
	aliases.Counter++
	return str
}

func expressionToSQL(expression physical.Expression, aliases *aliases) string {
	switch expression := expression.(type) {
	case *physical.Variable:
		if expression.Name.Source() == aliases.Alias {
			return expression.Name.String()
		}

		placeholder := aliases.newPlaceholder()
		aliases.PlaceholderToExpression[placeholder] = expression

		return placeholder

	default:
		placeholder := aliases.newPlaceholder()
		aliases.PlaceholderToExpression[placeholder] = expression

		return placeholder
	}
}

func formulaToSQL(formula physical.Formula, aliases *aliases) string {
	switch formula := formula.(type) {
	case *physical.And:
		left := formulaToSQL(formula.Left, aliases)

		right := formulaToSQL(formula.Right, aliases)

		return fmt.Sprintf("%s AND %s", parenthesize(left), parenthesize(right))

	case *physical.Or:
		left := formulaToSQL(formula.Left, aliases)

		right := formulaToSQL(formula.Right, aliases)

		return fmt.Sprintf("%s OR %s", parenthesize(left), parenthesize(right))

	case *physical.Not:
		child := formulaToSQL(formula.Child, aliases)

		return fmt.Sprintf("NOT %s", parenthesize(child))
	case *physical.Constant:
		if formula.Value {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case *physical.Predicate:
		left := expressionToSQL(formula.Left, aliases)

		right := expressionToSQL(formula.Right, aliases)

		relationString := relationToSQL(formula.Relation)

		return fmt.Sprintf("%s %s %s", parenthesize(left), relationString, parenthesize(right))
	default:
		panic("Unknown type of physical.Formula")
	}
}

func parenthesize(str string) string {
	return fmt.Sprintf("(%s)", str)
}

func (aliases *aliases) materializeAliases() (*executionAliases, error) {
	result := newExecutionAliases()

	ctx := context.Background()

	for placeholder, expression := range aliases.PlaceholderToExpression {
		exec, err := expression.Materialize(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't materialize expression")
		}

		result.PlaceholderToExpression[placeholder] = exec
	}

	return result, nil
}
