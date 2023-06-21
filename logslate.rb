class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.4'

  on_linux do
    sha256 '66bfcce233cf85e45f4552b5a398f747fa5a69554d13570172240d4981c6c5ed'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.4/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '0854bb5e59b7cb3922cce6c1bfade7f0a190f3767e0f19fd31410ad489de7fdf'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.4/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
