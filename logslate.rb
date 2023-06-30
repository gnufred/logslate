class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.9'

  on_linux do
    sha256 '53768da6e8f7854f8decade92d1784792e994ed0a398d17e6b5a715d98fa7566'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.9/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '519fdf8259520ce5fe0300202504bbea2675e0b5a5852c10c4f6cd3fc3f4992a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.9/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
