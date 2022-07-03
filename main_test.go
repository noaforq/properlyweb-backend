package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 0. 実装されている
// 1. 引数に取った値を文字列で返す。
// 2. ただし、３の倍数のときは ｢Fizz｣ を返す。
// 3. ただし、５の倍数のときは ｢Buzz｣ を返す。
// 4. ただし、３と５両方の倍数の場合には ｢FizzBuzz｣ を返す。
// 5. 引数が１から１００までの数でない場合は例外にする。

// 0. 実装されている
func Test実装されている(t *testing.T) {
	FizzBuzz(1)
}

// 1. 引数に取った値を文字列で返す。
func Test引数に取った値を文字列で返す(t *testing.T) {
	var result string = FizzBuzz(1)
	assert.Equal(t, "1", result)
}

// 2. ただし、３の倍数のときは ｢Fizz｣ を返す。
func Test３の倍数のときはFizzを返す(t *testing.T) {
	var result string = FizzBuzz(3)
	assert.Equal(t, "Fizz", result)
}

// 3. ただし、５の倍数のときは ｢Buzz｣ を返す。
func Test５の倍数のときはBuzzを返す(t *testing.T) {
	var result string = FizzBuzz(5)
	assert.Equal(t, "Buzz", result)
}

// 4. ただし、３と５両方の倍数の場合には ｢FizzBuzz｣ を返す。
func Test３と５両方の倍数の場合にはFizzBuzzを返す(t *testing.T) {
	var result string = FizzBuzz(15)
	assert.Equal(t, "FizzBuzz", result)
}

// 5. 引数が１から１００までの数でない場合は例外にする。
func Test引数が１から１００までの数でない場合は例外にする(t *testing.T) {
	defer func() {
		err := recover()
		assert.Equal(t, "argument exception", err)
	}()
	FizzBuzz(0)
}
