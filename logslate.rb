class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.0'

  on_linux do
    sha256 'd621b7de61d21265b9eb40aacf0badd2d8e657e3207ecf951a7032ff6271255d'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.0/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '889e1e8f46abcecb81ed0fa79e8d0cad7b4a84e204afe2b5676ebb7cf194a9f8'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.0/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
