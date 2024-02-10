class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.2'

  on_linux do
    sha256 'a7ae7baa76219643ce07d244f80afd0dcb6f456fe2001d2b1b8e61769f8cac3f'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.2/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '6c7f26fca4bf26bf7491b1e5eec0a2b53f06c2c5ac195ba85a8577528d589ddf'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.2/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
