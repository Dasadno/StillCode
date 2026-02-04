-- StillCode Seed Data
-- Tasks with LeetCode-style function signatures

-- Task 1: Two Sum (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Two Sum',
    E'Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\n**Example 1:**\n```\nInput: nums = [2,7,11,15], target = 9\nOutput: [0,1]\nExplanation: nums[0] + nums[1] = 2 + 7 = 9\n```\n\n**Example 2:**\n```\nInput: nums = [3,2,4], target = 6\nOutput: [1,2]\n```\n\n**Constraints:**\n- 2 <= nums.length <= 10^4\n- -10^9 <= nums[i] <= 10^9\n- Only one valid answer exists.',
    'easy',
    'twoSum',
    '[{"name":"nums","type":"[]int"},{"name":"target","type":"int"}]',
    E'def twoSum(nums, target):\n    # Your code here\n    pass',
    E'function twoSum(nums, target) {\n    // Your code here\n}',
    E'func twoSum(nums []int, target int) []int {\n    // Your code here\n    return nil\n}',
    E'vector<int> twoSum(vector<int>& nums, int target) {\n    // Your code here\n    return {};\n}',
    E'public int[] twoSum(int[] nums, int target) {\n    // Your code here\n    return new int[]{};\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(1, '{"nums": [2,7,11,15], "target": 9}', '[0,1]', FALSE),
(1, '{"nums": [3,2,4], "target": 6}', '[1,2]', FALSE),
(1, '{"nums": [3,3], "target": 6}', '[0,1]', FALSE),
(1, '{"nums": [1,2,3,4,5], "target": 9}', '[3,4]', TRUE),
(1, '{"nums": [-1,-2,-3,-4,-5], "target": -8}', '[2,4]', TRUE);

-- Task 2: Valid Palindrome (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Valid Palindrome',
    E'A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward.\n\nGiven a string `s`, return `true` if it is a palindrome, or `false` otherwise.\n\n**Example 1:**\n```\nInput: s = "A man, a plan, a canal: Panama"\nOutput: true\nExplanation: "amanaplanacanalpanama" is a palindrome.\n```\n\n**Example 2:**\n```\nInput: s = "race a car"\nOutput: false\n```\n\n**Constraints:**\n- 1 <= s.length <= 2 * 10^5\n- s consists only of printable ASCII characters.',
    'easy',
    'isPalindrome',
    '[{"name":"s","type":"string"}]',
    E'def isPalindrome(s):\n    # Your code here\n    pass',
    E'function isPalindrome(s) {\n    // Your code here\n}',
    E'func isPalindrome(s string) bool {\n    // Your code here\n    return false\n}',
    E'bool isPalindrome(string s) {\n    // Your code here\n    return false;\n}',
    E'public boolean isPalindrome(String s) {\n    // Your code here\n    return false;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(2, '{"s": "A man, a plan, a canal: Panama"}', 'true', FALSE),
(2, '{"s": "race a car"}', 'false', FALSE),
(2, '{"s": " "}', 'true', FALSE),
(2, '{"s": "Was it a car or a cat I saw?"}', 'true', TRUE);

-- Task 3: Maximum Subarray (Medium)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Maximum Subarray',
    E'Given an integer array `nums`, find the subarray with the largest sum, and return its sum.\n\n**Example 1:**\n```\nInput: nums = [-2,1,-3,4,-1,2,1,-5,4]\nOutput: 6\nExplanation: The subarray [4,-1,2,1] has the largest sum 6.\n```\n\n**Example 2:**\n```\nInput: nums = [1]\nOutput: 1\n```\n\n**Example 3:**\n```\nInput: nums = [5,4,-1,7,8]\nOutput: 23\n```\n\n**Constraints:**\n- 1 <= nums.length <= 10^5\n- -10^4 <= nums[i] <= 10^4',
    'medium',
    'maxSubArray',
    '[{"name":"nums","type":"[]int"}]',
    E'def maxSubArray(nums):\n    # Your code here\n    pass',
    E'function maxSubArray(nums) {\n    // Your code here\n}',
    E'func maxSubArray(nums []int) int {\n    // Your code here\n    return 0\n}',
    E'int maxSubArray(vector<int>& nums) {\n    // Your code here\n    return 0;\n}',
    E'public int maxSubArray(int[] nums) {\n    // Your code here\n    return 0;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(3, '{"nums": [-2,1,-3,4,-1,2,1,-5,4]}', '6', FALSE),
(3, '{"nums": [1]}', '1', FALSE),
(3, '{"nums": [5,4,-1,7,8]}', '23', FALSE),
(3, '{"nums": [-1]}', '-1', TRUE),
(3, '{"nums": [-2,-1]}', '-1', TRUE);

-- Task 4: Climbing Stairs (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Climbing Stairs',
    E'You are climbing a staircase. It takes `n` steps to reach the top.\n\nEach time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?\n\n**Example 1:**\n```\nInput: n = 2\nOutput: 2\nExplanation: There are two ways: 1+1 and 2.\n```\n\n**Example 2:**\n```\nInput: n = 3\nOutput: 3\nExplanation: There are three ways: 1+1+1, 1+2, and 2+1.\n```\n\n**Constraints:**\n- 1 <= n <= 45',
    'easy',
    'climbStairs',
    '[{"name":"n","type":"int"}]',
    E'def climbStairs(n):\n    # Your code here\n    pass',
    E'function climbStairs(n) {\n    // Your code here\n}',
    E'func climbStairs(n int) int {\n    // Your code here\n    return 0\n}',
    E'int climbStairs(int n) {\n    // Your code here\n    return 0;\n}',
    E'public int climbStairs(int n) {\n    // Your code here\n    return 0;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(4, '{"n": 2}', '2', FALSE),
(4, '{"n": 3}', '3', FALSE),
(4, '{"n": 4}', '5', FALSE),
(4, '{"n": 5}', '8', TRUE),
(4, '{"n": 10}', '89', TRUE);

-- Task 5: Valid Parentheses (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Valid Parentheses',
    E'Given a string `s` containing just the characters ''('', '')'', ''{'', ''}'', ''['' and '']'', determine if the input string is valid.\n\nAn input string is valid if:\n1. Open brackets must be closed by the same type of brackets.\n2. Open brackets must be closed in the correct order.\n3. Every close bracket has a corresponding open bracket of the same type.\n\n**Example 1:**\n```\nInput: s = "()"\nOutput: true\n```\n\n**Example 2:**\n```\nInput: s = "()[]{}"\nOutput: true\n```\n\n**Example 3:**\n```\nInput: s = "(]"\nOutput: false\n```\n\n**Constraints:**\n- 1 <= s.length <= 10^4\n- s consists of parentheses only.',
    'easy',
    'isValid',
    '[{"name":"s","type":"string"}]',
    E'def isValid(s):\n    # Your code here\n    pass',
    E'function isValid(s) {\n    // Your code here\n}',
    E'func isValid(s string) bool {\n    // Your code here\n    return false\n}',
    E'bool isValid(string s) {\n    // Your code here\n    return false;\n}',
    E'public boolean isValid(String s) {\n    // Your code here\n    return false;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(5, '{"s": "()"}', 'true', FALSE),
(5, '{"s": "()[]{}"}', 'true', FALSE),
(5, '{"s": "(]"}', 'false', FALSE),
(5, '{"s": "([)]"}', 'false', TRUE),
(5, '{"s": "{[]}"}', 'true', TRUE);

-- Task 6: Reverse String (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Reverse String',
    E'Write a function that reverses a string. The input string is given as an array of characters `s`.\n\nYou must do this by modifying the input array in-place with O(1) extra memory.\n\n**Example 1:**\n```\nInput: s = ["h","e","l","l","o"]\nOutput: ["o","l","l","e","h"]\n```\n\n**Example 2:**\n```\nInput: s = ["H","a","n","n","a","h"]\nOutput: ["h","a","n","n","a","H"]\n```\n\n**Constraints:**\n- 1 <= s.length <= 10^5\n- s[i] is a printable ASCII character.',
    'easy',
    'reverseString',
    '[{"name":"s","type":"[]string"}]',
    E'def reverseString(s):\n    # Your code here - modify s in-place\n    s.reverse()\n    return s',
    E'function reverseString(s) {\n    // Your code here - modify s in-place\n    return s.reverse();\n}',
    E'func reverseString(s []string) []string {\n    // Your code here\n    return s\n}',
    E'vector<string> reverseString(vector<string>& s) {\n    // Your code here\n    return s;\n}',
    E'public String[] reverseString(String[] s) {\n    // Your code here\n    return s;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(6, '{"s": ["h","e","l","l","o"]}', '["o","l","l","e","h"]', FALSE),
(6, '{"s": ["H","a","n","n","a","h"]}', '["h","a","n","n","a","H"]', FALSE),
(6, '{"s": ["a"]}', '["a"]', TRUE);

-- Task 7: Contains Duplicate (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Contains Duplicate',
    E'Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.\n\n**Example 1:**\n```\nInput: nums = [1,2,3,1]\nOutput: true\n```\n\n**Example 2:**\n```\nInput: nums = [1,2,3,4]\nOutput: false\n```\n\n**Example 3:**\n```\nInput: nums = [1,1,1,3,3,4,3,2,4,2]\nOutput: true\n```\n\n**Constraints:**\n- 1 <= nums.length <= 10^5\n- -10^9 <= nums[i] <= 10^9',
    'easy',
    'containsDuplicate',
    '[{"name":"nums","type":"[]int"}]',
    E'def containsDuplicate(nums):\n    # Your code here\n    pass',
    E'function containsDuplicate(nums) {\n    // Your code here\n}',
    E'func containsDuplicate(nums []int) bool {\n    // Your code here\n    return false\n}',
    E'bool containsDuplicate(vector<int>& nums) {\n    // Your code here\n    return false;\n}',
    E'public boolean containsDuplicate(int[] nums) {\n    // Your code here\n    return false;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(7, '{"nums": [1,2,3,1]}', 'true', FALSE),
(7, '{"nums": [1,2,3,4]}', 'false', FALSE),
(7, '{"nums": [1,1,1,3,3,4,3,2,4,2]}', 'true', FALSE),
(7, '{"nums": [1]}', 'false', TRUE);

-- Task 8: Binary Search (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Binary Search',
    E'Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.\n\nYou must write an algorithm with O(log n) runtime complexity.\n\n**Example 1:**\n```\nInput: nums = [-1,0,3,5,9,12], target = 9\nOutput: 4\nExplanation: 9 exists in nums and its index is 4\n```\n\n**Example 2:**\n```\nInput: nums = [-1,0,3,5,9,12], target = 2\nOutput: -1\nExplanation: 2 does not exist in nums so return -1\n```\n\n**Constraints:**\n- 1 <= nums.length <= 10^4\n- -10^4 < nums[i], target < 10^4\n- All the integers in nums are unique.\n- nums is sorted in ascending order.',
    'easy',
    'search',
    '[{"name":"nums","type":"[]int"},{"name":"target","type":"int"}]',
    E'def search(nums, target):\n    # Your code here\n    pass',
    E'function search(nums, target) {\n    // Your code here\n}',
    E'func search(nums []int, target int) int {\n    // Your code here\n    return -1\n}',
    E'int search(vector<int>& nums, int target) {\n    // Your code here\n    return -1;\n}',
    E'public int search(int[] nums, int target) {\n    // Your code here\n    return -1;\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(8, '{"nums": [-1,0,3,5,9,12], "target": 9}', '4', FALSE),
(8, '{"nums": [-1,0,3,5,9,12], "target": 2}', '-1', FALSE),
(8, '{"nums": [5], "target": 5}', '0', TRUE),
(8, '{"nums": [2,5], "target": 5}', '1', TRUE);

-- Task 9: Fizz Buzz (Easy)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Fizz Buzz',
    E'Given an integer `n`, return a string array `answer` (1-indexed) where:\n\n- `answer[i] == "FizzBuzz"` if i is divisible by 3 and 5.\n- `answer[i] == "Fizz"` if i is divisible by 3.\n- `answer[i] == "Buzz"` if i is divisible by 5.\n- `answer[i] == i` (as a string) if none of the above conditions are true.\n\n**Example 1:**\n```\nInput: n = 3\nOutput: ["1","2","Fizz"]\n```\n\n**Example 2:**\n```\nInput: n = 5\nOutput: ["1","2","Fizz","4","Buzz"]\n```\n\n**Example 3:**\n```\nInput: n = 15\nOutput: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]\n```\n\n**Constraints:**\n- 1 <= n <= 10^4',
    'easy',
    'fizzBuzz',
    '[{"name":"n","type":"int"}]',
    E'def fizzBuzz(n):\n    # Your code here\n    pass',
    E'function fizzBuzz(n) {\n    // Your code here\n}',
    E'func fizzBuzz(n int) []string {\n    // Your code here\n    return nil\n}',
    E'vector<string> fizzBuzz(int n) {\n    // Your code here\n    return {};\n}',
    E'public String[] fizzBuzz(int n) {\n    // Your code here\n    return new String[]{};\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(9, '{"n": 3}', '["1","2","Fizz"]', FALSE),
(9, '{"n": 5}', '["1","2","Fizz","4","Buzz"]', FALSE),
(9, '{"n": 15}', '["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]', FALSE);

-- Task 10: Merge Two Sorted Lists (Medium)
INSERT INTO tasks (title, description, difficulty, function_name, params, starter_code_python, starter_code_js, starter_code_go, starter_code_cpp, starter_code_java)
VALUES (
    'Merge Two Sorted Arrays',
    E'You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order. Merge them into a single sorted array and return it.\n\n**Example 1:**\n```\nInput: nums1 = [1,2,4], nums2 = [1,3,4]\nOutput: [1,1,2,3,4,4]\n```\n\n**Example 2:**\n```\nInput: nums1 = [], nums2 = [0]\nOutput: [0]\n```\n\n**Constraints:**\n- 0 <= nums1.length, nums2.length <= 200\n- -10^9 <= nums1[i], nums2[i] <= 10^9',
    'medium',
    'mergeSortedArrays',
    '[{"name":"nums1","type":"[]int"},{"name":"nums2","type":"[]int"}]',
    E'def mergeSortedArrays(nums1, nums2):\n    # Your code here\n    pass',
    E'function mergeSortedArrays(nums1, nums2) {\n    // Your code here\n}',
    E'func mergeSortedArrays(nums1 []int, nums2 []int) []int {\n    // Your code here\n    return nil\n}',
    E'vector<int> mergeSortedArrays(vector<int>& nums1, vector<int>& nums2) {\n    // Your code here\n    return {};\n}',
    E'public int[] mergeSortedArrays(int[] nums1, int[] nums2) {\n    // Your code here\n    return new int[]{};\n}'
);

INSERT INTO test_cases (task_id, input, expected, is_hidden) VALUES
(10, '{"nums1": [1,2,4], "nums2": [1,3,4]}', '[1,1,2,3,4,4]', FALSE),
(10, '{"nums1": [], "nums2": [0]}', '[0]', FALSE),
(10, '{"nums1": [0], "nums2": []}', '[0]', TRUE);
