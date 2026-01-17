# Go Quests


![Go Quests banner](https://images.steamusercontent.com/ugc/964242373533560541/F1BD3729A9743D8D062FE780774044B192356454/?imw=5000&imh=5000&ima=fit&impolicy=Letterbox&imcolor=%23000000&letterbox=false)

<p align="center"><strong>⭐ If these quests helped you learn Go, please consider starring the repo.</strong></p>



A beginner-friendly set of small Go exercises (“quests”). Each quest is intentionally small so you can practice one Go concept at a time.

You implement functions in `solution.go` and confirm correctness with unit tests in `solution_test.go`.

---

### Requirements

- [Golang](https://go.dev/dl/) (install from the official site)
- [Visual Studio Code](https://code.visualstudio.com/download) (install from the official site)
- [VS Code Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [VS Code Readme extension](https://marketplace.visualstudio.com/items?itemName=manishsencha.readme-preview)
- [Git](https://git-scm.com/install/) (optional, but recommended)


Check your Go version: 

```sh
go version
```

---

### Get the repo (two ways)

**Option A: Clone with Git (recommended)**

Repo: <https://github.com/lite-quests/go-quests>

1. Open Visual studio code.
2. File → Open Folder → Select the folder where you want the repository to reside.
2. Toolbar → Terminal → New Terminal

```sh
git clone https://github.com/lite-quests/go-quests.git
cd go-quests # This command helps you go into the quest directory from the terminal
```

**Option B: Download ZIP (no Git required)**
1. Open the repo on GitHub: <https://github.com/lite-quests/go-quests>
2. Click **Code** → **Download ZIP**
3. Unzip the file
4. Visual Studio Code → Open Folder → Select the extracted/unzipped folder
---

### Code structure

- `quests/` — exercises you edit (each quest is a Go package)

- `quests/<n.topic>/README.md` — instructions for that quest

- `quests/<n.topic>/solution.go` — your implementation (edit this)

- `quests/<n.topic>/solution_test.go` — tests (do not change)

- `solutions/` — reference implementations (use to compare after you try)

---

### How to solve a quest (intended workflow)
1. Open the repo folder in VS Code (Or any other IDE of your choice)
2. Pick a quest folder under `quests/` (start with smaller numbers first)
3. Read `quests/<n.topic>/README.md` thoroughly. Review all referenced links and documentation to gain a complete understanding of the underlying concept. If the preview does not render correctly, use **Ctrl + Shift + V** to open the README in preview mode for improved readability.
4. Open the solution file:
   - Example: `quests/10.structs/solution.go`
5. Implement the required functions/methods. 
6. Run the quest tests:

```sh
go test -v ./quests/10.structs
```

6. Repeat until all tests pass. If stuck, inspect the failing test output first.

7. Only after you’ve tried: compare with the reference in `solutions/`.

This repo is designed so that tests teach you the spec: make changes → run tests → refine → repeat.


---

### Tips that save beginners a lot of time

- **Don’t print unless the quest asks**  
  Tests usually check return values or exact stdout bytes.

- **Read test failures carefully**: they usually tell you the exact edge case you missed.

- **Avoid changing test files**
  If tests don’t pass, fix `solution.go` (the tests define the requirements).

- **Disable AI assistance while solving**
  You’ll learn faster if you struggle a bit, read the errors, and look things up in the Go docs. Use AI only after you’ve made a real attempt.

- **Prefer the intended concept over “test hacks”**  
  You might find a workaround that passes the tests, but the goal is to practice the concept the quest is designed to teach (the quest README will usually hint what to use) so please stick to it.


---

### Where are the answers?

- `solutions/` contains reference implementations for each quest.
- Best practice: try first, then compare after you pass (or when you’re stuck).


---

### Troubleshooting

- **“package not found”**: make sure you’re running test command from the correct directory.
- **Tests fail on stdout**: check newline requirements (`fmt.Print` vs `fmt.Println`). 
- **Changes don’t seem to affect tests**
  Clear the test cache:
  ```sh
  go clean -testcache
  ```

---

### Contact

For any issues, contact either:
- [Mani Yadla](https://x.com/mani_yadla_)
- [Ananya Pappula](https://x.com/AnanyaPappula)
