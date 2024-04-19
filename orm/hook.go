package orm

const (
    HookBefore = "before"
    HookAfter  = "after"
    HookError  = "error"
)

type Hook func(dsn string) error
