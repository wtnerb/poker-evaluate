# Evaluate
This microservice will evaluate and select the winner from a hand of poker.

The expected input is a poker table object as per the `poker/models` package. It will respond with the winner of the hand and the split of who gets what money.

Notably, how to handle ties or when someone has gone all-in add extra complexity to the problem. This will likely force the models to be updated.

This service will (at time of writing, subject to change) update the database with the players totals being updated based upon their current winnings (if any). Perhaps a different database service should be used?

Regardless, version 0.1 of this service is to recieve a post request with the table state information and (after validating that it is a game end table state) return with the winner of the hand.