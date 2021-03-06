package aggregates

import (
	"time"

	"github.com/cube2222/octosql"
	"github.com/cube2222/octosql/docs"
	"github.com/cube2222/octosql/execution"
	"github.com/pkg/errors"
)

type Average struct {
	averages   *execution.HashMap
	counts     *execution.HashMap
	typedValue octosql.Value
}

func NewAverage() *Average {
	return &Average{
		averages: execution.NewHashMap(),
		counts:   execution.NewHashMap(),
	}
}

func (agg *Average) Document() docs.Documentation {
	return docs.Section(
		agg.String(),
		docs.Body(
			docs.Section("Description", docs.Text("Averages Floats, Ints or Durations in the group. You may not mix types.")),
		),
	)
}

func (agg *Average) AddRecord(key octosql.Tuple, value octosql.Value) error {
	if agg.typedValue == nil {
		agg.typedValue = value
	}

	var floatValue octosql.Float
	switch value := value.(type) {
	case octosql.Float:
		_, typeOk := agg.typedValue.(octosql.Float)
		if !typeOk {
			return errors.Errorf("mixed types in avg: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}
		floatValue = value
	case octosql.Int:
		_, typeOk := agg.typedValue.(octosql.Int)
		if !typeOk {
			return errors.Errorf("mixed types in avg: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}
		floatValue = octosql.MakeFloat(float64(value.AsInt()))
	case octosql.Duration:
		_, typeOk := agg.typedValue.(octosql.Duration)
		if !typeOk {
			return errors.Errorf("mixed types in avg: %v and %v with values %v and %v",
				execution.GetType(value), execution.GetType(agg.typedValue),
				value, agg.typedValue)
		}
		floatValue = octosql.MakeFloat(float64(value.AsDuration()))
	default:
		return errors.Errorf("invalid type in average: %v with value %v", execution.GetType(value), value)
	}

	count, previousValueExists, err := agg.counts.Get(key)
	if err != nil {
		return errors.Wrap(err, "couldn't get current element count out of hashmap")
	}

	average, _, err := agg.averages.Get(key)
	if err != nil {
		return errors.Wrap(err, "couldn't get current average out of hashmap")
	}

	var newAverage octosql.Float
	var newCount int
	if previousValueExists {
		newCount = count.(int) + 1
		newAverage = (average.(octosql.Float)*octosql.MakeFloat(float64(newCount-1)) + floatValue) / octosql.MakeFloat(float64(newCount))
	} else {
		newCount = 1
		newAverage = floatValue
	}

	err = agg.counts.Set(key, newCount)
	if err != nil {
		return errors.Wrap(err, "couldn't put new element count into hashmap")
	}

	err = agg.averages.Set(key, newAverage)
	if err != nil {
		return errors.Wrap(err, "couldn't put new average into hashmap")
	}

	return nil
}

func (agg *Average) GetAggregated(key octosql.Tuple) (octosql.Value, error) {
	average, ok, err := agg.averages.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get average out of hashmap")
	}

	if !ok {
		return nil, errors.Errorf("average for key not found")
	}

	switch agg.typedValue.(type) {
	case octosql.Duration:
		return octosql.MakeDuration(time.Duration(average.(octosql.Float).AsFloat())), nil
	default:
		return average.(octosql.Float), nil
	}
}

func (agg *Average) String() string {
	return "avg"
}
