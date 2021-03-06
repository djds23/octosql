package aggregates

import (
	"reflect"
	"testing"

	"github.com/cube2222/octosql"
	"github.com/cube2222/octosql/execution"
)

func TestAverage(t *testing.T) {
	type kv struct {
		key   octosql.Tuple
		value octosql.Value
	}
	tests := []struct {
		name    string
		args    []kv
		key     octosql.Tuple
		want    octosql.Value
		wantErr bool
	}{
		{
			name: "one element",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key")}),
					value: octosql.MakeInt(5),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key")}),
			want: octosql.MakeFloat(5.0),
		},
		{
			name: "one element",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(6.0),
		},
		{
			name: "many single-element groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(4),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(4.0),
		},
		{
			name: "many single-element groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(4),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(6.0),
		},
		{
			name: "many groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(4),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(7),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(8),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(9),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(4.0),
		},
		{
			name: "many groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(4),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(7),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(8),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(9),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(7.5),
		},
		{
			name: "one element",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key")}),
					value: octosql.MakeFloat(5.0),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key")}),
			want: octosql.MakeFloat(5.0),
		},
		{
			name: "one element",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(6.0),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(6.0),
		},
		{
			name: "many single-element groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(6.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(4.0),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(4.0),
		},
		{
			name: "many single-element groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(6.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(4.0),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(6.0),
		},
		{
			name: "many groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(6.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(4.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(7.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(8.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(9.0),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(4.0),
		},
		{
			name: "many groups",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(6.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(4.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(7.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(8.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(9.0),
				},
			},
			key:  octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			want: octosql.MakeFloat(7.5),
		},
		{
			name: "type error",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(4.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(7),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(8.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(9.0),
				},
			},
			key:     octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			wantErr: true,
		},
		{
			name: "type error",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(6),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(4),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(7.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(8.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(9.0),
				},
			},
			key:     octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			wantErr: true,
		},
		{
			name: "type error",
			args: []kv{
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(6.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key1": octosql.MakeInt(1)})}),
					value: octosql.MakeInt(4),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(7.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(8.0),
				},
				{
					key:   octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
					value: octosql.MakeFloat(9.0),
				},
			},
			key:     octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			wantErr: true,
		},
		{
			name:    "average of zero elements",
			args:    []kv{},
			key:     octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1), octosql.MakeTuple([]octosql.Value{octosql.MakeString("key"), octosql.MakeInt(1)}), octosql.MakeObject(map[string]octosql.Value{"key": octosql.MakeInt(1)})}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agg := &Average{
				averages: execution.NewHashMap(),
				counts:   execution.NewHashMap(),
			}
			for i := range tt.args {
				if err := agg.AddRecord(tt.args[i].key, tt.args[i].value); err != nil {
					if !tt.wantErr {
						t.Errorf("Average.AddRecord() error = %v", err)
					}
					return
				}
			}

			got, err := agg.GetAggregated(tt.key)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Average.GetAggregated() error = %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("Average: wanted error")
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Average.GetAggregated() = %v, want %v", got, tt.want)
			}
		})
	}
}
