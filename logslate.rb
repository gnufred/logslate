class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.6'

  on_linux do
    sha256 '95ae8c12ef2c955300a349893443062a62d6c263fd42c787e91e3baecfd31b42'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.6/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '3bc433a0e97be93aae8e50455700c7c97688f4d1cd319a5b17be86b73cbb9a13'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.6/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
