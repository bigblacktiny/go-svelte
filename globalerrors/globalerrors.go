package globalerrors

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type ValidationError struct {
	Source      string
	FailureDesc []string
}

func (v *ValidationError) Error() string {
	var failureDesc string
	for _, desc := range v.FailureDesc {
		failureDesc += desc
	}
	return fmt.Sprintf("validation error in %s:\n %s ", v.Source, failureDesc)
}

type SrvError string

const (
	srvNoStartTxt               SrvError = "Unable to start %s server. Error: %v \n"
	srvNoHandlerTxt             SrvError = "Unable to register service handler. Error: %v"
	dbNoConnectionTxt           SrvError = "Unable to connect to DB %s. Error: %v\n"
	dbNoConnectionStringTxt     SrvError = "Unable to find DB connection string. Please set environment variable %s \n"
	dtProtoTimeStampToTimeStamp SrvError = "Unable to convert proto timestamp %v to timestamp. Error: %v \n"
	dtTimeStampToProtoTimeStamp SrvError = "Unable to convert timestamp %v to proto timestamp. Error: %v \n"
	dtInvalidValidityDates      SrvError = "The valid thru date (%v) must take place after the valid from date (%v)\n"
	missingField                SrvError = "%s must not be empty\n"
	authNoMetaData              SrvError = "Unable to read meta-date for end point: %s\n"
	authInvalidToken            SrvError = "invalid token\n"
	authNilToken                SrvError = "Invalid nil user token\n"
	authNilClaim                SrvError = "invalid nil %s claim\n"
	authInvalidClaim            SrvError = "Invalid %s claim\n"
	brkBadMarshall              SrvError = "Unable to marshall object %v. Error: %v\n"
	brkNoMessageSent            SrvError = "Unable to send message to broker for topic %s. Error: %v\n"
	brkNoConnection             SrvError = "Unable to connect to broker: Error: %v\n"
	brkUnableToSetSubs          SrvError = "Unable to setup broker subscription for topic %s. Error: %v\n"
	brkBadUnMarshall            SrvError = "Unable to unmarshall received object for topic %s. Message received: %v. Error: %v\n"
)

func (ge *SrvError) SrvNoStart(serviceName string, err error) string {
	return fmt.Sprintf(string(srvNoStartTxt), serviceName, err)
}

func (ge *SrvError) DbNoConnection(dbName string, err error) string {
	return fmt.Sprintf(string(dbNoConnectionTxt), dbName, err)
}

func (ge *SrvError) DbNoConnectionString(envVarName string) string {
	return fmt.Sprintf(string(dbNoConnectionStringTxt), envVarName)
}

func (ge *SrvError) SrvNoHandler(err error) string {
	return fmt.Sprintf(string(srvNoHandlerTxt), err)
}

func (ge *SrvError) DtProtoTimeStampToTimeStamp(currTimeStamp *timestamp.Timestamp, err error) string {
	return fmt.Sprintf(string(dtProtoTimeStampToTimeStamp), currTimeStamp, err)
}

func (ge *SrvError) DtTimeStampToProtoTimeStamp(currentTime time.Time, err error) string {
	return fmt.Sprintf(string(dtTimeStampToProtoTimeStamp), currentTime, err)
}

func (ge *SrvError) MissingField(fieldName string) string {
	return fmt.Sprintf(string(missingField), fieldName)
}

func (ge *SrvError) DtInvalidValidityDates(validFrom, validThru time.Time) string {
	return fmt.Sprintf(string(dtInvalidValidityDates), validFrom, validThru)
}

func (ge *SrvError) AuthNoMetaData(endpoint string) string {
	return fmt.Sprintf(string(authNoMetaData), endpoint)
}

func (ge *SrvError) AuthInvalidToken() string {
	return fmt.Sprintf(string(authInvalidToken))
}

func (ge *SrvError) AuthNilToken() string {
	return fmt.Sprintf(string(authNilToken))
}

func (ge *SrvError) AuthNilClaim(claimType string) string {
	return fmt.Sprintf(string(authNilClaim), claimType)
}

func (ge *SrvError) AuthInvalidClaim(claimType string) string {
	return fmt.Sprintf(string(authInvalidClaim), claimType)
}

func (ge *SrvError) BrkBadMarshall(objToMarshal string, err error) string {
	return fmt.Sprintf(string(brkBadMarshall), objToMarshal, err)
}

func (ge *SrvError) BrkNoMessageSent(objToMarshal string, err error) string {
	return fmt.Sprintf(string(brkNoMessageSent), objToMarshal, err)
}

func (ge *SrvError) BrkNoConnection(err error) string {
	return fmt.Sprintf(string(brkNoConnection), err)
}

func (ge *SrvError) BrkUnableToSetSubs(topic string, err error) string {
	return fmt.Sprintf(string(brkUnableToSetSubs), topic, err)
}

func (ge *SrvError) BrkBadUnMarshall(topic string, message []byte, err error) string {
	return fmt.Sprintf(string(brkBadUnMarshall), topic, message, err)
}
