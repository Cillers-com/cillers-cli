package templates

const LanguageDirectivesGoTemplate = `
<directive>
    Go code must comply with the guidelines specified in the following 'go-directive' elements. 
</directive>

<go-directive id="main-function" tags="go">
    The main Function must:
    * The entry point in the should be the main function in the main.go file.
    * The main function should call another function called 'app' that returns an error if unsuccessful.
    * The main function should exit with a non-zero error if the 'app' function returns an error. 
</directive>

<go-directive id="concise-functions">
    All functions should be as short as possible. Move as much logic as possible out from the function to helper functions.
</go-directive>

<go-directive id="naming-convention">
    Follow Go naming convention. 
</go-directive>

<go-directive id="only-non-deprecated-packages">
    The following Go packages are deprecated and must not be used: 
    * io/ioutil
</go-directive>

<go-directive id="limited-glob-functionality">
    Go doesn't support ** in its glob functionality.
</go-directive>
`
