class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.9'

  on_linux do
    sha256 '053eef6eec1058546d6634245c044591eda9277de52bd7ab98b1eab79ae9616d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.9/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'f984348df18176f482f38620fbaf7548114a06edeb51c32e5776e7b4acf26c68'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.9/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
