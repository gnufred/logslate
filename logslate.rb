class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.6.3-t3'

  on_linux do
    sha256 'c4d676d4412956085222aa79c341102c8878a5358f51dbbcb5b108b065ae60a8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3-t3/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '1ca062952ef153bcb7e8131e65c66c386951456d55dd6e2a5a2dd038aae6ae25'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.6.3-t3/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
