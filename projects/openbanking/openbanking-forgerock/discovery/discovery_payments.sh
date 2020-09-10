#!/bin/sh -eu

curl -s 'https://rs.aspsp.ob.forgerock.financial:443/open-banking/discovery' | jq '.Data.PaymentInitiationAPI | .[] | select(. | .Version=="v3.1")' > discovery/testdata/discovery_payments.json