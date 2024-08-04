package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("Saying hello to people", func(t *testing.T){
        got := Hello("Chris")
        want := "Hello, iChris"

        assertCorrectMessage(t, got, want)
    })
    t.Run ("Saying hello to nobody", func(t* testing.T){
        got := Hello("")
        want := "Hello, World"

        assertCorrectMessage(t,got, want)
    })
}


func assertCorrectMessage(t testing.TB, got, want string){
    t.Helper()
    if got != want{
        t.Errorf("got %q want %q", got, want)
    }
}
