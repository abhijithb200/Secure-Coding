## Static Analysis Approaches

<b>1. Pattern Matching</b>

Use grep tool to find all occurences of strcpy

<b>2. Lexical Analysis</b>

The source code is turned into stram of tokens, discarding whitespaces and matched against a database of known vulnerability patterns.

<b>3. Parsing and AST analysis</b>

Parsing the source code and create a AST representation. Helps in analyzing not only syntax but also the semantics of a program. The result can be improved by matching AST trees instead of sequences of tokens or characters. Type qualifiers and data-flow analysis can be buil using AST

<b>4. Type qualifiers</b>

Type qualifiers is properties that "qualify" the standard types in languages (const). Use source code to detect type inconsistencies by type qualifier inference.

<b>5. Data flow analysis</b>

<b>6. Taint Analysis</b>

Data from all sources are tainted and propagating its status to all locations where data is used

## Compiler

<b>1. Lexing and Parsing</b>

Lexing 

- removes whitespace and comments
- code to stream of tokens 

Parsing

- analyse syntax of code and check it conforms to language grammer

<b>2. Abstract Syntax Tree</b>

<b>3. Static Single Assignment Form</b>

Create multiple version of a variable one for each assignment and use PHI functions at joint points of multiple versions

<b>4. Data flow analysis</b>

Constant Propagation - find variables that are assigned constants and replace all their uses with a direct use of constant

Any variable have a lattice value : UNDEFINED, VARYING and CONSTANT. All variables are set to UNDEFINED and later set to CONSTANT or VARYING. After a first pass through the code is over and find all the constant. Another pass is performed to replace all their uses with the constant value (Constant Propagation).

<b>
5. Value Range Propagation</b>

Determine the range of possible values for each variable. It find the value range of variables [start,end].

## Solution

<b>1. Taint Analysis</b>

Two lattice values : TAINTED and NOT TAINTED. Any variable that meets a TAINTED value becomes TAINTED.

        NOT TAINTED + NOT TAINTED = NOT TAINTED
        anything    + TAINTED     = TAINTED

The algorithm:

1. Initialize all variables as NOT TAINTED
2. Mark the values returned by those functions which read data from an untrusted source as TAINTED
3. Propagate the tainted values through the program. If a tainted value is used in an expression, mark the result of the expression as TAINTED
4. Repeat step 3 until a fixed point is reached
5. Find all calls to potentially vulnerable functions(like memcpy, strcpy, gets). If one of their arguments is tainted, report this as vulnerability.

```c
1 char src[10], dst[10];
2 int n, m;
3 n = read_int();
4 m = n + 1;
5 memcpy (dst, src, m);
````
n will be marked as tainted and the propagation function will figure out n is used to compute m and hance m is also tainted. Finally the tainted value is used as the argument of vulnerable function memcpy so this is marked as vulnerable.

<b>2. Value Range Propagation</b>

Taint propagation alone is not sufficient. Patterson's value range propagation algorithm can be used to find a range of possible values for each variable.

The algorithm:

1. Use tha value range propagation algorithm to determine a range for each variable
2. Find all calls to function that read data from an untrusted sourve and mark their results as TAINTED
3. Use the taint propagation algorithm to find all tainted values in the program
4. Find all calls to potentially vulnerable functions and inspect their arguments. If an argument is TAINTED and its range is outside the safe range for the function, repott this as a vulnerability 

## Implementation

Contains three parts - Initialization, taint analysis and vulnerability reporting. First, it loads an annotation file (not programme) with vulnerable function and functions returning untrusted data. Four annotations include:

1. __USER_DATA_ : Function returning untrusted data
2. __VULN_USER_ : The parameter is passed untrusted data 
3. __VULN_RANGE : Used for function which is vulnerable if one of their parameters is tainted and outside a safe range
4. __VULN_FUNC_ : Considered vulnerable and appear beginning of the function

```c
ssize_t read (int filedes,void *buffer __USER_DATA__,size_t size);
void * memcpy (void *restrict to, const void *restrict from,
size_t size __VULN_RANGE__);
int printf (const char *template __VULN_USER__, ...);
__VULN_FUNC__ char * gets (char *s __USER_DATA__);
```

After loading annotation and program, it iterates through the statements in the program and find all function calls.

- If the name of the calle matches one of the functons annotated with __USER_DATA, the corresponding variable is marked as TAINTED and perform taint propagation algorithm

- Tainted values used as parameters to __VULN_USER functions are reported as vulnerability

- Report vulnerability if parameter is tainted and its upper bount is equal to MAXINT for __VULN_RANGE annotation