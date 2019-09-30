# Booking App

This is an AWS Lambda function that books a squash court.

A message containing the details of the court to be booked is placed on SQS.
When this Lambda function is triggered it reads the next message from the SQS queue
and books the court according to the parameters in the message.

The message is encoded in json and should to be in the following format:

```json
{
    "Id": "1",
    "Date": "19/10/2019 19:50",
}
```

The Lambda function also requires 3 Environment Variables:
* BOOKING_QUEUE - Which SQS to read messages from
* MEMBERSHIP_NUMBER - What membership number should be used to book a court
* NAME - The full name of the person who is booking the court

## TODO
* ~~Use environment variables~~
* ~~create a function to derive the timeslot parameter from the court number and time~~
* ~~create a function to derive the days parameter from the date~~
* logging
* testing
* multiple messages
* how to handle bookings beyond 21 days - possibly link to the function to derive the days parameter
* front end to add bookings
