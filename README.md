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
* ~~Use environment variables~~
* create a function to derive the timeslot parameter from the court number and time
* create a function to derive the days parameter from the date
* logging
* testing
* multiple messages
* how to handle bookings beyond 21 days - possibly link to the function to derive the days parameter
* front end to add bookings
