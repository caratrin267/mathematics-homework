package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/alexcg/go-homework"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // Example problem: Find the greatest common divisor (GCD) of two numbers
    a := 24
    b := 36

    result, err := homework.FindGcd(a, b)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("The GCD of %d and %d is: %d\n", a, b, result)
    }

    // Another example problem: Calculate the factorial of a number
    number := 5
    result, err = homework.Factorial(number)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("The factorial of %d is: %d\n", number, result)
    }

    // Another example problem: Sort a list of numbers using insertion sort
    data := []int{8, 6, 2, -4, 5, 1}
    insertSort(data)

    // Sorting algorithm with bubble sort as a reference
    for i := 0; i < len(data)-1; i++ {
        for j := 0; j < len(data)-1-i; j++ {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j]
            }
        }
    }

    // Sorting algorithm with quicksort as a reference
    quickSort(data)

    fmt.Println("Data sorted:", data)
}

func insertSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1

        // Move elements of arr[0..i-1], that are greater than key,
        // to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j + 1] = key
    }
}

func quickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }

    left, right := 0, len(arr)-1

    pivotIndex := left

    for pivotIndex <= right {
        if arr[pivotIndex] > arr[right] {
            swap(&arr[pivotIndex], &arr[right])
            pivotIndex++
        } else {
            pivotIndex++
        }
    }

    // Recursively sort the first part and then the rest
    quickSort(arr[:pivotIndex])
    quickSort(arr[pivotIndex+1:])
}

func swap(a []int, b []int) {
    a[0], b[0] = b[0], a[0]
}
