package templates

const GeneralDirectivesTemplate = `<content><directive id="self-documenting-names">
    Use self-documenting variable, function, class and file names. Functions should be named to reflect the output they are expected to produce. Files should be named with the category of functions that it contains. 
</directive>

<directive id="import-organization">
    Import Organization:
    Group imports as follows:
    * Standard library imports
    * Third-party library imports
    * Local application imports
    Within each group, sort imports alphabetically
</directive>

<directive id="package-dependencies">
    Package Dependencies
    Never use deprecated packages
    Prefer standard library packages over third-party alternatives when possible.
</directive>

<directive id="single-responsibility-files">
    Each file must only have one type of function, e.g. getting the current user or defining API endpoints, to improve the organization and maintainability of the code. 
</directive>

<directive id="short-functions">
    Functions MUST be short. Separate concerns by delegating logic to helper functions. 
</directive>

<directive id="consistent-indentation">
    Use consistent indentation with 4 spaces per level. Do not use tabs for indentation.
</directive>

<directive id="error-handling">
    Errors should never be masked. If an error occurs, it must be handled properly, with for example a retry, or the code should fail with a good error message. 
</directive>

<directive id="minimal-variable-mutation">
    Never change a variable, except if there is only one variable in a function and the purpose of that function is to generate the object that is held in the variable. 
</directive>

<directive id="immutable-params">
    Never change a param. Treat all params as frozen objects. 
</directive>

<directive id="no-trivial-comments">
    Never include trivial comments in the code. 
</directive>

<directive id="ensure-functionality">
    Make sure the code works.
</directive>

<directive id="non-zero-exit-on-failure">
    The process must exit with a non-zero code if it was not successful. 
</directive>
</content>`
