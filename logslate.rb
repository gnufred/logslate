# `brew` package management script

class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate/'
  version '0.0.8'

  on_linux do
    sha256 '595861b839c9b9a91a14638fb221808f4384dc979c41eb8ba8b6a43f53debff1'
    url "https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/#{version}/logslate-linux-amd64"
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '762b0df8032a6cec6a6a2e5ca7b49854d153e4a445777f417daa028e421a18b7'
    url "https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/#{version}/logslate-darwin-amd64"
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end

  test do
    logslate version
  end

end
