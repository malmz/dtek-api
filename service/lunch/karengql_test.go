package lunch

import "testing"

func TestGql(t *testing.T) {
	resp, err := Test()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
