class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.9'

  on_linux do
    sha256 '69dc398ec8c44df812df8f8d620b55278beb36ead86b8dd1995cd6db285e29ca'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.9/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '90dc629f11a11d70509b8b6fc846d5075de91c87c826e046de6f4a00cd44550d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.9/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
