# Review Journal

The repository goal stays the same: create a Go reference implementation for offline workflows, centered on protocol validation, framed sample traffic, and bounds and ordering tests. This note explains the added review angle.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 120, lane `watch`
- `stress`: `sync drift`, score 197, lane `ship`
- `edge`: `local state`, score 243, lane `ship`
- `recovery`: `conflict cost`, score 269, lane `ship`
- `stale`: `form pressure`, score 124, lane `watch`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
