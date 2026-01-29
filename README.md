# Mini Storage Engine with B+‑Tree Index (Learning Project)

## Goal

Build a small, in‑memory‑backed (disk‑simulated) storage engine to understand:

- How B+‑trees work **as a system component**
- Why databases use pages, buffer pools, and wide trees
- How indexing impacts performance, I/O, and memory
- Trade‑offs senior engineers reason about when using databases

This is **not** a production database and **not** a full DBMS.

---

## Non‑Goals

- No concurrency
- No transactions
- No SQL
- No crash recovery (WAL)
- No replication
- No full deletion correctness

---

## High‑Level Architecture

```
+---------------------+
|   Storage Engine    |
|                     |
|  +---------------+  |
|  |  B+ Tree       |  |
|  |  (Index)       |  |
|  +---------------+  |
|          |           |
|  +---------------+  |
|  | Buffer Pool    |  |
|  +---------------+  |
|          |           |
|  +---------------+  |
|  | Disk Manager   |  |
|  +---------------+  |
+---------------------+
```

---

# OVERALL TODO (Master Checklist)

- [ ] Define system constraints and parameters
- [ ] Implement page abstraction
- [ ] Design page layout (internal vs leaf)
- [ ] Implement disk manager (simulated)
- [ ] Implement buffer pool
- [ ] Implement B+‑tree search
- [ ] Implement insert‑only B+‑tree
- [ ] Implement node splitting logic
- [ ] Link leaf nodes
- [ ] Support range scans
- [ ] Add instrumentation (I/O metrics)
- [ ] Validate invariants
- [ ] Write a short reflection / notes

---

# PART 1: System Constraints & Parameters

### Purpose
Force real‑world limits that shape the design.

### TODO
- [ ] Choose page size (e.g. 4KB or 8KB)
- [ ] Choose key type (fixed‑size integer)
- [ ] Choose value type (fixed‑size payload or pointer)
- [ ] Define maximum tree height expectation
- [ ] Decide: in‑memory only or disk‑backed file

### Questions to Answer
- Why this page size?
- How many keys fit per page?
- How does page size affect tree height?

---

# PART 2: Page Abstraction

### Purpose
Pages are the unit of I/O, not nodes.

### TODO
- [ ] Define a unique page ID system
- [ ] Define common page header fields
- [ ] Distinguish internal vs leaf pages
- [ ] Enforce: one tree node = one page
- [ ] Ensure page size is never exceeded

### Invariants
- A page must fully represent a node
- No pointers to in‑memory structs outside the buffer pool

---

# PART 3: Page Layout Design

### Purpose
Understand how data is physically stored.

### TODO
- [ ] Design internal page layout:
  - keys
  - child page IDs
- [ ] Design leaf page layout:
  - keys
  - values
  - next leaf page ID
- [ ] Decide key ordering rules
- [ ] Decide where free space lives

### Questions to Answer
- Why do internal pages not store values?
- Why do leaf pages need sibling pointers?

---

# PART 4: Disk Manager (Simulated)

### Purpose
Make I/O explicit and measurable.

### TODO
- [ ] Implement page read by page ID
- [ ] Implement page write by page ID
- [ ] Track read/write counts
- [ ] Simulate disk latency (optional)
- [ ] Disallow direct memory access to pages

### Key Insight
Every tree operation must go through this layer.

---

# PART 5: Buffer Pool

### Purpose
Model memory vs disk trade‑offs.

### TODO
- [ ] Fixed‑size buffer pool
- [ ] Page pin/unpin logic
- [ ] Eviction policy (FIFO or LRU)
- [ ] Dirty page tracking
- [ ] Write‑back on eviction

### Questions to Answer
- What happens if the buffer pool is too small?
- Why do reads get faster over time?

---

# PART 6: B+‑Tree Search

### Purpose
Validate traversal logic and page structure.

### TODO
- [ ] Start search from root page
- [ ] Binary search keys inside a page
- [ ] Follow child pointers correctly
- [ ] Always end search in a leaf
- [ ] Verify correctness with test keys

### Invariants
- All leaves are at the same depth
- Internal nodes guide search, never store data

---

# PART 7: Insert‑Only B+‑Tree

### Purpose
Understand structural modification under constraints.

### TODO
- [ ] Insert into leaf if space allows
- [ ] Detect page overflow
- [ ] Split leaf pages
- [ ] Promote separator key to parent
- [ ] Handle cascading splits
- [ ] Handle root split

### Key Difference from B‑Tree
- Separator keys are **copied**, not removed, from leaves

---

# PART 8: Leaf Linking & Range Scans

### Purpose
Enable real database queries.

### TODO
- [ ] Maintain next‑leaf pointers
- [ ] Ensure correct ordering after splits
- [ ] Implement range scan using leaf chain
- [ ] Measure page reads during scan

### Questions to Answer
- Why is this faster than repeated point lookups?
- Why do databases love B+‑trees for ORDER BY?

---

# PART 9: Instrumentation & Metrics

### Purpose
Turn the project into an experiment.

### TODO
- [ ] Count page reads
- [ ] Count page writes
- [ ] Track cache hit/miss ratio
- [ ] Measure tree height
- [ ] Measure keys per page

### Experiments
- Insert sorted keys
- Insert random keys
- Run cold vs warm cache scans

---

# PART 10: Validation & Invariants

### Purpose
Ensure correctness under growth.

### TODO
- [ ] Verify key ordering in all nodes
- [ ] Verify leaf depth uniformity
- [ ] Verify max/min keys per page
- [ ] Verify linked leaf order
- [ ] Detect invariant violations early

---

# PART 11: Reflection (Important)

### Purpose
Convert work into senior‑level insight.

### TODO
- [ ] Write 1–2 pages of notes:
  - What surprised you?
  - What broke first?
  - What trade‑offs became obvious?
- [ ] Answer:
  - Why B+‑trees over B‑trees?
  - Why page size matters?
  - Why range scans are cheap?

---

## Definition of “Done”

You are **done** when:
- Inserts work
- Searches work
- Range scans work
- Metrics make sense
- You can explain *why* the design exists

Not when everything is perfect.

---

## Final Reminder

> **Stop when learning plateaus.**  
> This project is a means, not an identity.

---
