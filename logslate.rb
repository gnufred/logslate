class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.0'

  on_linux do
    sha256 '5a02b6f224c5b70f16f73f54dbf8517bb0123f622e439a2d94106a27247c8fd7'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.0/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '37d474d89e764662d325c64fda93762f21c111dad3d9f5f84d068c030652185d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.0/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
