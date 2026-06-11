PROBLEM: Producer / Consumer
============================
Run N producer goroutines that push items into a bounded channel of capacity B,
and M consumer goroutines that read and process them. Wait for everything to
drain cleanly, with no leaked goroutines and no deadlocks.

Why it's medium: the classic Go coordination question. Tests `sync.WaitGroup`
discipline, who-closes-the-channel ownership (producers, never consumers), and
that `for range ch` naturally terminates on close.

Example:
  3 producers x 4 items each, 2 consumers, channel buffer 2
  → all 12 items processed, both consumers exit, no deadlock
