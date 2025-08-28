class Notelin < Formula
  desc "No-Tel-in: Telemetry & Sentry Management Scanner"
  homepage "https://github.com/notelin/notelin"
  version "1.0.0"
  
  if OS.mac?
    if Hardware::CPU.arm?
      url "https://github.com/notelin/notelin/releases/download/v1.0.0/notelin-1.0.0-darwin-arm64.tar.gz"
      sha256 "REPLACE_WITH_ARM64_SHA256"
    else
      url "https://github.com/notelin/notelin/releases/download/v1.0.0/notelin-1.0.0-darwin-amd64.tar.gz"
      sha256 "REPLACE_WITH_AMD64_SHA256"
    end
  elsif OS.linux?
    if Hardware::CPU.arm?
      url "https://github.com/notelin/notelin/releases/download/v1.0.0/notelin-1.0.0-linux-arm64.tar.gz"
      sha256 "REPLACE_WITH_LINUX_ARM64_SHA256"
    else
      url "https://github.com/notelin/notelin/releases/download/v1.0.0/notelin-1.0.0-linux-amd64.tar.gz"
      sha256 "REPLACE_WITH_LINUX_AMD64_SHA256"
    end
  end

  depends_on "gum"

  def install
    if OS.mac?
      if Hardware::CPU.arm?
        bin.install "notelin-darwin-arm64" => "notelin"
      else
        bin.install "notelin-darwin-amd64" => "notelin"
      end
    elsif OS.linux?
      if Hardware::CPU.arm?
        bin.install "notelin-linux-arm64" => "notelin"
      else
        bin.install "notelin-linux-amd64" => "notelin"
      end
    end
  end

  test do
    system "#{bin}/notelin", "--version"
    system "#{bin}/notelin", "--help"
  end

  def caveats
    <<~EOS
      No-Tel-in analyzes applications for telemetry and data collection.
      
      Usage:
        notelin                 # Interactive mode
        notelin scan            # Quick system scan  
        notelin report          # Generate privacy report
      
      For more information: https://github.com/notelin/notelin
      
      Note: This tool requires gum for the interactive interface.
      If you encounter issues, ensure gum is properly installed.
    EOS
  end
end