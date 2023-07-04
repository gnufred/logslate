class Logslate < Formula
  desc 'Logslate (log translate)'
  homepage 'https://gitlab.com/gnufred/logslate'
  version '0.2.10'

  on_linux do
    sha256 'ddd0872f1d02102bc6d7c907d32e230fc09dc0990fddd1d6b691a808175b17ab'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.10/logslate-linux-amd64'
    def install
      bin.install 'logslate-linux-amd64' => 'logslate'
    end
  end

  on_macos do
    sha256 '8e1d00ee43957bac2bbce636b68dcd9a58a992bc005ebe75cb54fed781d06756'
    url 'https://gitlab.com/api/v4/projects/45181636/packages/generic/logslate/0.2.10/logslate-darwin-amd64'
    def install
      bin.install 'logslate-darwin-amd64' => 'logslate'
    end
  end
end
