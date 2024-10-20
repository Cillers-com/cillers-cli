package templates

const GeneralDirectivesTemplate = `
<directive id="self-documenting-names">
    Use self-documenting variable, function, class and file names. Functions should be named to reflect the output they are expected to produce. Files should be named with the category of functions that it contains. 
</directive>

<response-directive>
    Always respond with the full contents of all code files. Never leave any code out in a file.  
</response-directive>

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

<directive id="max-100-characters-per-line">
    Max 100 characters per line.
</directive>

<directive id="bash-directive">
    Bash scripts must comply with the guidelines specified in the 'bash-directive' elements.
</directive>

<bash-directive id='no-dependencies-on-non-out-of-the-box-tools'>
    Bash script must not depend on any tools that are not available on osx out-of-the-box, including any language that is the source language that installs the bash script. 
</bash-directive>

<bash-directive id="escape-values-in-json">
    When generating JSON, you must escape values to make sure that the JSON doesn't become malformed.
</bash-directive>

<directive id="anthropic-api">
    Any code using the Anthropic Claude Sonnet REST API should be based on the following template which includes the current latest version numbers:

curl https://api.anthropic.com/v1/messages \
     --header &quot;x-api-key: $ANTHROPIC_API_KEY&quot; \
     --header &quot;anthropic-version: 2023-06-01&quot; \
     --header &quot;content-type: application/json&quot; \
     --data \
&apos;{
    &quot;model&quot;: &quot;claude-3-5-sonnet-20240620&quot;,
    &quot;max_tokens&quot;: 1,
    &quot;messages&quot;: [
        {&quot;role&quot;: &quot;user&quot;, &quot;content&quot;: &quot;What is latin for Ant? (A) Apoidea, (B) Rhopalocera, (C) Formicidae&quot;},
        {&quot;role&quot;: &quot;assistant&quot;, &quot;content&quot;: &quot;The answer is (&quot;}
    ]
}&apos;
</directive>
`

