class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.1.0'

  on_linux do
    sha256 ''
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.0/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 ''
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.1.0/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
