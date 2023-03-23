## Taint Style Vulnerability

Tainted data may enter the program at specific places and can spread across the program via assignmnets and similat constructs. Using set suitable operations, tainted data can be sanitized, removing its harmful properties. The points to consider:

- Entry points into the program: GET, POST and COOKIE arrays
- Sanitation Routines: the inbuild functions that make it harmless
- Sensitive Sinks: All routines that return data to the browser such as echo(), print() and printf()

## Data flow analysis

Used to determine if the tainted data reaches sensitive sinks without being properly sanitized. The purpose of data flow analysis is to statically compute certain informatioon for every single program point(function). To perform data flow analysis, a control flow graph (CFG) of the program is computed.

Taint analysis is performed on the three adress code representation generated by the front end. Identify where the taint data can enter the program, propagate taint values along assignments and similar constructs and infrom the user of every sink that raceive tainted input.

<b>Alias analysis:</b> Whenever a variable assigned a tained value, the value must not be propageated only to the variable itself but also to all its aliases( variables pointing to the same memory location)

<b>Literal analysis:</b> Know about the literal values that varibles and constants may hold at each program point

<b>Path pruning: </b> used to reduce the number of possible execution paths that a program can take to analyze the code without actually running the program. It reduce number of execution paths. eg,  If a variable is not used, then any paths that rely on that variable can be pruned. Similarly, if a particular condition always evaluates to true or false, then any paths that rely on that condition can be pruned.