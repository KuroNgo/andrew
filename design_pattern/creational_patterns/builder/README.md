## Builder Design Pattern - Reusing an algorithm to create many implementation of an interface

- The Builder pattern helps us construct complex object without directly instantiating their struct writing the logic they require.
Imagine an object that could have dozens of fields that are more complex structs themselves.
Now imagine that you have many objects with these charateristics, and you could have more.
We don't want to write the logic to create all these object in the package that just needs to use the object

### Description
Instance create can be as simple as providing the opening and closing braces {} and leaving the instance with zelo values,
or as complex as an object that needs to make some API calls, check states, and create object for its fields.
You could also have an object that is composed of many object something that's really idiomatic in Go, as it doesn't support inheritance