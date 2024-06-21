package standard

import (
	"reflect"
	"testing"

	"github.com/vasilesk/fool/pkg/card"
)

//nolint:funlen
func Test_stdDeck_GetMax(t *testing.T) {
	t.Parallel()

	ordered := newOrdered()

	type fields struct {
		cards []card.Card
		trump card.Card
		pos   int
	}

	type args struct {
		n int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []card.Card
	}{
		{
			name: "three acquired",
			fields: fields{
				cards: ordered.cards,
				trump: ordered.trumpCard,
				pos:   0,
			},
			args: args{n: 3},
			want: []card.Card{
				card.New(card.SuitHearts, card.WeightSix),
				card.New(card.SuitHearts, card.WeightSeven),
				card.New(card.SuitHearts, card.WeightEight),
			},
		},
		{
			name: "all acquired",
			fields: fields{
				cards: ordered.cards,
				trump: ordered.trumpCard,
				pos:   0,
			},
			args: args{n: 36},
			want: ordered.cards,
		},
		{
			name: "too much acquired",
			fields: fields{
				cards: ordered.cards,
				trump: ordered.trumpCard,
				pos:   0,
			},
			args: args{n: 100},
			want: ordered.cards,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := &stdDeck{
				cards:     tt.fields.cards,
				trumpCard: tt.fields.trump,
				pos:       tt.fields.pos,
			}

			got := d.TakeMax(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TakeMax() got = %v, want %v", got, tt.want)
			}
		})
	}
}
