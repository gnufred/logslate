class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.0'

  on_linux do
    sha256 '1e05949071bc797b2fb49d1207f258afa0f3ba4202da1470e62f3567fc6c4a17'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.0/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'ba3db3b743fae09151999ad57e575a8af480c44c85f806ed5b17b2c255a076df'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.0/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
