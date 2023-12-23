class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.5'

  on_linux do
    sha256 '67728d8295deb4724055169a102b8af593462b10774a2c23ab17643a4650721a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.5/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'b5919206626c2524d5f6703fb7e682f7f456aeb7aaef8b6259cb7dd19a6ae573'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.5/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
