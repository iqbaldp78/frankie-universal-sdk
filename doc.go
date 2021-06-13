package main

/**
 * @api {POST} /isgood Validation API
 * @apiDescription This API allows developers to test the Universal SDK output to ensure it looks right..
 * @apiGroup Isgood
 * @apiVersion 0.1.0
 * @apiExample {curl} Example usage:
 *   curl -i http://localhost:8080/isgood
 * @apiExample {json} Example usage:
 * 	[
 * 		{
 *   		"checkType": "DEVICE",
 *   		"activityType": "SIGNUP",
 *   		"checkSessionKey": "string",
 *   		"activityData": [
 *     			{
 *       			"kvpKey": "ip.address",
 *       			"kvpValue": "1.23.45.123",
 *       			"kvpType": "general.string"
 *     			}
 *   		]
 * 		}
 *	]
 *
 * @apiSuccess {String} checkType Describes the type of check service we need to verify with.
 * @apiSuccess {String} activityType The type of activity we're checking.
 * @apiSuccess {String} checkSessionKey The unique session based ID that will be checked against the service.
 * @apiSuccess {Object[]} activityData A collection of loosely typed Key-Value-Pairs, which contain arbitrary data to be passed on to the verification services.
 * @apiSuccess {String} activityData.kvpKey Name of the data.
 * @apiSuccess {String} activityData.kvpValue Value of the data.
 * @apiSuccess {String} activityData.kvpType Used to describe the contents of the KVP data.
 * @apiSampleRequest url
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *   "puppy": true
 * }

 * @apiError wrongParameter The <code>005</code> <field> is a required field.
 * @api {POST} /isgood
 * @apiErrorExample {json} Error-Response:
 * {
 *   "code": "005",
 *   "message": "Wrong parameter",
 *   "error": "CheckType is a required field",
 *   "status_code": 400
 * }

 * @apiError wrongParameter The <code>005</code> CheckType must be DEVICE | BIOMETRIC | COMBO.
 * @api {POST} /isgood
 * @apiErrorExample {json} Error-Response:
 * {
 *   "code": "005",
 *   "message": "Wrong parameter",
 *   "error": "CheckType must be DEVICE | BIOMETRIC | COMBO",
 *   "status_code": 400
 * }

 * @apiError wrongParameter The <code>005</code> ActivityType must be SIGNUP | LOGIN | PAYMENT | CONFIRMATION | _<Vendor Specific List>.
 * @api {POST} /isgood
 * @apiErrorExample {json} Error-Response:
 * {
 *   "code": "005",
 *   "message": "Wrong parameter",
 *   "error": "ActivityType must be SIGNUP | LOGIN | PAYMENT | CONFIRMATION | _<Vendor Specific List>",
 *   "status_code": 400
 * }
 */
