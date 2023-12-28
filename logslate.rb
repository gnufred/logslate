class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.5'

  on_linux do
    sha256 '0a3df278e533c0f20612f30983ed139578dc50c72063163286b7f16c148a1633'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.5/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '5d5c531bbaa2fb962103d0833bcc2b55df15ba930dbbfedc774a00fb58dc73f8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.5/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
