package structs

import (
    "testing"
    "math"
)

func assertCorrectCalculation(t testing.TB, got, want float64 ){
    t.Helper()
    if got != want{
        t.Errorf("got %g, want %g", got, want)
    }
}

func TestPerimeter(t *testing.T) {
    t.Run("Perimeter rectangle", func(t *testing.T) {
        test_rectangle := Rectangle{10.0,10.0}
        assertCorrectCalculation(t, test_rectangle.Perimeter(), 40.0)
    })

    t.Run("Perimeter circle", func(t* testing.T) {
        test_circle := Circle{10.0}
        assertCorrectCalculation(t, test_circle.Perimeter(), 62.83185307179586)
    })
    
}

func TestArea(t *testing.T) {
    t.Run("Area rectangle", func(t *testing.T) {
        test_rectangle := Rectangle{10.0,10.0}
        assertCorrectCalculation(t, test_rectangle.Area(), 100.0)
    })

    t.Run("Area circle", func(t* testing.T) {
        test_circle := Circle{10.0}
        assertCorrectCalculation(t, test_circle.Area(), 100 * math.Pi) 
    }) 
    
}

