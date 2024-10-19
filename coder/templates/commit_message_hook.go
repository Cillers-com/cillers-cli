package templates 

func GenerateCommitMsgHook() string {
    return `#!/bin/sh
COMMIT_MSG_FILE=$1
COMMIT_SOURCE=$2
SHA1=$3

# Only generate commit message for commits created by the user
if [ -z "$COMMIT_SOURCE" ]; then
    STAGED_DIFF=$(git diff --staged | base64)
    API_KEY=$(grep 'api_key:' .cillers/secrets_and_local_config/secrets.yml | awk '{print $2}')
    
    if [ -z "$API_KEY" ]; then
        echo "Error: Anthropic API key not found in .cillers/secrets_and_local_config/secrets.yml" >&2
        exit 1
    fi

    JSON_PAYLOAD=$(cat <<EOF
{
    "model": "claude-3-5-sonnet-20240620",
    "max_tokens": 1024,
    "messages": [
        {
            "role": "user",
            "content": "Generate a Git commit message for the following changes (base64 encoded, please decode before processing):\n\n$STAGED_DIFF"
        }
    ]
}
EOF
)

    RESPONSE_FILE=$(mktemp)
    API_RESPONSE=$(curl -s -w "\n%{http_code}" https://api.anthropic.com/v1/messages \
        -H "x-api-key: $API_KEY" -o "$RESPONSE_FILE" \
        -H "anthropic-version: 2023-06-01" \
        -H "content-type: application/json" \
        -d "$JSON_PAYLOAD")

    HTTP_STATUS=$(echo "$API_RESPONSE" | tail -n1)

    if ! grep -q '"type":"text"' "$RESPONSE_FILE"; then
        echo "Error: Unexpected API response format. Falling back to default editor." >&2
        echo "Response body file: $RESPONSE_FILE" >&2
        rm -f "$RESPONSE_FILE"
        exit 0
    fi

    if [ "$HTTP_STATUS" != "200" ]; then
        echo "Error: API request failed with status $HTTP_STATUS" >&2
        echo "Response body file: $RESPONSE_FILE" >&2
        exit 1
    fi

    COMMIT_MSG=$(sed -n 's/.*"text":"\([^"]*\)".*/\1/p' "$RESPONSE_FILE" | sed 's/\\n/\n/g' | sed 's/\\"/"/g')

    echo "Generated commit message:"
    echo "$COMMIT_MSG"
    echo "Response body file: $RESPONSE_FILE"

    if [ -z "$COMMIT_MSG" ]; then
        echo "Error: Failed to extract commit message from API response. Falling back to default editor." >&2
        echo "Response body file: $RESPONSE_FILE" >&2
        exit 0
    fi

    if [ $(echo "$COMMIT_MSG" | awk 'NF' | wc -l) -eq 0 ]; then
        echo "Error: Generated commit message is empty. Falling back to default editor." >&2
        echo "Response body file: $RESPONSE_FILE" >&2
        exit 0
    fi

    if [ ${#COMMIT_MSG} -gt 2000 ]; then
        echo "Warning: Generated commit message is too long. Truncating to 2000 characters." >&2
        COMMIT_MSG="${COMMIT_MSG:0:2000}"
    fi

    echo "$COMMIT_MSG" > "$COMMIT_MSG_FILE"
    echo "Commit message saved to: $COMMIT_MSG_FILE"
fi
`
}
