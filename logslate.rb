class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.7.2'

  on_linux do
    sha256 'e97f6e81edd55d62ab055853cf32a92d5b732fd772aa18a9c082f9e43e68222d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7.2/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'a4842cfda198b7792b2870d67cc239beb0c42cd7b1b003f9f77ba8e17c7817d8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7.2/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
