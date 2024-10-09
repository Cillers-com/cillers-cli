#!/bin/sh

set -e

# Name of your app
APP_NAME="cillers-cli"

# Google Cloud Storage bucket name
GCS_BUCKET="cillers-cli"

# Homebrew tap repository (use SSH URL)
HOMEBREW_TAP_REPO="git@github.com:Cillers-com/homebrew-tap.git"
FORMULA_PATH="Formula/cillers.rb"

# Function to extract version from config.go
extract_version() {
    grep 'Version:' config/config.go | sed -E 's/.*Version:[ \t]*"([^"]+)".*/\1/'
}

# Get version
VERSION=$(extract_version)
if [ -z "$VERSION" ]; then
    echo "Failed to extract version. Aborting."
    exit 1
fi

echo "Building version: $VERSION"

# Supported platforms
PLATFORMS="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64"

# Build directory
BUILD_DIR="build"

# Create build directory if it doesn't exist
mkdir -p $BUILD_DIR

# Initialize a string to store checksums
CHECKSUMS=""

# Loop through platforms
for PLATFORM in $PLATFORMS; do
    GOOS=$(echo $PLATFORM | cut -d'/' -f1)
    GOARCH=$(echo $PLATFORM | cut -d'/' -f2)
    
    if [ "$GOOS" = "darwin" ]; then
        OS="macos"
    else
        OS="linux"
    fi
    IMAGE_VARIANT="${OS}-${GOARCH}"
    
    OUTPUT_NAME=$APP_NAME
    
    echo "Building for $GOOS $GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $BUILD_DIR/$OUTPUT_NAME

    ARCHIVE_NAME="${APP_NAME}-${VERSION}-${IMAGE_VARIANT}.tar.gz"
    LATEST_ARCHIVE_NAME="${APP_NAME}-latest-${IMAGE_VARIANT}.tar.gz"
    tar -czf "$BUILD_DIR/$ARCHIVE_NAME" -C "$BUILD_DIR" "$OUTPUT_NAME"
    
    # Calculate checksum
    CHECKSUM=$(shasum -a 256 "$BUILD_DIR/$ARCHIVE_NAME" | awk '{print $1}')
    CHECKSUMS="${CHECKSUMS}${IMAGE_VARIANT}:${CHECKSUM}\n"
    
    rm "$BUILD_DIR/$OUTPUT_NAME"
    
    echo "Created $ARCHIVE_NAME"
    
    # Upload versioned archive to Google Cloud Storage
    echo "Uploading $ARCHIVE_NAME to Google Cloud Storage..."
    gsutil cp "$BUILD_DIR/$ARCHIVE_NAME" "gs://$GCS_BUCKET/$ARCHIVE_NAME"
    
    # Upload and overwrite the "latest" version
    echo "Uploading latest version..."
    gsutil cp "$BUILD_DIR/$ARCHIVE_NAME" "gs://$GCS_BUCKET/$LATEST_ARCHIVE_NAME"
    
    echo "Uploaded:"
    echo "  - gs://$GCS_BUCKET/$ARCHIVE_NAME"
    echo "  - gs://$GCS_BUCKET/$LATEST_ARCHIVE_NAME (Latest stable link)"
done

echo "Build process completed!"

# Update Homebrew formula
echo "Updating Homebrew formula..."

# Clone the homebrew-tap repository
git clone $HOMEBREW_TAP_REPO homebrew-tap
cd homebrew-tap

# Update the formula file
sed -i.bak "s/version \".*\"/version \"$VERSION\"/" $FORMULA_PATH

# Update checksums
echo "$CHECKSUMS" | while IFS=':' read -r variant checksum; do
    sed -i.bak "s/\"$variant\" => \".*\"/\"$variant\" => \"$checksum\"/" $FORMULA_PATH
done

# Remove backup files
rm -f ${FORMULA_PATH}.bak

# Commit and push changes
git config user.name "Cillers CLI Build Script"
git config user.email "noreply@cillers.com"
git add $FORMULA_PATH
git commit -m "Update Cillers CLI to version $VERSION"
git push

cd ..
rm -rf homebrew-tap

echo "Homebrew formula updated successfully!"
echo "Note: Ensure your bucket is configured for public access if you want these files to be publicly downloadable."
