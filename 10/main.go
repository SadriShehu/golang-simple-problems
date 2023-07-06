package main

import (
	"fmt"
	"regexp"
)

func main() {
	matchCountry := regexp.MustCompile(`3ds-([a-z]+).jokrtech.com`)
	fmt.Println(matchCountry.FindStringSubmatch("3ds-ks.jokrtech.com")[1])

	hubRegexp := regexp.MustCompile(`Hub__c/([\w|0-9]+)`)
	orderRegexp := regexp.MustCompile(`Sales_Order__c/([\w|0-9]+)`)
	lineItemRegexp := regexp.MustCompile(`Sales_Order_Line_Item__c/([\w|0-9]+)`)
	data := `{
		"graphs": [
			{
				"graphId": "867305de-5ac6-4df0-8644-5dbf1073b3c4",
				"graphResponse": {
					"compositeResponse": [
						{
							"body": {
								"attributes": {
									"type": "Hub__c",
									"url": "/services/data/v55.0/sobjects/Hub__c/a065e000008CfueAAC"
								},
								"Id": "a065e000008CfueAAC",
								"OwnerId": "0055e000001bpHLAAY",
								"IsDeleted": false,
								"Name": "SCL002",
								"CurrencyIsoCode": "CLP",
								"CreatedDate": "2021-09-03T18:09:05.000+0000",
								"CreatedById": "0055e000003fiKNAAY",
								"LastModifiedDate": "2022-09-15T08:30:25.000+0000",
								"LastModifiedById": "0055e000005sXkvAAE",
								"SystemModstamp": "2022-09-15T08:30:25.000+0000",
								"LastViewedDate": "2022-12-06T16:50:42.000+0000",
								"LastReferencedDate": "2022-12-06T16:50:42.000+0000",
								"Hub_Status__c": "Active",
								"Case_Safe_ID__c": "a065e000008CfueAAC",
								"JOKR_Entity__c": "0015e00000JqxaRAAR",
								"Hub_Name__c": "Holanda",
								"Hub_Street__c": "Holanda 2900",
								"Hub_City__c": "Santiago",
								"Hub_Postal_Code__c": "8320000",
								"Hub_Account__c": "0015e00000Jtm89AAB",
								"Regional_Account__c": "0015e00000JqQ0IAAV",
								"Hub_Country__c": "Chile",
								"NetsuiteId__c": null,
								"Hub_Geolocation__Latitude__s": null,
								"Hub_Geolocation__Longitude__s": null,
								"Hub_Geolocation__c": null,
								"Hub_State__c": null,
								"Hub_Type__c": "Hub",
								"Zone_10__c": null,
								"City_Code__c": "SCL",
								"Safety_Lead_Time__c": null,
								"Zone_1__c": null,
								"Zone_2__c": null,
								"Zone_3__c": null,
								"Zone_4__c": null,
								"Zone_5__c": null,
								"Zone_6__c": null,
								"Zone_7__c": null,
								"Zone_8__c": null,
								"Zone_9__c": null,
								"Area_Account__c": "0015e00000yoL1BAAU"
							},
							"httpHeaders": {
								"Last-Modified": "Thu, 15 Sep 2022 08:30:25 GMT"
							},
							"httpStatusCode": 200,
							"referenceId": "refHub"
						},
						{
							"body": {
								"id": "a0CS000000DYJI6MAP",
								"success": true,
								"errors": []
							},
							"httpHeaders": {
								"Location": "/services/data/v55.0/sobjects/Sales_Order__c/a0CS000000DYJI6MAP"
							},
							"httpStatusCode": 201,
							"referenceId": "refOrder"
						},
						{
							"body": {
								"id": "a0DS000000D4xGUMAZ",
								"success": true,
								"errors": []
							},
							"httpHeaders": {
								"Location": "/services/data/v55.0/sobjects/Sales_Order_Line_Item__c/a0DS000000D4xGUMAZ"
							},
							"httpStatusCode": 201,
							"referenceId": "refLineItem_0"
						}
					]
				},
				"isSuccessful": true
			}
		]
	}`

	btlrProductsResponse := lineItemRegexp.FindAllStringSubmatch(string(data), -1)

	fmt.Println(btlrProductsResponse[0][1])

	hubID := ""
	hubIDRes := hubRegexp.FindStringSubmatch(string(data))
	if len(hubIDRes) > 1 {
		hubID = hubIDRes[1]
	}
	fmt.Println(hubID)

	orderID := ""
	orderIDRes := orderRegexp.FindStringSubmatch(string(data))
	if len(orderIDRes) > 1 {
		orderID = orderIDRes[1]
	}
	fmt.Println(orderID)

}
