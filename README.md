# tk-assert

A set of common assertions to reduce the boilerplate of conditionals in tests. Each assertion takes a `*testing.T` object and a condition to check.

## Equality

Takes an interface for both objects to check and compares with `==`

Equal: `tk.Equal(t, 1, 1)`

Not Equal: `tk.NotEqual(t, 1, 2)`

## String comparisons

Checks if a substring exists in a string or not

Contains: `tk.Contains(t, "foo", "foobars")`

Excludes: `tk.Excludes(t, "foo", "doesn't exist")`

## Boolean and nil

Check if the given argument is the expected value

True: `tk.True(t, 1 == 1)`

False: `tk.False(t, 1 == 2)`

Nil: `tk.Nil(t, nil)`

Not Nil: `tk.NotNil(t, "not nil")`

## Errors

Whether an error is thrown, Success is effectively equivalent to Nil but the naming can improve clarity.

Error: `tk.Error(t, errors.New("throwing an error")`

Success: `tk.Success(t, nil)`
