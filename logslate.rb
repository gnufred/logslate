class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.7'

  on_linux do
    sha256 '5ccd8d3f114a1f905854a0bd56e782dc8263643a71bf2e8525c359063b7e8789'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.7/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '3372e33a0bfc9d53b7c3cdb90bd40994ae5db3850f96b0fc8853cefe132a9f89'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.7/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
