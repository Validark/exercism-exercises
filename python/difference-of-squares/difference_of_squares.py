def difference(n):
    return square_of_sum(n) - sum_of_squares(n)

def square_of_sum(n):
    return sum([i + 1 for i in range(n)])**2

def sum_of_squares(n): 
    return sum([(i+1)**2 for i in range(n)])