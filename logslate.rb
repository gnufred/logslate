class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.2'

  on_linux do
    sha256 'ff39e01118a2b8f38ee76dc2412d9e115ea1f56893694070d785df3ea79ab531'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '84ab6bb9b41f80fe64152f1da7a5abc6be919e7a58ef7df511dc033e9e9161d0'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
