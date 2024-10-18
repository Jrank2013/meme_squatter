build-squat: clean
	go-assets-builder -p web pkg/templates -o pkg/web/assets.go
	cd pkg/public/; \
		go-assets-builder -p web -v PublicAssets -o ../web/public_assets.go .
	go build -o squat cmd/squat/main.go

clean:
	rm squat || true