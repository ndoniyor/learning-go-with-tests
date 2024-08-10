package integers

import "testing"

func TestAdd(t *testing.T){
    sum := Add(2,3)
    expected := 5

    if sum != expected{
        t.Errorf("Got sum %d; Expecting %d", sum, expected)
    }
}
