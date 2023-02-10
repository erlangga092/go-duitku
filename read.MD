<!-- endpoint sandbox  -->

https://api-sandbox.duitku.com/api/merchant/createInvoice

<!-- request header  -->

Accept: application/json
Content-Type: application/json
x-duitku-signature: FORMAT SHA256(<MERCHANT_CODE_ANDA> - <TIMESTAMP_JAKARTA> - <API_KEY_ANDA>)
x-duitku-timestamp: <TIMESTAMP_JAKARTA>(Milliseconds)
x-duitku-merchantcode: <YOUR_MERCHANT_CODE_HERE>

```json
{
  "paymentAmount": 40000,
  "merchantOrderId": "1648542419",
  "productDetails": "Test Pay with duitku",
  "additionalParam": "",
  "merchantUserInfo": "",
  "customerVaName": "John Doe",
  "email": "test@test.com",
  "phoneNumber": "08123456789",
  "itemDetails": [
    {
      "name": "Test Item 1",
      "price": 10000,
      "quantity": 1
    },
    {
      "name": "Test Item 2",
      "price": 30000,
      "quantity": 3
    }
  ],
  "customerDetail": {
    "firstName": "John",
    "lastName": "Doe",
    "email": "test@test.com",
    "phoneNumber": "08123456789",
    "billingAddress": {
      "firstName": "John",
      "lastName": "Doe",
      "address": "Jl. Kembangan Raya",
      "city": "Jakarta",
      "postalCode": "11530",
      "phone": "08123456789",
      "countryCode": "ID"
    },
    "shippingAddress": {
      "firstName": "John",
      "lastName": "Doe",
      "address": "Jl. Kembangan Raya",
      "city": "Jakarta",
      "postalCode": "11530",
      "phone": "08123456789",
      "countryCode": "ID"
    }
  },
  "callbackUrl": "https://example.com/api-pop/backend/callback.php",
  "returnUrl": "https://example.com/api-pop/backend/redirect.php",
  "expiryPeriod": 10
}
```

<!-- respon  -->

```json
{
  "merchantCode": "DXXXX",
  "reference": "DXXXXS875LXXXX32IJZ7",
  "paymentUrl": "https://app-sandbox.duitku.com/redirect_checkout?reference=DXXXXS875LXXXX32IJZ7",
  "statusCode": "00",
  "statusMessage": "SUCCESS"
}
```

<!-- json object  -->

```json
"itemDetails": [{
    "name": "Apel",
    "quantity": 2,
    "price": 50000
}]

```

<!-- snap frontend  -->

```javascript
<script src="https://app-sandbox.duitku.com/lib/js/duitku.js"></script>
```

<!-- snap  -->

```js
https://app-sandbox.duitku.com/redirect_checkout?reference=DXXXXS875LXXXX32IJZ7&lang=id
```
