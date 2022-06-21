# Домашнее задание к занятию "7.5. Основы golang"

1. Напишите программу для перевода метров в футы (1 фут = 0.3048 метр). Можно запросить исходные данные 
у пользователя, а можно статически задать в коде.
    Для взаимодействия с пользователем можно использовать функцию `Scanf`:

```go
package main

import "fmt"

func main() {
    fmt.Print("Enter value in meters: ")
    var input float64
    var multiplier = 3.28084
    fmt.Scanf("%f", &input)
    
    output := input * multiplier
    
    fmt.Println("Result in feet:", output)
}
```
 
1. Напишите программу, которая найдет наименьший элемент в любом заданном списке, например:
    ```
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    ```
```go
package main

import "fmt"

func main() {
	// input slice
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	// accept the first element as min
	min := x[0]

	for i := 1; i < len(x); i++ {
		// if current element less than min - redefine min variable
		if min > x[i] {
			min = x[i]
		}
	}
	// print min value
	fmt.Println(min)
}
```

1. Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. То есть `(3, 6, 9, …)`.

```go
package main

import "fmt"

func main() {

	// empty slice
	var s []int
	
	for i := 1; i <= 100; i++ {
		// if remainder of a division by 3 equal 0 - append slice s
		if i % 3 == 0 {
			s = append(s, i)
		}		
	}
	// print numbers that are divisible by 3 without a remainder
	fmt.Println(s)
}
```