class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.9'

  on_linux do
    sha256 'c8f4a2e6830fe07019213d00aa4ba0588811b3ab2e8542bdca7554aab73d38c9'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.9/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '6d1c07c35385318aac9909b036dd678ab24b09870bf8dcab0fa82fcd4a6fc8dd'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.9/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
