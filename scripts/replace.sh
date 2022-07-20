#!/bin/bash

echo "add import swagger doc"
sed -i -e 's/\/\/swagger_doc_import/\_ \"github\.com\/davidenq\/tweets-timeline-challenge\/docs\/api"/g' ./app/ui/api/server.go
