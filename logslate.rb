class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '20-dotfiles'

  on_linux do
    sha256 'f9d523cba56da4977ec8cfcc0032aae56319ad5d99c0465fe4f01ba6e4d190ed'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/20-dotfiles/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '0e7becf3a60cee0929130457fd5744b633c185727f89eea76e68e9b89690a763'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/20-dotfiles/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
