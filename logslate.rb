class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.5.4'

  on_linux do
    sha256 '8d78fb7c99c907cf44988bc61efeb40cc189ac803ed4c1b780ccc37e44b523bf'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.4/logslate-linux-amd64.tar.gz'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'b9f609101a184c779be2a3d9fc1df67d1da9994324cfd22733c25d7e647ee4e1'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.5.4/logslate-darwin-amd64.tar.gz'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
