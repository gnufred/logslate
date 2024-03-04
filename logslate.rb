class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.7.8'

  on_linux do
    sha256 '1349a3525dd1dfe495cc31dc931c7db8a63e23896c7d56748fcca291a6c9ab47'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.8/logslate-linux-amd64.xz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'a2f39fd65a919eb49ceb75371bcf38e9de9cd1c50cfdd40d1271b3406917abc0'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.7.8/logslate-darwin-amd64.xz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
