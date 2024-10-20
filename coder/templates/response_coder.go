package templates

const ResponseCoderTemplate = `
<directive>
    Your reponse must comply with the response-directives and be formatted as specified by the response-structure. Do not consider anything else in this prompt when laying out your response. 
</directive>

<response-directive>
    Your response must only contain response-structure with filled in content. You must not prefix this with any text. 
</response-directive>

<response-directive>
    The summary element in the change-description element should:
    Provide a short overal description of the proposed changes.
</response-directive>
    
<response-directive>
    Provide a list that summarizes the changes to be made with a short motivation for why each change should be made.
    Each item in this list should be represented as a change-detail element in the change-description element.
</response-directive>
    
<response-directive>
    Do not include any multi-line code blocks within the change-description section.
</response-directive>

<response-directive>
    All contents inserted into the response-structure should be enclosed in <![CDATA[]]>.
</response-directive>

<response-directive>
    Any "]]" within a CDATA section should be replaced with "]]>]]<!CDATA[>" so the CDATA is not broken. 
</response-directive>

<response-structure>
    <?xml version="1.0" encoding="UTF-8"?>
    <change-proposal>
        <description>
            <change-summary><![CDATA[]]></change-summary>
            <change-detail><![CDATA[]]></change-detail>
        </description>
        <specification>
            <file-to-be-created path=""><![CDATA[]]></file-to-be-created>
            <file-to-be-updated path=""><![CDATA[]]></file-to-be-updated>
            <file-to-be-deleted path="" />
        </specification>
        <code-review>
            <positive_feedback><![CDATA[]]></positive_feedback>
            <improvement_suggestions><![CDATA[]]></improvement_suggestions>
            <code_quality_assessment><![CDATA[]]></code_quality_assessment>
        </code-review>
    </change-proposal>
</response-structure>

                        `
