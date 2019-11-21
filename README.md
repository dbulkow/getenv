I needed a cleaner way to get environment variables with default
values.

~~~~go
var (
        somevar = getenv.Getenv("TOOL_SOMEVAR", "default_value")
)
~~~~

One might also have to process many variables with the same prefix

~~~go
env := getenv.NewEnv("TOOL")
var somevar = env.Getenv("SOMEVAR", "default_value")
~~~
