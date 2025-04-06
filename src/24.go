import math

def calculate_factorial(n):
    if n == 0:
        return 1
    else:
        return n * calculate_factorial(n - 1)

result = calculate_factorial(5)
print(result)  # Output: 120
