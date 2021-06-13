<a name="top"></a>
# Frankie SDK v0.1.0

API Doc Frankie Financial Universal SDK Tester API (Internal Only)

- [Isgood](#isgood)
	- [](#)
	


# <a name='isgood'></a> Isgood

## <a name=''></a> 
[Back to top](#top)

<p>This API allows developers to test the Universal SDK output to ensure it looks right..</p>

	POST /isgood



### Examples

Example usage:

```
curl -i http://localhost:8080/isgood
```
Example usage:

```
	[
		{
  		"checkType": "DEVICE",
  		"activityType": "SIGNUP",
  		"checkSessionKey": "string",
  		"activityData": [
    			{
      			"kvpKey": "ip.address",
      			"kvpValue": "1.23.45.123",
      			"kvpType": "general.string"
    			}
  		]
		}
	]
```


### Success Response

Success-Response:

```
HTTP/1.1 200 Ok
{
  "puppy": true
}
```

### Success 200

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  checkType | String | <p>Describes the type of check service we need to verify with.</p>|
|  activityType | String | <p>The type of activity we're checking.</p>|
|  checkSessionKey | String | <p>The unique session based ID that will be checked against the service.</p>|
|  activityData | Object[] | <p>A collection of loosely typed Key-Value-Pairs, which contain arbitrary data to be passed on to the verification services.</p>|
| &nbsp;&nbsp;&nbsp;&nbsp; activityData.kvpKey | String | <p>Name of the data.</p>|
| &nbsp;&nbsp;&nbsp;&nbsp; activityData.kvpValue | String | <p>Value of the data.</p>|
| &nbsp;&nbsp;&nbsp;&nbsp; activityData.kvpType | String | <p>Used to describe the contents of the KVP data.</p>|

### Error Response

Error-Response:

```
{
  "code": "005",
  "message": "Wrong parameter",
  "error": "CheckType is a required field",
  "status_code": 400
}
```
Error-Response:

```
{
  "code": "005",
  "message": "Wrong parameter",
  "error": "CheckType must be DEVICE | BIOMETRIC | COMBO",
  "status_code": 400
}
```
Error-Response:

```
{
  "code": "005",
  "message": "Wrong parameter",
  "error": "ActivityType must be SIGNUP | LOGIN | PAYMENT | CONFIRMATION | _<Vendor Specific List>",
  "status_code": 400
}
```
