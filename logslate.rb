class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.6-w2'

  on_linux do
    sha256 '8bb8b6e5fe9811de8838f1e45dadbd6ec6130b10d0187b64d986dc54e943d75b'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.6-w2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '0dc06d72b9d1f879a28253eb8ca4a6c5b21bdcbaa577440a65cb3942155c7fce'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.6-w2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
