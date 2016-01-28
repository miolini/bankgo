FORMAT: 1A
HOST: http://localhost:14080/

# BankGO

BankGO is simple API for stroring and change user balances.

## List All Balances [GET /balances]

Get all users balances.

+ Response 200 (application/json)

        [
            {
                "UserID": 1,
                "Value": 100
            },
            {
                "UserID": 2,
                "Value": 15000,
            }
        ]

## Create transaction [POST /transaction]

Create transaction.

+ Request (application/json)

        {
            "UserID": 1,
            "Value": 100
        }

+ Response 200 (application/json)

        {
            "UserID": 1,
            "Value": 15000
        }
