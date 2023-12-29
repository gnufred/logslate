class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.7'

  on_linux do
    sha256 '2f78590a68fe0966b68fc487caedaaa2ac61eacbd13d5bf44cd76f4b5035c922'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.7/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '1f09ce9c4abf05e51647abe1fe688288dc2da3237a90f4e91beb01aaf20c92b2'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.7/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
