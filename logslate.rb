class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.5'

  on_linux do
    sha256 'd0068334c858bbb0a5e9989a6964f28da7f6c84196049f4db515b33c6464b423'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.5/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'ff20a7102a1b2f38f4ba2cee51373c878df982c020836227b54239aea7993d5e'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.5/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
