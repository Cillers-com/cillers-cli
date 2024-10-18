package templates

const ResponseCoderTemplate = `<instructions>
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
    <instruction>
        'list' elements have a sub element 'of' that specifies which elements the list should contain. 
    </instruction>
    <instruction>
        The h1, h2 and h3 elements should be equivalent to HTML header tags. h1 should have the largest bold text, h2 should be slightly smaller and h3 should be bold but same size as normal text. 
    </instruction>
</instructions>
<content_structure>
    <section>
        <h1>
            Summary Of Changes
        </h1>
        <content>
            Provide a list that summarizes the changes to be made with a short motivation for why each change should be made.
            Do not include any multi-line code blocks in this section. 
        </content>
    </section>
    <section>
        <h1>
            Files To Be Added
        </h1>
        <content>
            For each file that should be added provide:
            * the path of the file formatted as a markdown header
            * a motivation for why the file should be added with "Motivation" as sub-header. 
            * the full code in a code block with "Full Code" as sub-header.
        </content>
        <only_include_if>
            There are files that should be added.
        </only_include_if>
    </section>
    <section>
        <h1>
            Files To Be Changed
        </h1>
        <content>
            <list>
                <of>All files tht should be changed</of>
                <list_item>
                    <h2>
                        The path of the file.
                    </h2>
                    <h3>
                        Change List
                    </h3>
                    <content>
                        A bullet list with high-level descriptions of the changes to be made in this file with a motivation.
                    </content>
                    <h3>
                        Diff
                    </h3>
                    <content>
                        A git-style code diff in a code block without code that didn't change.
                    </content>
                    <h3>
                        Full Code
                    </h3>
                    <content>
                        The full code in a code block. This should not just be a diff of the code. It should be the entire contents of the file after the update. 
                    </content>
                </list_item>
            </list>
        </content>
        <only_include_if>
            There are files that should be changed.
        </only_include_if>
    </section>
    <section>
        <h1>
            Files To Be Removed
        </h1>
        <content>
            For each file that should be removed provide:
            * the path of the file formatted as a markdown header
            * a motivation for why the file should be removed.
        </content>
        <only_include_if>
            There are files that should be removed.
        </only_include_if>
    </section>
    <section>
        <h1>
            Peer Review
        </h1>
        <content>
            Provide a likely peer review of a developer that is keen on engineering excellence, simplicity, and readability. Provide one section with positive feedback, and one section with what could have been done better. Are the namings the files, functions and variables self-documenting. 
        </content>
    </section>
</content_structure>`
