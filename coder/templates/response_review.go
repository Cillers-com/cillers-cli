package templates

const ResponseReviewTemplate = `<instructions>
    <instruction>
        Format your response in the structure specified in the 'content_structure' element. 
    </instruction>
    <instruction>
        Each section should have the title that is specified in the 'title' element. 
    </instruction>
    <instruction>
        Each section should include the information specified in the 'content' element. 
    </instruction>
    <instruction>
        Exclude sections if they have a 'only_include_if' element that is not satisfied.
    </instruction>
    <instruction>
        The files should be ordered in the order they appear in directory tree listing, except files in the root directory that should come first.
    </instruction>
</instructions>
<content_structure>
    <section>
        <h1>
            Summary
        </h1>
        <content>
            Provide a summary of the review, including overall impressions and key findings.
        </content>
    </section>
    <section>
        <h1>
            Directive Violations
        </h1>
        <content>
            List all violations of the directives specified in the 'directive' elements. For each violation, provide:
            * The directive that was violated
            * The file and line number where the violation occurred
            * A brief explanation of how the code violates the directive
            * A suggestion for how to fix the violation
        </content>
        <only_include_if>
            There are directive violations.
        </only_include_if>
    </section>
    <section>
        <h1>
            Code Quality Issues
        </h1>
        <content>
            List any code quality issues that are not direct violations of the specified directives. For each issue, provide:
            * A description of the issue
            * The file and line number where the issue occurs
            * A suggestion for how to improve the code
        </content>
        <only_include_if>
            There are code quality issues.
        </only_include_if>
    </section>
    <section>
        <h1>
            Positive Aspects
        </h1>
        <content>
            Highlight positive aspects of the code, such as:
            * Good adherence to directives
            * Well-structured code
            * Efficient algorithms or implementations
            * Clear and consistent naming conventions
        </content>
    </section>
    <section>
        <h1>
            Suggestions for Improvement
        </h1>
        <content>
            Provide overall suggestions for how the code could be improved, beyond fixing specific violations or issues.
        </content>
    </section>
</content_structure>`
