package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("Saying hello to people", func(t *testing.T){
        got := Hello("Chris", "English")
        want := "Hello, Chris"

        assertCorrectMessage(t, got, want)
    })
    t.Run ("Saying hello to nobody", func(t* testing.T){
        got := Hello("", "English")
        want := "Hello, World"

        assertCorrectMessage(t,got, want)
    })
    t.Run("Saying hello in Uzbek", func(t* testing.T){
        got := Hello("Bahodir", "Uzbek")
        want := "Salom, Bahodir"

        assertCorrectMessage(t, got, want)
    })
    t.Run("Saying hello in empty language", func(t* testing.T){
        got := Hello("Bahodir", "")
        want := "Salom, Bahodir"

        assertCorrectMessage(t, got, want)
    })
}


func assertCorrectMessage(t testing.TB, got, want string){
    t.Helper()
    if got != want{
        t.Errorf("got %q want %q", got, want)
    }
}
