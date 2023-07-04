class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.0'

  on_linux do
    sha256 'ca1faa52e3d1e59de805e954cda576f52df673664e557102d941f87ec674498f'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.0/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'e8c3902eec72fdd47226ec915bbc92555f1ef65d6f01887d0b75d1643c89af7a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.0/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
