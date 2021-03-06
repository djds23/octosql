package aggregates

import (
	"github.com/cube2222/octosql"
	"github.com/cube2222/octosql/docs"
	"github.com/cube2222/octosql/execution"
	"github.com/pkg/errors"
)

type Max struct {
	maxes      *execution.HashMap
	typedValue octosql.Value
}

func NewMax() *Max {
	return &Max{
		maxes: execution.NewHashMap(),
	}
}

func (agg *Max) Document() docs.Documentation {
	return docs.Section(
		agg.String(),
		docs.Body(
			docs.Section("Description", docs.Text("Takes the maximum element in the group. Works with Ints, Floats, Strings, Booleans, Times, Durations.")),
		),
	)
}

func (agg *Max) AddRecord(key octosql.Tuple, value octosql.Value) error {
	max, previousValueExists, err := agg.maxes.Get(key)
	if err != nil {
		return errors.Wrap(err, "couldn't get current max out of hashmap")
	}

	if agg.typedValue == nil {
		agg.typedValue = value
	}
	switch value := value.(type) {
	case octosql.Int:
		_, typeOk := agg.typedValue.(octosql.Int)
		if !typeOk {
			return errors.Errorf("mixed types in max: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}

		if !previousValueExists || value > max.(octosql.Int) {
			max = value
		}

	case octosql.Float:
		_, typeOk := agg.typedValue.(octosql.Float)
		if !typeOk {
			return errors.Errorf("mixed types in max: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}

		if !previousValueExists || value > max.(octosql.Float) {
			max = value
		}

	case octosql.String:
		_, typeOk := agg.typedValue.(octosql.String)
		if !typeOk {
			return errors.Errorf("mixed types in max: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}

		if !previousValueExists || value > max.(octosql.String) {
			max = value
		}

	case octosql.Bool:
		_, typeOk := agg.typedValue.(octosql.Bool)
		if !typeOk {
			return errors.Errorf("mixed types in max: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}

		if !previousValueExists || value.AsBool() {
			max = value
		}

	case octosql.Duration:
		_, typeOk := agg.typedValue.(octosql.Duration)
		if !typeOk {
			return errors.Errorf("mixed types in max: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}

		if !previousValueExists || value > max.(octosql.Duration) {
			max = value
		}

	case octosql.Time:
		_, typeOk := agg.typedValue.(octosql.Time)
		if !typeOk {
			return errors.Errorf("mixed types in max: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}

		if !previousValueExists || value.AsTime().After(max.(octosql.Time).AsTime()) {
			max = value
		}
	default:
		return errors.Errorf("invalid type in max: %v with value %v", execution.GetType(value), value)
	}

	err = agg.maxes.Set(key, max)
	if err != nil {
		return errors.Wrap(err, "couldn't put new max into hashmap")
	}

	return nil
}

func (agg *Max) GetAggregated(key octosql.Tuple) (octosql.Value, error) {
	max, ok, err := agg.maxes.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get max out of hashmap")
	}

	if !ok {
		return nil, errors.Errorf("max for key not found")
	}

	return max.(octosql.Value), nil
}

func (agg *Max) String() string {
	return "max"
}
