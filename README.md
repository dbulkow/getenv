I needed a cleaner way to get environment variables with default
values. Cleaner as in the code that uses this library:

~~~~go
var (
        somevar = getenv.Getenv("TOOL_SOMEVAR", "default_value")
)
~~~~
