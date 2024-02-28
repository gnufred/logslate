class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.3'

  on_linux do
    sha256 'cf3010f8e0317c6a7fe27819da6a840af388404c6d5763d00c04c51680c17d4d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.3/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'c5e34026a84cc88a545d2a37ae01f9ea2d1af8a60c6c397717def72210ca2358'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.3/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
