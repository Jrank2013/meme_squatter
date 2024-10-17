buildSquat:
	go-assets-builder -p web pkg/templates -o pkg/web/assets.go
	go build -o squat cmd/squat/main.go