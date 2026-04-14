---
description: Create test cases for a LeetCode problem in Go
---

# Create LeetCode Test

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
   - **Do not provide the solution, approach, or step-by-step algorithm** for the LeetCode problem unless the user explicitly asks for the **source code** (e.g., "直接給我答案", "幫我寫出做法跟答案"). Only write solution code when the user's request clearly and unambiguously asks for the complete implementation.
   - When the user asks for "思路" or "hints", only provide:
     - The **category/pattern** of the problem (e.g., "這題可以往 hash map 的方向想").
     - A **guiding question** to nudge thinking (e.g., "有沒有辦法用 O(n) 的方式把頻率分組？").
     - **Complexity analysis** of the user's current approach.
   - Do NOT provide: full algorithm steps, specific data structure usage walkthrough, or pseudo-code — these count as the answer.
   - Focus only on answering questions about **Go syntax and usage** (e.g., "golang 要如何排序？").

## Capabilities (Can Do / Cannot Do)

### Can Do

- Scaffold a new problem folder following step 1–5 (`problems/pXXXX_snake_case_name/` with `solution.go` stub + `solution_test.go`).
- Write table-driven tests using `reflect.DeepEqual`, normalize order with `sort.Ints` etc. when output ordering is irrelevant.
- Run `go test -v ./problems/pXXXX_.../...` to verify the stub fails as expected.
- Explain Go syntax, standard library usage (slices, maps, sort, container/heap, reflect, etc.).
- Reuse shared types from `kit/` (`ListNode`, `TreeNode`) instead of redefining them per problem.
- Add or extend test cases, fix table entries, adjust comparison/normalization logic.
- Refactor non-solution code (test helpers, kit utilities, workflow docs).
- Offer **minimal** hints when the user asks for a nudge: problem category/pattern, a guiding question, or complexity analysis only — never step-by-step algorithm or specific data structure walkthrough.

### Cannot Do (by workflow rules)

- Provide the full algorithm, step-by-step approach, pseudo-code, or final answer for a LeetCode problem unless the user explicitly asks for the **source code** (e.g., "直接給我答案", "幫我寫出做法跟答案").
- List or compare multiple solution approaches (e.g., "常見有三種做法…") — this effectively gives away the answer.
- Fill in the body of `solution.go` with working logic before being asked.
- Reveal full solutions from other problems in the repo as a direct answer.