# Booking App

This is an AWS Lambda function that books a squash court.

A message containing the details of the court to be booked is placed on SQS.  
When this Lambda function is triggered it reads the next message from the SQS queue  
and books the court according to the variableis in the message.

The message needs to be in the following format:

```json
{
    "Court": "1",
    "Days": "21",
    "Hour": "19",
    "Min": "50",
    "Timeslot": "17"
}
```

## TODO
* logging
* testing
* multiple messages
* how to handle bookings beyond 21 days
