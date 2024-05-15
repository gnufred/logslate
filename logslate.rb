class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.11'

  on_linux do
    sha256 '62afb4ca5a4c1a23b682e128932575b26dd27f7e56ffd2de6d828691c471074d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.11/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'a4fdf406e5fbd3d8f576e2da33f9bd489febe898823bcd0e438120d8eb5b87e8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.11/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
