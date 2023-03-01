https://vdocuments.mx/secure-programming-with-static-analysis.html?page=177

# Building a model

## Lexical Analysis

- Transforming the code into series of tokens discarding comments and whitespace
- Use regular expression to identify tokens
- ID token need name of the identifier
- For error reporting, token need to contain position on the source text
- It matches the name of the dangerous function

    - Analyzer look on the token stream to look for idetifier
    - Match them against a list of dangerous function names and report results

## Parsing

- Use context free grammer to match the token stream
- Grammer consists of a set of productions that describe the symbols in the language. The parser performs a derivation by matching the token stream against the production rules. If each symbol is connected to the symbol from which it was derived, parse tree is formed.
- Certain types of stylistic check are best performed on a parse tree because it contains the most direst representation of the code just as the programmer wrote it.

## Abstract Syntax

- The datastructure that standardized version of the program suitable for later analysis
- AST can contain a more limited number of constructs than the source language(For eg, for and do loops can be converted to while loops) :<b>lowering</b>
- Languges that are closely related like C and C++ can be lowered into same AST format

## Semantic Analysis

- As the AST is build, also symbol table is created. For each identifier in the program, the symbol table associateds teh identifier with its type and a pointer to its declaration or definition
- With AST and symbol table, tool can perform type checking
- Semantic Analysis : symbolr resolutiona and type checking because now the tool can infer the meaning of the symbol in the progam
- Structural Analysis : tool can perform useful checks on the structure of the program
- Now a modern compiler uses AST and symbol and type information to generate an intermediate representation for optimization and later convert it to the platform specific object code
- Now for a static analyser, AST is transformed or generate its own varity of intermediate representation suitable to its needs

## Tracking Control Flow

- Explore different execution paths that can take place when a function is executed.
- For this, control flow graph on top of AST is created. The <b>basic block</b> contain the sequence of instructions are the nodes of control flow graph.
- The instruction is the basic block not contain actual source code but contain the pointer to the AST nodeor the nodes for the tool's intermediate representation
- A <b>trace</b> is a sequence of basic blocks that define a path through the code
- <b>Call graph</b> represents control flow between functions. Node in the graph represent functions and directed edges represent the potential for one function to invoke another

## Tracking Dataflow

- Examine the way the data move through the program
- It includes traversing a function's control flow graph and noting where data values are generated and where they are used. 
- For this a function is converted to Static Single Assignment(SSA) form. SSA only allow to assign a value to a variable only once. If a variable x is assigned three times, the new program contain the variable x1, x2, x2
- If a SSA variable is ever assigned a constant value, the constant value can replace all uses of the SSA varible : <b>Constant propagation</b>. It help finding security problems such as hard-coded passwords or encryption keys.
- If a variable is assigned different values along different control flow paths in SSA form, the variable must be reconciled at the point where the control flow paths merge using a PHI function

## Taint Propagation

- Identify the values in the program the attacker could control. It use data control to determine what an attacker can control
- Perl's taint mode uses runtime mechanism to make sure the user-supplied data are validated against a regular expression before they are used as part of a sensitive operation

## Pointer Aliasing

- Another dataflow problem
- Used to understant which pointers could possibly refer to the same memory location

# Analysis Algorithm

1. Local Analysis : analyzing individual function (use Control Flow Graph)
2. Global Analysis : analyzing interaction between functions (use Call Graph)

The possibility of buffer overflow by `strcpy(dest, src);` can be removed by making `assert(alloc_size(dest) > strlen(src));` as true. If the condition will be fail somehow, it should be reported. It can be well used for SQLi, XSS etc

- <b>Assertion checking problem </b> : decide what to check and how it can be performed is important

Three varieties of assertions

1. Programmers who trust input, so a tool needs to check assertions related to the level of trust afforded to data as they move through the program
2. Sometime more than one value need to be controlled to succesfully perform the attack - <b>Range analysis</b>
3. Tools are less concerned with he particular data values and more concerned with the type of an object, because varible can have a different type at each point in the code : <b>type state</b>

<b>Program Slicing: </b>removes all the code that cannot affect the outcome of the assert predicate hence by increase efficiency and avoid false paths

## Local Analysis Approaches

- Abstract Interpretation : Abstracting away aspects of the program that are not relevant to the properties of interest and then performing an interpretation
- Flow insensitive analysis : The order the statements are excuted is not taken into account. 
- Predicate Transformers : A weakes precondition of a program is the fewest requirements a function need to reach a final state.
- Model Checking : introduce some specifications to follow and then compares the specification to the model

## Global Analysis

- Inlining : Replacing each function call in the program with the definition of the called function
- Function Summary : When a local analysis alogorithm encounters a functioncall, the function's summary is applied as a stand in for the function

# Rules