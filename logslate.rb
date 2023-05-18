class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.17'

  on_linux do
    sha256 '62b78b6b60cf6d02eb361f5493aee2b5d75535d35be3709a5797b2ffff3516eb'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.17/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'c4432139894f7f1e72a4f32e0fd6c247655b76bc71e8530feec69498ea2f88c6'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.17/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
