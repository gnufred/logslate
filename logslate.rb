class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.8-w5'

  on_linux do
    sha256 '2651f0e618732abd94a0e084c36130eca6f484c157a90a216e06778e68281478'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8-w5/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '0467bf242010079f4f4fac5d840acafa342e151e19739fddb897090fb8f600b9'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.8-w5/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
