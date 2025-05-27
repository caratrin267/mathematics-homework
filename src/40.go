def gcd(a, b):
    while b != 0:
        temp = b
        b = a % b
        a = temp
    return abs(a)

def lcm(a, b):
    return a * b // gcd(a, b)
