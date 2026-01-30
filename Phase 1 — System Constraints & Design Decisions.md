# Phase 1 — System Constraints & Design Decisions

## Purpose

This phase defines the **physical and operational constraints** of the storage engine.  
These choices shape every later decision and are intentionally fixed early.

The goal is not realism for its own sake, but to **model the same pressures real databases face**, while keeping the system small enough to reason about.

---

## Storage Model

### Decision
**In‑memory storage with a simulated disk layer**

All pages are stored in memory, but they are accessed exclusively through a “disk manager” abstraction that exposes `readPage` and `writePage` behavior.

### Reasoning
- Real databases are I/O‑bound, not CPU‑bound
- The important learning is *when* I/O happens, not how fast the OS is
- An in‑memory disk avoids:
  - File system complexity
  - Platform‑specific issues
  - Debugging noise

### Trade‑offs
**Pros**
- Faster iteration
- Easier debugging
- Deterministic behavior

**Cons**
- No real persistence
- No OS‑level caching effects

This is acceptable because persistence is not the learning goal.

---

## Page Size

### Decision
**Page size: 4KB**

### Reasoning
- 4KB is the most common OS page size
- Widely used in real systems
- Small enough to make capacity constraints visible
- Large enough to demonstrate high fan‑out

### Trade‑offs
**Smaller pages**
- More tree levels
- More I/O per lookup

**Larger pages**
- Fewer levels
- Higher write amplification
- More memory waste for small nodes

4KB is a balanced, realistic default.

---

## Key Design

### Decision
**Keys are 64‑bit unsigned integers**

- Fixed‑length
- Unique
- Totally ordered

### Reasoning
- Fixed‑size keys simplify page layout
- No need to deal with variable‑length encoding
- Comparison is cheap and predictable
- 64‑bit range avoids artificial limits

### Trade‑offs
**Pros**
- Simple layout
- Easy capacity math
- Fast comparisons

**Cons**
- Does not model strings or composite keys
- No collation or locale behavior

This project focuses on structure, not data modeling.

---

## Value Design

### Decision
**Leaf pages store fixed‑size values (64‑bit integers)**

Values represent payloads, not pointers to external storage.

### Reasoning
- Keeps the system self‑contained
- Avoids designing a separate heap or record store
- Makes range scans meaningful (keys → values)

### Trade‑offs
**Pros**
- Simplicity
- Clear separation between internal and leaf nodes

**Cons**
- Unrealistic for large records
- Does not model row IDs or tuple storage

This is acceptable because the index behavior is the focus.

---

## Page Types

### Decision
Two explicit page types:
1. **Internal pages**
2. **Leaf pages**

### Internal Pages
- Store only keys and child page IDs
- Used exclusively for navigation

### Leaf Pages
- Store keys and values
- Linked to neighboring leaf pages

### Reasoning
This follows the canonical B+‑tree design used in databases:
- Internal pages stay small and cache‑friendly
- All real data lives at the leaf level
- Range scans are efficient via leaf links

### Trade‑offs
**Pros**
- Excellent range query performance
- Predictable page layouts

**Cons**
- Duplicate keys appear in internal nodes
- Slightly more complex insertion logic

Databases accept this trade‑off for better I/O behavior.

---

## Page Header Design

### Decision
Every page contains a fixed‑size header with:

Common fields:
- Page ID
- Page type (internal or leaf)
- Number of keys stored

Leaf‑only fields:
- Next leaf page ID

### Reasoning
- Pages must be self‑describing
- No external metadata should be required to interpret a page
- This mirrors real on‑disk formats

### Trade‑offs
**Pros**
- Robustness
- Easier debugging and validation

**Cons**
- Header space reduces usable payload

The cost is minimal compared to clarity gained.

---

## Page Capacity Calculations

### Internal Pages

Each internal page stores:
- Header
- N keys
- N+1 child page IDs

This creates a **high fan‑out**, which:
- Keeps tree height low
- Reduces disk reads per lookup

### Leaf Pages

Each leaf page stores:
- Header
- N keys
- N values
- One sibling pointer

### Key Insight
Page size, not algorithmic complexity, determines performance.

This is why B+‑trees work well on disk.

---

## Tree Order

### Decision
Tree order is **derived from page capacity**, not chosen manually.

### Reasoning
- In real systems, order emerges from:
  - Page size
  - Key size
  - Pointer size
- Artificially choosing an order hides real constraints

This keeps the model realistic.

---

## Root Page Special Case

### Decision
- Root page may have fewer keys than other pages
- Root split creates a new root and increases tree height

### Reasoning
- Matches standard B+‑tree behavior
- Simplifies growth logic
- Avoids special handling elsewhere

---

## Expected Tree Height

Given 4KB pages and fixed‑size keys:
- Tree height is expected to be:
  - ~2 levels for thousands of keys
  - ~3 levels for millions of keys

### Implication
Most lookups require:
- 2–4 page reads total

This explains why B+‑trees scale well.

---

## Operational Assumptions

### Decisions
- Single‑threaded execution
- No concurrent readers or writers
- No crashes mid‑operation
- No partial page writes

### Reasoning
These assumptions:
- Reduce scope
- Keep focus on data layout and I/O
- Avoid premature complexity

Concurrency and recovery are separate problems.

---
