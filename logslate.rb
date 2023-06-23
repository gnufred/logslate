class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.5'

  on_linux do
    sha256 '6462f744ce2e2d0600bbb0f925e46fa5f9c9fd73617a236cd94b7f03a9308814'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.5/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'b73a21de34d514dd2b92ddb44bdb4811ecf6eaf95da9bc6511747cd91d25c725'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.5/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
