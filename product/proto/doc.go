/*
Package proto contains the definition of protobuffer messages and services related to the product service. Note that the
.pb.go and .pb.micro.go files are generated and should not be manually modified.
Note that the product message contains a couple of atypical fields:
- _key: Collection primary key. This is unique field required by the DB and will be auto generated if not provided
- _id: Auto-generated field composed of the collection name and the _key. This is generated by the DB
- _rev: Revision number. This is auto-generated by the DB. Since ArangoDB is a write-only DB, revision number are kept to
	enable delayed record deletion. AQL will always pull the latest version of of a record unless otherwise specified
- extraFields: This field does not directly exist in the DB. It is populated by the service with any fields that are on
- the collection record but not in the protobuf definition. At the time of this writing this is a read only field
*/
package proto
