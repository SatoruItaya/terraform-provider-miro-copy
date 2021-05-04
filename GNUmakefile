version = 99.0.0
provider_macos_path = registry.terraform.io/SatoruItaya/miro/$(version)/darwin_amd64/

install_macos:
	@go build -o terraform-provider-miro_$(version)
	@mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)
	@mv terraform-provider-miro_$(version)  ~/Library/Application\ Support/io.terraform/plugins/$(provider_macos_path)

.PHONY: build test testacc fmt cassettes vet fmtcheck errcheck install_macos
