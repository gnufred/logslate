class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.1'

  on_linux do
    sha256 '7117ab5916ad8d9a4619ae3e32bc5fdf6f623f4052ffcffe806d1b0fedef861d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'c218d4090a14b69aa623699156f38eca3865d4f68308d2d8c58fcb853c65cd2e'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
