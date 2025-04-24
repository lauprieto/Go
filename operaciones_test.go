package main

import "testing"

func TestSumar(t *testing.T) {
    if got := sumar(2, 3); got != 5 {
        t.Errorf("sumar(2, 3) = %d; want 5", got)
    }
}

func TestRestar(t *testing.T) {
    if got := restar(5, 4); got != 1 {
        t.Errorf("restar(5, 4) = %d; want 1", got)
    }
}

func TestMultiplicar(t *testing.T) {
    if got := multiplicar(2, 3); got != 6 {
        t.Errorf("multiplicar(2, 3) = %d; want 6", got)
    }
}
