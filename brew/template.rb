#!/bin/bash

echo "class ${BREW_BIN^} < Formula
  desc '${BREW_DESCRIPTION}'
  homepage 'https://${BREW_GITLAB_URL}'
  version '${BREW_RELEASE_VERSION}'

  on_linux do
    sha256 '${BREW_SHA256_PACKAGE_LINUX}'
    url '${PACKAGE_REGISTRY_CURRENT}/${LINUX_AMD64_BINARY}'
    def install
      bin.install '${LINUX_AMD64_BINARY}' => '${BREW_BIN}'
    end
  end

  on_macos do
    sha256 '${BREW_SHA256_PACKAGE_DARWIN}'
    url '${PACKAGE_REGISTRY_CURRENT}/${DARWIN_AMD64_BINARY}'
    def install
      bin.install '${DARWIN_AMD64_BINARY}' => '${BREW_BIN}'
    end
  end
end"
