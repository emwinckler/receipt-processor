I've not shared a go app before, however I believe the only thing you may have to do to get the app running is
	go get github.com/gorilla/mux
	go get github.com/shopspring/decimal
and     go run main.go


There's likely a simpler way of doing this, however I've been using the command line to POST and GET.
This can be done like so,
	curl -X POST localhost:8000/receipts/process -d '{"retailer": "M&M Corner Market","purchaseDate": "2022-03-20","purchaseTime": "14:33","items": [{"shortDescription": "Gatorade","price": "2.25"},{"shortDescription": "Gatorade","price": "2.25"},{"shortDescription": "Gatorade","price": "2.25"},{"shortDescription": "Gatorade","price": "2.25"}],"total": "9.00"}'
for posting, and like so,
	curl localhost:8000/receipts/1/points
for getting.

The id is simply based on the length of the map that's storing the receipts for simplicity, however it could easily be replaced by an id generator.