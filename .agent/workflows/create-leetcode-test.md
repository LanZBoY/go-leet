---
description: Create test cases for a LeetCode problem in Go
---

This workflow guides the process of setting up a new LeetCode problem directory with a solution stub and test cases.

1. **Determine Folder Name**
   Identify the problem number (e.g., 347) and descriptive name (e.g., top-k-frequent-elements). 
   Format: `problems/pXXXX_snake_case_name`.
   Example: `problems/p0347_top_k_frequent_elements`.

2. **Create Directory**
   // turbo
   `mkdir -p problems/pXXXX_snake_case_name`

3. **Initialize Solution Stub**
   Create `solution.go` inside the folder with the package name matching the folder name.
   Define the function signature provided by LeetCode.

4. **Initialize Test Cases**
   Create `solution_test.go` with table-driven tests.
   - Use `reflect.DeepEqual` for comparison.
   - If the output order doesn't matter, use `sort.Ints` or similar to normalize results before comparison.
   - Include all examples from the problem description.

5. **Verify Setup**
   // turbo
   `go test -v ./problems/pXXXX_snake_case_name/...`
   Confirm that tests run and fail as expected (since the solution is a stub).

6. **Interaction Rules**
   - **Do not provide the solution or approach** for the LeetCode problem unless the user explicitly asks for it (e.g., "幫我寫出做法跟答案").
   - Focus only on answering questions about **Go syntax and usage** (e.g., "golang 要如何排序？").
