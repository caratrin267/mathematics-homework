def main():
    # Example 1: Calculate Fibonacci sequence
    n = int(input("Enter the number of terms in the Fibonacci sequence (n): "))
    if n <= 0:
        print("Please enter a positive integer.")
    else:
        fib_sequence = [0, 1]
        for i in range(2, n):
            next_fib = fib_sequence[i-1] + fib_sequence[i-2]
            fib_sequence.append(next_fib)
        
        # Output the Fibonacci sequence
        print("Fibonacci sequence:", fib_sequence)

    # Example 2: Calculate powers and roots
    base = float(input("Enter a base for exponentiation (base): "))
    exponent = int(input("Enter an exponent (exponent): "))

    if base < 0 or exponent <= 0:
        print("Please enter valid values.")
    else:
        result = base ** exponent
        print(f"Result of base {base} raised to the power of {exponent} is: {result}")

if __name__ == "__main__":
    main()
