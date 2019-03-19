# Evaluate
This microservice will evaluate and select the winner from a hand of poker.

The expected input is a poker table object as per the `poker/models` package. It will respond with the winner(s) of the hand.

Notably, how to handle ties or when someone has gone all-in add extra complexity to the problem. This will likely force the models to be updated.

Regardless, version 0.1 of this service is to recieve a post request with the table state information and (after validating that it is a game end table state) return with the winner of the hand.

## Log
2019-03-19 v0.1 server works, can compare hands of different rank, some error checking, test coverage 91%