class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.12'

  on_linux do
    sha256 'f9bd4659c572c22dedb3d40d2c197dc18efce4687fbfa8c18ac93e0fb7f3973d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.12/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '67173af971adcdc4e570ba992251100f8d3b12ab57ce0005eaae248da025b9e6'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.12/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
