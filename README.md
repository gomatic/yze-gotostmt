# yze-go-gotostmt

A [`yze`](https://github.com/gomatic/yze) analyzer (group `go`, category `patterns`) that forbids the `goto` statement, per the gomatic Go standard that control flow uses early returns and extracted helpers rather than `goto`.

- **Rule:** `yze/go/gotostmt`
- **Library:** exports `Analyzer` (a standard `go/analysis` analyzer) and `Registration` for the [`yze`](https://github.com/gomatic/yze) aggregator and [`stickler`](https://github.com/gomatic/stickler) runner.
- **Binary:** `cmd/yze-go-gotostmt` runs it standalone (`text`/`-json`/`-fix`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
