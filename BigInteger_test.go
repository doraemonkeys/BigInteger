package BigInteger

import (
	"testing"
)

//go test -cover -coverprofile cover.out
//go tool cover -html cover.out

func TestMultiply(t *testing.T) {
	type test struct {
		input []string
		want  string
	}
	tests := []test{
		{input: []string{"0", "0"}, want: "0"},
		{input: []string{"0", "1"}, want: "0"},
		{input: []string{"1", "1"}, want: "1"},
		{input: []string{"-1", "0"}, want: "0"},
		{input: []string{"-1", "1"}, want: "-1"},
		{input: []string{"-1", "-1"}, want: "1"},
		{input: []string{"8", "8"}, want: "64"},
		{input: []string{"2", "2"}, want: "4"},
		{input: []string{"-2", "-2"}, want: "4"},
		{input: []string{"-2", "2"}, want: "-4"},
		{input: []string{"11", "2"}, want: "22"},
		{input: []string{"-99", "1"}, want: "-99"},
		{input: []string{"100", "2"}, want: "200"},
		{input: []string{"99999", "22"}, want: "2199978"},
		{input: []string{"-8888", "77777777"}, want: "-691288881976"},
		{input: []string{"8888", "-2222"}, want: "-19749136"},
		{input: []string{"9999", "77"}, want: "769923"},
		{input: []string{"8888", "9"}, want: "79992"},
		{input: []string{"8888", "-9"}, want: "-79992"},
		{input: []string{"9876422", "45454"}, want: "448922885588"},
	}
	for _, v := range tests {
		got := BigInteger(v.input[0]).Multiply(BigInteger(v.input[1]))
		if got != BigInteger(v.want) {
			t.Errorf("error,input:{%s %s} got:%s want:%s", v.input[0], v.input[1], got, v.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	type test struct {
		input []string
		want  string
	}
	tests := []test{
		{input: []string{"0", "0"}, want: "0"},
		{input: []string{"0", "1"}, want: "-1"},
		{input: []string{"1", "1"}, want: "0"},
		{input: []string{"-1", "0"}, want: "-1"},
		{input: []string{"-1", "1"}, want: "-2"},
		{input: []string{"-1", "-1"}, want: "0"},
		{input: []string{"8", "8"}, want: "0"},
		{input: []string{"2", "2"}, want: "0"},
		{input: []string{"-2", "-2"}, want: "0"},
		{input: []string{"-2", "2"}, want: "-4"},
		{input: []string{"11", "2"}, want: "9"},
		{input: []string{"-99", "1"}, want: "-100"},
		{input: []string{"100", "2"}, want: "98"},
		{input: []string{"-8", "-9"}, want: "1"},
		{input: []string{"-9", "-8"}, want: "-1"},
		{input: []string{"-9999", "-8"}, want: "-9991"},
		{input: []string{"-8888", "-2222"}, want: "-6666"},
		{input: []string{"8888", "2222"}, want: "6666"},
		{input: []string{"8888", "-2222"}, want: "11110"},
		{input: []string{"9999", "77"}, want: "9922"},
		{input: []string{"8888", "9"}, want: "8879"},
		{input: []string{"8888", "-9"}, want: "8897"},
		{input: []string{"9876422", "45454"}, want: "9830968"},
	}
	for _, v := range tests {
		got := BigInteger(v.input[0]).Subtract(BigInteger(v.input[1]))
		if got != BigInteger(v.want) {
			t.Errorf("error,input:{%s %s} got:%s want:%s", v.input[0], v.input[1], got, v.want)
		}
	}
}

func TestFilp(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "0", want: "0"},
		{input: "1", want: "-1"},
		{input: "-1", want: "1"},
		{input: "9", want: "-9"},
		{input: "9808", want: "-9808"},
		{input: "-7", want: "7"},
		{input: "-8090", want: "8090"},
	}
	for _, v := range tests {
		got := BigInteger(v.input).Flip()
		if got != BigInteger(v.want) {
			t.Errorf("error,input:%s  got:%s want:%s", v.input, got, v.want)
		}
	}
}

func TestGreaterThan(t *testing.T) {
	type test struct {
		input []string
		want  bool
	}
	tests := []test{
		{input: []string{"0", "0"}, want: false},
		{input: []string{"0", "1"}, want: false},
		{input: []string{"1", "1"}, want: false},
		{input: []string{"-1", "0"}, want: false},
		{input: []string{"-1", "1"}, want: false},
		{input: []string{"-1", "-1"}, want: false},
		{input: []string{"8", "8"}, want: false},
		{input: []string{"2", "2"}, want: false},
		{input: []string{"-2", "-2"}, want: false},
		{input: []string{"-2", "2"}, want: false},
		{input: []string{"11", "2"}, want: true},
		{input: []string{"-99", "1"}, want: false},
		{input: []string{"100", "-2"}, want: true},
		{input: []string{"99999", "22"}, want: true},
		{input: []string{"-8888", "77777777"}, want: false},
	}
	for _, v := range tests {
		got := BigInteger(v.input[0]).GreaterThan(BigInteger(v.input[1]))
		if got != v.want {
			t.Errorf("error,input:{%s %s} got:%v want:%v", v.input[0], v.input[1], got, v.want)
		}
	}
}

func TestAdd(t *testing.T) {
	type test struct {
		input []string
		want  string
	}
	tests := []test{
		{input: []string{"0", "0"}, want: "0"},
		{input: []string{"0", "1"}, want: "1"},
		{input: []string{"1", "1"}, want: "2"},
		{input: []string{"-1", "0"}, want: "-1"},
		{input: []string{"-1", "1"}, want: "0"},
		{input: []string{"-1", "-1"}, want: "-2"},
		{input: []string{"8", "8"}, want: "16"},
		{input: []string{"2", "2"}, want: "4"},
		{input: []string{"-2", "-2"}, want: "-4"},
		{input: []string{"-2", "2"}, want: "0"},
		{input: []string{"11", "2"}, want: "13"},
		{input: []string{"-99", "1"}, want: "-98"},
		{input: []string{"100", "2"}, want: "102"},
		{input: []string{"-8", "-9"}, want: "-17"},
		{input: []string{"-9", "-8"}, want: "-17"},
		{input: []string{"-9999", "-8"}, want: "-10007"},
		{input: []string{"-8888", "-2222"}, want: "-11110"},
		{input: []string{"8888", "2222"}, want: "11110"},
		{input: []string{"8888", "-2222"}, want: "6666"},
		{input: []string{"-9999", "77"}, want: "-9922"},
		{input: []string{"8888", "9"}, want: "8897"},
		{input: []string{"8888", "-9"}, want: "8879"},
		{input: []string{"9876422", "45454"}, want: "9921876"},
		{input: []string{"1111", "666666"}, want: "667777"},
	}
	for _, v := range tests {
		got := BigInteger(v.input[0]).Add(BigInteger(v.input[1]))
		if got != BigInteger(v.want) {
			t.Errorf("error,input:{%s %s} got:%s want:%s", v.input[0], v.input[1], got, v.want)
		}
	}
}