package main

// #include "calculator.c"
import "C"
import (
	"errors"
	"fmt"
)

func main() {
  var x float64 = 5
  var y float64 = 8

  addition, err := additionC(x, y)
  if err != nil {
    panic(err)
  }
  subtract, err := subtractC(x, y)
  if err != nil {
    panic(err)
  }
  multiply, err := multiplyC(x, y)
  if err != nil {
    panic(err)
  }
  divide, err := divideC(x, y)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%.2f + %.2f = %.2f\n", x, y, addition)
  fmt.Printf("%.2f - %.2f = %.2f\n", x, y, subtract)
  fmt.Printf("%.2f * %.2f = %.2f\n", x, y, multiply)
  fmt.Printf("%.2f / %.2f = %.2f\n", x, y, divide)
}

func additionC(x float64, y float64) (float64, error) {
  xC := C.double(x)
  yC := C.double(y)

  sum, err := C.addition(xC, yC)
  if err != nil {
    return 0, errors.New("failed to call addition: " + err.Error())
  }

  return float64(sum), nil
}

func subtractC(x float64, y float64) (float64, error) {
  xC := C.double(x)
  yC := C.double(y)

  sum, err := C.subtract(xC, yC)
  if err != nil {
    return 0, errors.New("failed to call subtract: " + err.Error())
  }

  return float64(sum), nil
}

func multiplyC(x float64, y float64) (float64, error) {
  xC := C.double(x)
  yC := C.double(y)

  sum, err := C.multiply(xC, yC)
  if err != nil {
    return 0, errors.New("failed to call multiply: " + err.Error())
  }

  return float64(sum), nil
}

func divideC(x float64, y float64) (float64, error) {
  xC := C.double(x)
  yC := C.double(y)

  sum, err := C.divide(xC, yC)
  if err != nil {
    return 0, errors.New("failed to call divide: " + err.Error())
  }

  return float64(sum), nil
}

