build-squat: clean
	cd pkg/templates; \
		go-assets-builder -p web -o ../web/assets.go .
	cd pkg/public/; \
		go-assets-builder -p web -v PublicAssets -o ../web/public_assets.go .
	go build -o squat cmd/squat/main.go

clean:
	rm squat || true