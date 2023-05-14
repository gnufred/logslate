class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.0.13'

  on_linux do
    sha256 'c435db95ed6e2f9a716cce5fa97c10fba5cd2be35062c0713d8d8ace588a6951'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.13/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'cb06addb3d436c5929f2a81495ba2788815aeb58daf64af259e14b96e3831626'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.0.13/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
