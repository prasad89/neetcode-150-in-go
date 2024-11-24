# Neetcode 150 Solutions in GoLang

This repository contains my solutions to the **Neetcode 150** problems, implemented in **GoLang**. It serves as a structured approach to mastering data structures, algorithms, and problem-solving for technical interviews.

## Topics Covered

### Arrays & Hashing

- [Contains Duplicate](./Arrays%20%26%20Hashing/containsDuplicate.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Check if a given array contains duplicate elements.
    - **Approach**:
      1. Use a hash set to keep track of seen elements.
      2. Iterate through the array; if an element is already in the set, return `true`.
      3. Otherwise, add the element to the set.
      4. Return `false` if no duplicates are found.
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    
  </details>

- [Valid Anagram](./Arrays%20%26%20Hashing/validAnagram.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Determine if two strings are anagrams of each other.
    - **Approach**:
      1. Count the frequency of characters in the first string.
      2. Decrement the frequency counts using characters from the second string.
      3. If the frequency of any character becomes negative, return `false`. Otherwise, return `true`.
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(1)
    
  </details>

- [Two Sum](./Arrays%20%26%20Hashing/twoSum.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Find indices of two numbers that add up to a given target.
    - **Approach**:
      1. Use a hash map to store the difference of each number and the target.
      2. If the current number exists in the hash map, return its index (stored as value in hash map) along with the current index.
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    
  </details>

- [Group Anagrams](./Arrays%20%26%20Hashing/groupAnagrams.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Group strings that are anagrams of each other.
    - **Approach**:
      1. Create a map where the key is an array of size 26 representing character frequencies (one slot for each letter 'a' to 'z'), and the value is a list of strings.
      2. For each string in the input array:
         - Create a frequency count array for the string by counting occurrences of each character (based on ASCII offsets).
         - Use this array as the key and append the string to the corresponding value in the map.
      3. Iterate through the map and collect all grouped anagrams into the result.
      4. Return the grouped anagrams as a list of lists.
    - **Time Complexity**: O(n * m)
    - **Space Complexity**: O(n * m)
    
  </details>

- [Top K Frequent Elements](./Arrays%20%26%20Hashing/topKFrequent.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Return the k most frequent elements from an array.
    - **Approach**:
      1. Use a hash map to count the frequency of each number in the array.
      2. Create a 2D array (bucket) where the index represents the frequency, and each entry contains the numbers with that frequency.
      3. Iterate through the 2D array from the highest frequency (end) to the lowest, collecting numbers into the result until k elements are gathered.
      4. Return the result containing the top k frequent elements.
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    
  </details>

- [Encode and Decode Strings](./Arrays%20%26%20Hashing/encodingDecoding.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Implement methods to encode and decode a list of strings.
    - **Approach**:
      1. **Encoding**:
         - For each string, append its length, a delimiter (`#`), and the string itself to the encoded result.
      2. **Decoding**:
         - Use two pointers, `i` and `j`, to traverse the encoded string.
         - Move `j` until the delimiter (`#`) is found, parse the substring between `i` and `j` to determine the length of the next string.
         - Extract the substring of this length, append it to the decoded list, and update `i` to point to the next encoded segment.
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    
  </details>

- [Product of Array Except Self](./Arrays%20%26%20Hashing/productOfArrayExceptSelf.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Return an array where each element is the product of all elements except itself.
    - **Approach**:
      1. Compute prefix and postfix products in two separate passes.  
      2. During the postfix pass, multiply the postfix product with the corresponding prefix product to build the result array.  
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(1)
    
  </details>

- [Valid Sudoku](./Arrays%20%26%20Hashing/validSudoku.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Check if a Sudoku board is valid.
    - **Approach**:
      1. Use hash sets to track elements in rows, columns, and 3x3 sub-boxes.
      2. **Traverse the Sudoku board**:
        - For each value at position `(r, c)`, check if that value already exists in the respective row, column, or sub-box using the hash sets.
        - If the value is already present in any of these, return `false`.
      3. **For the sub-boxes**:
        - The index of the sub-box can be calculated using the formula: `(r/3)*3 + c/3`.
      4. If no conflicts are found while traversing the entire board, return `true`.
    - **Time Complexity**: O(1)
    - **Space Complexity**: O(1)
    
  </details>

- [Longest Consecutive Sequence](./Arrays%20%26%20Hashing/longestConsecutiveSequence.go)
  <details>
    <summary>Approach</summary>
    
    - **Problem Statement**: Find the length of the longest consecutive sequence in an array.
    - **Approach**:
      1. **Use a hash set** to store unique elements from the array.
      2. For each element, check if it starts a new sequence (i.e., `num-1` is not present in the set). If it does, initialize `length = 1`.
      3. Continue iterating through the array, checking if `num + length` exists in the set. If found, increment `length`.
      4. If a sequence is not found, break the loop and update the longest sequence length if the current sequence is longer.
    - **Time Complexity**: O(n)
    - **Space Complexity**: O(n)
    
  </details>

### Two Pointers

- [Valid Palindrome](./Two%20Pointers/validPalindrome.go)
- [Two Sum II Input Array Is Sorted](./Two%20Pointers/twoSumII.go)
- [3Sum](./Two%20Pointers/threeSum.go)
- [Container With Most Water](./Two%20Pointers/containerWithMostWater.go)
- [Trapping Rain Water](./Two%20Pointers/trappingRainWater.go)

## Getting Started

Each topic will have its own folder, with GoLang files for individual problems. To run a solution:

```bash
go run path/to/solution.go
```

## Status

This repository is currently private, but it will be made public in the future to share my solutions and help others studying for technical interviews.

---

_Stay tuned for updates!_
