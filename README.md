# notion-mob-offline-lens

`notion-mob-offline-lens` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for offline workflows, centered on protocol validation, framed sample traffic, and bounds and ordering tests.

## Project Rationale

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Notion Mob Offline Lens Review Notes

Start with `conflict cost` and `form pressure`. Those cases create the widest score spread in this repo, so they are the best quick check when the model changes.

## Feature Set

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/notion-mob-offline-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `conflict cost` and `form pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Architecture

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `form pressure`, `sync drift`, `local state`, and `conflict cost`.

The added Go path is deliberately direct, with fixtures doing most of the explaining.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Test Command

The check exercises the source code and the review fixture. `recovery` is the high score at 269; `baseline` is the low score at 120.

## Next Improvements

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
