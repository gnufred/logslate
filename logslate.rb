class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '00-more_ci'

  on_linux do
    sha256 '7ef8263b0d3f0f7fc7e91978d2a43878347d02c572a0db624f67c5a865dc17d5'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/00-more_ci/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '1c47626a90b93c9cd8f1686fd71903b7efd0ce7e472ab98ab5e199fb570d0976'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/00-more_ci/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
