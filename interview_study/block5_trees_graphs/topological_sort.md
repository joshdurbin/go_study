PROBLEM: Course Schedule (Topological Sort / Cycle Detection)
=============================================================
There are numCourses courses (0..numCourses-1).
prerequisites[i] = [a, b] means: must take b before a.
Return true if you can finish all courses (graph has no cycle).

Example:
  numCourses=2, prerequisites=[[1,0]]         -> true
  numCourses=2, prerequisites=[[1,0],[0,1]]   -> false (cycle)

Pattern: Kahn's algorithm (BFS topological sort)
  1. Build adjacency list + in-degree count per node
  2. Enqueue all nodes with in-degree 0 (no prerequisites)
  3. BFS: dequeue node, decrement neighbors' in-degree
     If neighbor hits 0, enqueue it
  4. If processed count == numCourses -> no cycle -> true

Why: a cycle means some nodes can never reach in-degree 0 and never get processed.

Time: O(V+E)  Space: O(V+E)
