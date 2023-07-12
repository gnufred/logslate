class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.3.3'

  on_linux do
    sha256 'ce0f8447664869eea4be706b1a3dd70c9960f1cb6d0994b68f46f2768d153e3e'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.3/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '275fd3efb22d0ca2685f89bd32deced536ae8af87e1b4bacc366905e91b13d07'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.3.3/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
