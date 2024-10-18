package templates

const IgnoreTemplate = `# Ignore the .cillers directory itself
.cillers

# Ignore common version control directories
.git
.svn

# Ignore build artifacts and dependencies
/build
/dist
/node_modules
/vendor

# Ignore log files
*.log

# Ignore system and hidden files
.DS_Store
Thumbs.db

# Ignore IDE and editor specific files
.vscode
.idea
*.swp
*.swo

# Ignore compiled binaries
*.exe
*.dll
*.so
*.dylib

# Ignore temporary files
*.tmp
*.temp
`
