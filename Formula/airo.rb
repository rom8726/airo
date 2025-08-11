class Airo < Formula
  desc "Spec-first Golang project scaffold generator"
  homepage "https://github.com/rom8726/airo"
  url "https://github.com/rom8726/airo/archive/refs/tags/v1.12.0.tar.gz"
  sha256 "077ebf011a8acc8c77a69177a68ac4ee46f74079fd704e4874c889cb85c5146f"
  license "Apache-2.0"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/airo"
  end

  test do
    system "#{bin}/airo", "--version"
  end
end
