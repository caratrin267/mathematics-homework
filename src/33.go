def fibonacci(n):
    """
    Generate an infinite sequence of Fibonacci numbers.
    
    Args:
        n (int): The length of the sequence.
        
    Yields:
        int: The next number in the sequence.
    """
    a, b = 0, 1
    for _ in range(n-1):
        yield a
        a, b = b, a + b

for num in fibonacci(10):
    print(num)
