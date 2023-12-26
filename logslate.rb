class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.4.7.1'

  on_linux do
    sha256 'fd4ef507724ec700f668fed3deac5b5d02f3facc4d9ed7c16fdbed59731ce5d5'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7.1/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 'b31775eb566317038e06b0149248c8a209e3aacf7d1baf05eaf5023277dafb8b'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.4.7.1/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
