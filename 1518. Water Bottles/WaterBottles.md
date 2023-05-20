# Problem 1518
[Link Problem](https://leetcode.com/problems/water-bottles/)

There are `numBottles` water bottles that are initially full of water. You can exchange `numExchange` empty water bottles from the market with one full water bottle.

The operation of drinking a full water bottle turns it into an empty bottle.

Given the two integers `numBottles` and `numExchange`, return the **maximum** *number of water bottles you can drink*.

**Example 1**

```sh
Input: numBottles = 9, numExchange = 3
Output: 13
Explanation: You can exchange 3 empty bottles to get 1 full water bottle.
Number of water bottles you can drink: 9 + 3 + 1 = 13.
```

**Example 2**

```sh
Input: numBottles = 15, numExchange = 4
Output: 19
Explanation: You can exchange 4 empty bottles to get 1 full water bottle. 
Number of water bottles you can drink: 15 + 3 + 1 = 19.
```
# Solution
```sh
- Big O (Time Complexity): O(logn) -> (Bởi vì vòng lặp chạy khi số chai rỗng nhỏ hơn số chai được phép chuyển đổi)
- Big O (Space): O(1) -> (Bởi vì thuật toán chỉ khai báo 2 hằng số)
```
