# pprof_GLS

## Functions

### `func pprof_gls.Do(context.Context, pprof.LabelSet, func(context.Context))`

- `pprof.Do()`-alike
- but will fetch back already stacked pprof's labels from GLS if you didn't bother to track your `context.Context` back from your previous ancestor call to `pprof.Do()`
- otherwise, same API / behavior than `pprof.Do()`