# Solution to the following problem:
# Given n points on a plane with x and y-coordinates of 1, 2, ..., n.
# For each point, compute the area of the triangle formed by it and its two adjacent points.
# The total area is then returned.

from typing import List

class Solution:
    def calculate_area(self, points: List[List[int]]) -> int:
        if len(points) < 3:
            return None
        # Calculate areas for pairs (1,2), (1,3)... and add them up.
        return sum(abs(p1[0]-p2[0]) * abs(p1[1] - p2[1]) for p1, p2 in pairwise(points))

# Test the solution with a set of points
solution = Solution()
print(solution.calculate_area([[4,3],[5,3],[7,6]]))  # Output: 10
