class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.9'

  on_linux do
    sha256 '6aef379817473d52622b9dbd5277ccf6e8a5dabcfdbff311dd86f9d9aaaa4497'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.9/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '33d950aefc0706a80ccbd2b39f7bee7098c4b9621400a47902e59505b111d27a'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.9/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
