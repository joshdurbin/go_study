PROBLEM: Dijkstra's Shortest Path
=================================
Given a weighted directed graph (non-negative weights) and a source, compute
shortest distances from source to every reachable node.

Why this block: BFS handles unweighted graphs; for weights, BFS fails because
short hops can win over long single edges. Dijkstra = BFS + a min-heap keyed
by tentative distance. Pop the smallest, relax outgoing edges, repeat.

Example:
  Edges: 0→1 (w=4), 0→2 (w=1), 2→1 (w=2), 1→3 (w=1), 2→3 (w=5)
  From 0: dist = [0, 3, 1, 4]
