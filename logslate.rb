class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.2'

  on_linux do
    sha256 '49ae6e844ab5ffa945a8805d9933644d92a8719671907ff2f0a9d12af7b8e4a6'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'cae4a32423ee127e1043f9a1773189f7586a9de94bc755aa56b52e23fb88d493'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
