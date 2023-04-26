# `brew` package management script

class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate/'
  version '00-ci_upgrades'

  on_linux do
    sha256 '8cf269956accbab8715bb226675b4945c25cd0711e10d7d3c7076e8b33f52cab'
    url "https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/#{version}/logslate-linux-amd64"
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '4f8d83c134df6a40c7908dd4669899363ddc27552962ff4347b1de84a1c15fde'
    url "https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/#{version}/logslate-darwin-amd64"
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end

  test do
    logslate version
  end

end
