class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://github.com/gnufred/logslate'
  version '0.8.0'

  on_linux do
    sha256 'baa7228dd9e8eb3680b623bcdaaa16ad472904e10fa911fad9d99d81337f117e'
    url 'https://github.com/gnufred/logslate/releases/download/0.8.0/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '9c82fc2090ec5715982214ba00f6f83cb8cb288b42baba05a3a1586729850eff'
    url 'https://github.com/gnufred/logslate/releases/download/0.8.0/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
